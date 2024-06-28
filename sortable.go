package contract

// Sortable sortable
type Sortable[T any] interface {
	Sort(func(T, T) int)
}
