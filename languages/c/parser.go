package c

import "louis/interfaces"

type Parser struct {
}

func (p *Parser) Parse(t interfaces.Token) (interfaces.Node, error) {
	return nil, nil
}
