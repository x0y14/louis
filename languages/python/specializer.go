package python

import (
	"louis/interfaces"
)

type Specializer struct {
}

func (s *Specializer) Specialize(n interfaces.Node) (interfaces.Node, error) {
	_ = n
	return nil, nil
}
