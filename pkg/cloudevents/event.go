package cloudevents

import "github.com/cloudevents/sdk-go/pkg/cloudevents/context"

// Event represents the canonical representation of a CloudEvent.
type Event struct {
	Context context.EventContext
	Data    interface{}
}
