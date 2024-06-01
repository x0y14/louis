package c

import "louis/interfaces"

type Generator struct {
}

func (g *Generator) Generate(n interfaces.Node) (string, error) {
	return "", nil
}
