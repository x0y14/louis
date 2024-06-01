package interfaces

import "louis/languages/ir"

type Commonizer interface {
	Commonize(n Node) (*ir.Node, error)
}
