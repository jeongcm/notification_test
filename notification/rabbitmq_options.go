package notification

import "test/monitor"

type durableQueueKey struct{}
type headersKey struct{}
type queueArgumentsKey struct{}
type exchangeKey struct{}
type requeueOnErrorKey struct{}
type deliveryMode struct{}
type priorityKey struct{}
type durableExchange struct{}

// DurableQueue creates a durable queue when subscribing.
func DurableQueue() monitor.SubscribeOption {
	return setSubscribeOption(durableQueueKey{}, true)
}

// Headers adds headers used by the headers exchange
func Headers(h map[string]interface{}) monitor.SubscribeOption {
	return setSubscribeOption(headersKey{}, h)
}

// QueueArguments sets arguments for queue creation
func QueueArguments(h map[string]interface{}) monitor.SubscribeOption {
	return setSubscribeOption(queueArgumentsKey{}, h)
}

// RequeueOnError calls Nack(muliple:false, requeue:true) on amqp delivery when handler returns error
func RequeueOnError() monitor.SubscribeOption {
	return setSubscribeOption(requeueOnErrorKey{}, true)
}

// DeliveryMode sets a delivery mode for publishing
func DeliveryMode(value uint8) monitor.PublishOption {
	return setPublishOption(deliveryMode{}, value)
}

// Priority sets a priority level for publishing
func Priority(value uint8) monitor.PublishOption {
	return setPublishOption(priorityKey{}, value)
}
