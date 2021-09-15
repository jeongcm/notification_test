package notification

//
// All credit to Mondo
//

import (
	"crypto/tls"
	"errors"
	"github.com/streadway/amqp"
	"log"
	"strings"
	"sync"
	"time"
)

var (
	// The amqp library does not seem to set these when using amqp.DialConfig
	// (even though it says so in the comments) so we set them manually to make
	// sure to not brake any existing functionality
	defaultAmqpConfig = amqp.Config{
		Heartbeat: 10 * time.Second,
		Locale:    "en_US",
	}
)

type rabbitMQConn struct {
	Connection     *amqp.Connection
	accountInfo    string
	address        string
	port           string
	prefetchCount  int
	prefetchGlobal bool

	sync.Mutex
	connected bool
	close     chan bool

	waitConnection chan struct{}
}

func newRabbitMQConn(prefetchCount int, prefetchGlobal bool, accountInfo string, address string) *rabbitMQConn {
	ret := &rabbitMQConn{
		prefetchCount:  prefetchCount,
		prefetchGlobal: prefetchGlobal,
		close:          make(chan bool),
		waitConnection: make(chan struct{}),
		accountInfo:    accountInfo,
		address:        address,
	}

	// its bad case of nil == waitConnection, so close it at start
	close(ret.waitConnection)
	return ret
}

func (r *rabbitMQConn) connect(secure bool, config *amqp.Config) error {
	// try connect
	if err := r.tryConnect(secure, config); err != nil {
		return err
	}

	// connected
	r.Lock()
	r.connected = true
	r.Unlock()

	// create reconnect loop
	go r.reconnect(secure, config)
	return nil
}

func (r *rabbitMQConn) reconnect(secure bool, config *amqp.Config) {
	// skip first connect
	var connect bool

	for {
		if connect {
			// try reconnect
			if err := r.tryConnect(secure, config); err != nil {
				time.Sleep(1 * time.Second)
				continue
			}

			// connected
			r.Lock()
			r.connected = true
			r.Unlock()
			//unblock resubscribe cycle - close channel
			//at this point channel is created and unclosed - close it without any additional checks
			close(r.waitConnection)
		}

		connect = true
		notifyClose := make(chan *amqp.Error)
		r.Connection.NotifyClose(notifyClose)

		// block until closed
		select {
		case <-notifyClose:
			// block all resubscribe attempt - they are useless because there is no connection to rabbitmq
			// create channel 'waitConnection' (at this point channel is nil or closed, create it without unnecessary checks)
			r.Lock()
			r.connected = false
			r.waitConnection = make(chan struct{})
			r.Unlock()
		case <-r.close:
			return
		}
	}
}

func (r *rabbitMQConn) Connect(secure bool, config *amqp.Config) error {
	r.Lock()
	// already connected
	if r.connected {
		r.Unlock()
		return nil
	}

	// check it was closed
	select {
	case <-r.close:
		r.close = make(chan bool)
	default:
		// no op
		// new conn
	}

	r.Unlock()

	return r.connect(secure, config)
}

func (r *rabbitMQConn) Close() error {
	r.Lock()
	defer r.Unlock()
	select {
	case <-r.close:
		return nil
	default:
		close(r.close)
		r.connected = false
	}

	return r.Connection.Close()
}

// registry 에서 rabbitmq서비스 정보를 얻어와 연결 시도
func (r *rabbitMQConn) doTryConnect(secure bool, config *amqp.Config) error {
	var err error
	url := "amqp://" + r.accountInfo + "@" + r.address + "/"

	if secure || config.TLSClientConfig != nil || strings.HasPrefix(url, "amqps://") {
		if config.TLSClientConfig == nil {
			config.TLSClientConfig = &tls.Config{
				InsecureSkipVerify: true,
			}
		}

		url = strings.Replace(url, "amqp://", "amqps://", 1)
	}

	r.Connection, err = amqp.DialConfig(url, *config)
	if err == nil { //Connection success
		log.Printf("connect success, %v", r.Connection.Config.Dial)
		return nil
	}

	return errors.New("could not connect to openstack notification")
}

func (r *rabbitMQConn) tryConnect(secure bool, config *amqp.Config) error {
	var err error

	if config == nil {
		config = &defaultAmqpConfig
	}

	if err = r.doTryConnect(secure, config); err != nil {
		return err
	}

	return nil
}

func (r *rabbitMQConn) Consume(queue string, autoAck bool) (*amqp.Channel, <-chan amqp.Delivery, error) {
	c, err := r.Connection.Channel()
	if err != nil {
		return nil, nil, err
	}

	deliveries, err := c.Consume(queue,
		"",      // consumer
		autoAck, // auto-ack
		false,   // exclusive
		false,   // no local
		false,   // no wait
		nil,     // arguments
	)
	if err != nil {
		return nil, nil, err
	}

	//deliveries, err := consumerChannel.ConsumeQueue(queue, autoAck)
	//if err != nil {
	//	return nil, nil, err
	//}

	return c, deliveries, nil
}
