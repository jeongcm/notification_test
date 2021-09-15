package notification

import (
	"context"
	"test/monitor"
)

// setSubscribeOption returns a function to setup a context with given value
func setSubscribeOption(k, v interface{}) monitor.SubscribeOption {
	return func(o *monitor.SubscribeOptions) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, k, v)
	}
}

// setPublishOption returns a function to setup a context with given value
func setPublishOption(k, v interface{}) monitor.PublishOption {
	return func(o *monitor.PublishOptions) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, k, v)
	}
}
