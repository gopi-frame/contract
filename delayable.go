package contract

import (
	"encoding/json"
	"time"
)

// Delayable delayable
type Delayable[T any] interface {
	json.Marshaler
	json.Unmarshaler
	// Until returns the available time
	Until() time.Time
	// Value returns the underlying value
	Value() T
}
