package queue

import "github.com/gopi-frame/contract/event"

// Subcriber subcriber
type Subcriber interface {
	event.Subscriber
	OnBeforeHandle(event event.Event) bool
	OnAfterHandle(event event.Event) bool
	OnFailedHandle(event event.Event) bool
	OnProgressUpdated(event event.Event) bool
}
