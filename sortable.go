package contract

// Sortable sortable
type Sortable[T any] interface {
	// Sort sorts
	Sort(func(T, T) int)
}
