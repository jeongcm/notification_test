package monitor

// Monitor openstack client monitor 인터페이스
type Monitor interface {
	Options() Options
	Connect() error
	Disconnect() error
	Subscribe(clusterID int, topic string, h Handler, requeueOnError bool, opts ...SubscribeOption) (Subscriber, error)
}

// Handler is used to process messages via a subscription of a topic.
// The handler is passed a publication interface which contains the
// message and optional Ack method to acknowledge receipt of the message.
type Handler func(int, Event) error

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
