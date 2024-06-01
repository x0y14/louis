package interfaces

type Parser interface {
	Parse(t Token) (Node, error)
}
