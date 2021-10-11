package monitor

import "errors"

type clusterMonitorCreationFunc func(string) Monitor

var clusterMonitorCreationFuncMap map[string]clusterMonitorCreationFunc

// RegisterClusterMonitorCreationFunc 는 클러스터 타입별 Monitor 구조체 생성 함수의 맵이다.
func RegisterClusterMonitorCreationFunc(typeCode string, fn clusterMonitorCreationFunc) {
	if clusterMonitorCreationFuncMap == nil {
		clusterMonitorCreationFuncMap = make(map[string]clusterMonitorCreationFunc)
	}

	clusterMonitorCreationFuncMap[typeCode] = fn
}

// Monitor monitor 인터페이스
type Monitor interface {
	Options() Options
	Connect() error
	Disconnect() error
	Subscribe(queueName string, h Handler, requeueOnError bool, opts ...SubscribeOption) (Subscriber, error)
}

// Handler is used to process messages via a subscription of a topic.
// The handler is passed a publication interface which contains the
// message and optional Ack method to acknowledge receipt of the message.
type Handler func(Event) error

// Message structure for monitor
type Message struct {
	Header map[string]string
	Body   []byte
}

// Event is given to a subscription handler for processing
type Event interface {
	Topic() string
	Message() *Message
}

// Subscriber is a convenience return type for the Subscribe method
type Subscriber interface {
	Options() SubscribeOptions
	Unsubscribe() error
}

// New 는 클러스터 타입별 모니터 인터페이스를 초기화하는 함수
func New(typeCode, serverURL string) (Monitor, error) {
	if fn, ok := clusterMonitorCreationFuncMap[typeCode]; ok {
		return fn(serverURL), nil
	}

	return nil, errors.New("not found cluster")
}
