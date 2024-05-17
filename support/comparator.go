package support

// Comparator comparator
type Comparator[T any] interface {
	Compare(a, b T) int
}
