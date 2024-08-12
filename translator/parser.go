package translator

type Parser interface {
	Parse([]byte) (MessagePack, error)
}
