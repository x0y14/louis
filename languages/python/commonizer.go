package python

import (
	"louis/interfaces"
	"louis/languages/ir"
)

type Commonizer struct {
}

func (c *Commonizer) Commonize(n interfaces.Node) (*ir.Node, error) {
	_ = n
	return nil, nil
}
