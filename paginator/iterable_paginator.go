package paginator

// IterablePaginator is an interface that extends the Paginator interface and adds a Next() method.
// The Next() method is used to advance the paginator to the next page of results.
type IterablePaginator[T any] interface {
	Paginator[T]
	Next() bool
}
