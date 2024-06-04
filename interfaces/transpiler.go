package interfaces

type Transpiler interface {
	Tokenize(s string) (Token, error)
	Parse(t Token) (Node, error)
	Commonize(n Node) (Node, error)
	Specialize(n Node) (Node, error)
	Generate(n Node) (string, error)
}
