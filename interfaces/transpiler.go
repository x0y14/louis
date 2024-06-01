package interfaces

import "louis/languages/ir"

type Transpiler interface {
	Tokenize(s string) (Token, error)
	Parse(t Token) (Node, error)
	Commonize(n Node) (*ir.Node, error)
	Specialize(n *ir.Node) (Node, error)
	Generate(n Node) (string, error)
}
