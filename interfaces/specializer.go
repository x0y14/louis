package interfaces

import "louis/languages/ir"

type Specializer interface {
	Specialize(n *ir.Node) (Node, error)
}
