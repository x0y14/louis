package c

import (
	"louis/interfaces"
	"louis/languages/ir"
)

type Specializer struct {
}

func (s *Specializer) Specialize(n *ir.Node) (interfaces.Node, error) {
	return nil, nil
}
