package contract

// Ordered ordered interface
type Ordered interface {
	Less(any) bool
}
