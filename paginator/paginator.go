package paginator

import (
	"encoding/json"
	"github.com/gopi-frame/contract"
)

// Paginator is a paginator interface
type Paginator[T any] interface {
	contract.Mappable[string, any]
	contract.Jsonable
	json.Marshaler
	json.Unmarshaler
	// Items returns the items of current page
	Items() []T
	// Total returns the total number of items
	Total() int64
	// LastPage returns the number of the last page.
	LastPage() int
}
