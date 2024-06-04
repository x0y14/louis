package c

import (
	"louis/interfaces"
)

type Specializer struct {
}

func (s *Specializer) Specialize(n interfaces.Node) (interfaces.Node, error) {
	return nil, nil
}
