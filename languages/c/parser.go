package c

import (
	"fmt"
	"louis/interfaces"
)

type Parser struct {
	curt interfaces.Token
}

func (p *Parser) isEof() bool {
	return p.curt.GetKind().(TokenKind) == TK_EOF
}

func (p *Parser) peekKind(k TokenKind) *Token {
	if p.curt.GetKind() == k {
		return p.curt.(*Token)
	}
	return nil
}

func (p *Parser) peekOrder(ks []TokenKind) *Token {
	current := p.curt
	for !p.isEof() && len(ks) > 0 {
		if current.GetKind() == ks[0] {
			ks = ks[1:]
			current = current.GetNext()
		} else {
			return nil
		}
	}
	return p.curt.(*Token)
}

func (p *Parser) consumeKind(k TokenKind) *Token {
	if p.curt.GetKind() == k {
		t := p.curt
		p.curt = t.GetNext()
		return t.(*Token)
	}
	return nil
}

func (p *Parser) consumeIdent(id string) *Token {
	ident := p.peekKind(TK_IDENT)
	if ident.s == id {
		return p.consumeKind(TK_IDENT)
	}
	return nil
}

func (p *Parser) expectKind(k TokenKind) (*Token, error) {
	if p.curt.GetKind() == k {
		expected := p.curt
		p.curt = expected.GetNext()
		return expected.(*Token), nil
	}
	return nil, fmt.Errorf(
		"unexpected token error: want: %s, got: %s", k.String(), p.curt.GetKind().String())
}

func (p *Parser) variableDecl() (*Node, error) {
	return nil, nil
}

func (p *Parser) functionParam() (*Node, error) {
	return nil, nil
}

func (p *Parser) primary() (*Node, error) {
	switch {
	case p.curt.GetKind() == TK_INT:
		i, _ := p.expectKind(TK_INT)
		return &Node{
			Kind:       ND_LIT_INT,
			Next:       nil,
			Literal:    &Literal[int]{i.i},
			Identifier: nil,
			Field:      nil,
		}, nil
	}
	return nil, nil
}

func (p *Parser) unary() (*Node, error) {
	return p.primary()
}

func (p *Parser) mul() (*Node, error) {
	return p.unary()
}

func (p *Parser) add() (*Node, error) {
	return p.mul()
}

func (p *Parser) relational() (*Node, error) {
	return p.add()
}

func (p *Parser) equality() (*Node, error) {
	return p.relational()
}

func (p *Parser) andor() (*Node, error) {
	return p.equality()
}

func (p *Parser) assign() (*Node, error) {
	return p.andor()
}

func (p *Parser) expr() (*Node, error) {
	return p.assign()
}

func (p *Parser) statement() (*Node, error) {
	switch {
	case p.consumeKind(TK_LCB) != nil: // block
		var head Node
		curt := &head

		for p.consumeKind(TK_RCB) == nil {
			stmt, err := p.statement()
			if err != nil {
				return nil, err
			}
			curt.Next = stmt
			curt = stmt
		}

		return &Node{
			Kind:       ND_BLOCK,
			Next:       nil,
			Literal:    nil,
			Identifier: nil,
			Field:      &BlockField{Stmts: head.Next.(*Node)},
		}, nil
	case p.consumeIdent(KW_RETURN.String()) != nil: // return
		var head Node
		curt := &head
		retValue, err := p.expr()
		if err != nil {
			return nil, err
		}
		curt.Next = retValue
		curt = retValue

		if _, err := p.expectKind(TK_SEMI); err != nil {
			return nil, err
		}
		return &Node{
			Kind:       ND_RETURN,
			Next:       nil,
			Literal:    nil,
			Identifier: nil,
			Field: &ReturnField{
				Values: head.Next.(*Node),
			},
		}, nil
	}
	return nil, nil
}

func (p *Parser) functionDefDecl() (*Node, error) {
	// `int` main (...) {...}
	fnRetType, err := p.expectKind(TK_IDENT)
	if err != nil {
		return nil, err
	}
	fnRetTypeNode := &Node{
		Kind:       ND_IDENT,
		Next:       nil,
		Literal:    nil,
		Identifier: &Identifier{Name: fnRetType.s},
	}
	// int `main` (...) {...}
	fnNameIdent, err := p.expectKind(TK_IDENT)
	if err != nil {
		return nil, err
	}
	fnNameIdentNode := &Node{
		Kind:       ND_IDENT,
		Next:       nil,
		Literal:    nil,
		Identifier: &Identifier{Name: fnNameIdent.s},
	}

	// int main `(...)` {...}
	// `(`
	_, err = p.expectKind(TK_LRB)
	if err != nil {
		return nil, err
	}
	// (`...`)
	fnParamsNode, err := p.functionParam()
	if err != nil {
		return nil, err
	}
	// `)`
	_, err = p.expectKind(TK_RRB)
	if err != nil {
		return nil, err
	}

	// `;` -> function declare
	if p.peekKind(TK_SEMI) != nil {
		return &Node{
			Kind:       0,
			Next:       nil,
			Literal:    nil,
			Identifier: nil,
		}, nil
	}

	// `{` -> function define
	// int main(...) `{...}`
	blockNode, err := p.statement()
	if err != nil {
		return nil, err
	}

	return &Node{
		Kind:       ND_FUNCTION_DEFN,
		Next:       nil,
		Literal:    nil,
		Identifier: nil,
		Field:      &FunctionDefField{ReturnType: fnRetTypeNode, Ident: fnNameIdentNode, Params: fnParamsNode, Block: blockNode},
	}, nil
}

func (p *Parser) program() (*Node, error) {
	switch {
	case p.peekOrder([]TokenKind{TK_IDENT, TK_IDENT, TK_SEMI}) != nil: // variable declare
		return p.variableDecl()
	case p.peekOrder([]TokenKind{TK_IDENT, TK_IDENT, TK_LRB}) != nil: // function define/declare
		return p.functionDefDecl()
	}
	return nil, nil
}

func (p *Parser) Parse(t interfaces.Token) (interfaces.Node, error) {
	p.curt = t
	var head Node
	curt := &head

	for !p.isEof() {
		n, err := p.program()
		if err != nil {
			return nil, err
		}
		curt.Next = n
		curt = n
	}

	return head.Next, nil
}
