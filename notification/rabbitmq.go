package notification

// Package rabbitmq provides a RabbitMQ Openstack Monitor

import (
	"context"
	"errors"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"net"
	"net/url"
	"sync"
	"test/monitor"
	"time"
)

func init() {
	monitor.RegisterClusterMonitorCreationFunc("type.openstack", New)
}

// OsloMessage notification event message body 구조체
type OsloMessage struct {
	EventType string                 `json:"event_type"`
	Payload   map[string]interface{} `json:"payload"`
}

type notification struct {
	conn           *rabbitMQConn
	opts           monitor.Options
	prefetchCount  int
	prefetchGlobal bool
	mtx            sync.Mutex
	wg             sync.WaitGroup
	address        string
	auth           string
}

type subscriber struct {
	mtx       sync.Mutex
	mayRun    bool
	opts      monitor.SubscribeOptions
	ch        *amqp.Channel
	queueArgs map[string]interface{}
	notifier  *notification
	fn        func(msg amqp.Delivery)
	headers   map[string]interface{}
}

type publication struct {
	d   amqp.Delivery
	m   *monitor.Message
	t   string
	err error
}

func (p *publication) Topic() string {
	return p.t
}

func (p *publication) Message() *monitor.Message {
	return p.m
}

func (s *subscriber) Options() monitor.SubscribeOptions {
	return s.opts
}

func (s *subscriber) Unsubscribe() error {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.mayRun = false
	if s.ch != nil {
		return s.ch.Close()
	}
	return nil
}

func (s *subscriber) resubscribe() {
	minResubscribeDelay := 100 * time.Millisecond
	maxResubscribeDelay := 30 * time.Second
	expFactor := time.Duration(2)
	reSubscribeDelay := minResubscribeDelay

	//loop until unsubscribe
	for {
		s.mtx.Lock()
		mayRun := s.mayRun
		s.mtx.Unlock()
		if !mayRun {
			// we are unsubscribed, showdown routine
			return
		}

		select {
		//check shutdown case
		case <-s.notifier.conn.close:
			//yep, its shutdown case
			return
			//wait until we reconect to rabbit
		case <-s.notifier.conn.waitConnection:
		}

		// it may crash (panic) in case of Consume without connection, so recheck it
		s.notifier.mtx.Lock()
		if !s.notifier.conn.connected {
			s.notifier.mtx.Unlock()
			continue
		}
		ch, sub, err := s.notifier.conn.Consume(
			s.opts.Queue,
			s.opts.AutoAck,
		)

		s.notifier.mtx.Unlock()
		switch err {
		case nil:
			reSubscribeDelay = minResubscribeDelay
			s.mtx.Lock()
			s.ch = ch
			s.mtx.Unlock()
		default:
			if reSubscribeDelay > maxResubscribeDelay {
				reSubscribeDelay = maxResubscribeDelay
			}
			time.Sleep(reSubscribeDelay)
			reSubscribeDelay *= expFactor
			continue
		}
		for d := range sub {
			s.notifier.wg.Add(1)
			s.fn(d)
			s.notifier.wg.Done()
		}
	}
}

// 큐를 지속적으로 유지하기위해 DurableQueue옵션, 데이터 안정성을 위해 DisableAutoAck, AckOnSuccess를 사용한다
// 특히 핸들러 에러 시 메시지의 requeue여부를 플래그로 전달하며 이 플래그에 따라 RequeueOnError() 옵션을 호출한다.
func (n *notification) Subscribe(queueName string, handler monitor.Handler, requeueOnError bool, opts ...monitor.SubscribeOption) (monitor.Subscriber, error) {
	opts = append(opts, monitor.Queue(queueName), monitor.DisableAutoAck())

	if requeueOnError == true {
		opts = append(opts, RequeueOnError())
	}

	return n.subscribe(handler, opts...)
}

func (n *notification) subscribe(handler monitor.Handler, opts ...monitor.SubscribeOption) (monitor.Subscriber, error) {
	if n.conn == nil {
		return nil, errors.New("not connected openstack notification")
	}

	options := monitor.SubscribeOptions{
		Context: context.Background(),
		AutoAck: true,
	}
	for _, o := range opts {
		o(&options)
	}

	var requeueOnError bool
	requeueOnError, _ = options.Context.Value(requeueOnErrorKey{}).(bool)

	var qArgs map[string]interface{}
	if qa, ok := options.Context.Value(queueArgumentsKey{}).(map[string]interface{}); ok {
		qArgs = qa
	}

	var headers map[string]interface{}
	if h, ok := options.Context.Value(headersKey{}).(map[string]interface{}); ok {
		headers = h
	}

	fn := func(msg amqp.Delivery) {
		header := make(map[string]string)
		for k, v := range msg.Headers {
			header[k], _ = v.(string)
		}
		m := &monitor.Message{
			Header: header,
			Body:   msg.Body,
		}

		p := &publication{d: msg, m: m, t: msg.RoutingKey}
		p.err = handler(p)
		if p.err == nil && !options.AutoAck {
			if err := msg.Ack(false); err != nil {
				log.Printf("could not acknowledge on a delivery. cause: %v\n", err)
			}
		} else if p.err != nil && !options.AutoAck {
			if err := msg.Nack(false, requeueOnError); err != nil {
				log.Printf("could not negatively acknowledge on a delivery. cause: %v\n", err)
			}
		}
	}

	sret := &subscriber{opts: options, mayRun: true, notifier: n,
		fn: fn, headers: headers, queueArgs: qArgs}

	go sret.resubscribe()

	return sret, nil
}

func (n *notification) Options() monitor.Options {
	return n.opts
}

func (n *notification) Connect() error {
	conn := n.conn

	if n.conn == nil {
		conn = newRabbitMQConn(
			0,
			false,
			n.auth,
			n.address)
	}

	conf := defaultAmqpConfig
	conf.TLSClientConfig = n.opts.TLSConfig

	if err := conn.Connect(n.opts.Secure, &conf); err != nil {
		return err
	}

	n.conn = conn
	return nil
}

func (n *notification) Disconnect() error {
	if n.conn == nil {
		return errors.New("connection is nil")
	}
	n.mtx.Lock()
	ret := n.conn.Close()
	n.mtx.Unlock()

	n.wg.Wait() // wait all goroutines
	return ret
}

// New 함수는 새로운 monitor interface 를 생성한다.
func New(serverURL string) monitor.Monitor {
	//TODO auth 의 경우 임시로 ID:PASSWORD(ex.guest:guest)를 쓰지만
	//	   사용자 입력에 의한 Cluster 의 MetaData 로 저장될 필요가 있음.
	//     마찬가지로 임시로 client 의 api server url 과 고정된 port(ex.192.168.1.1:5672) 를 쓰지만
	//     사용자 입력에 의한 Cluster 의 MetaData 로 저장될 필요가 있음.
	auth := "guest:guest"
	defaultPort := "5672"

	u, _ := url.Parse(serverURL)
	ip, _, _ := net.SplitHostPort(u.Host)

	options := monitor.Options{
		Context: context.Background(),
	}

	log.Println(fmt.Sprintf("%s:%s", ip, defaultPort))
	return &notification{
		opts:    options,
		auth:    auth,
		address: fmt.Sprintf("%s:%s", ip, defaultPort),
	}
}
