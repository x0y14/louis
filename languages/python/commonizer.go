package python

import (
	"louis/interfaces"
)

type Commonizer struct {
}

func (c *Commonizer) Commonize(n interfaces.Node) (interfaces.Node, error) {
	_ = n
	return nil, nil
}
