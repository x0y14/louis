package c

import (
	"fmt"
	"louis/interfaces"
)

type NodeKind int

const (
	ND_UNKNOWN NodeKind = iota
	ND_VARIABLE_DECL
	ND_FUNCTION_DECL
	ND_FUNCTION_DEFN
	ND_BLOCK
	ND_RETURN
	ND_IF
	ND_WHILE
	ND_FOR
	ND_EXPR_ADD
	ND_EXPR_SUB
	ND_EXPR_MUL
	ND_EXPR_DIV
	ND_EXPR_MOD
	ND_EXPR_EQL
	ND_EXPR_LSS
	ND_EXPR_GTE
	ND_EXPR_ASSIGN
	ND_EXPR_NOT
	ND_EXPR_NEQ
	ND_EXPR_LEQ
	ND_EXPR_GEQ
	ND_EXPR_AND
	ND_EXPR_OR
	ND_IDENT
	ND_LIT_INT
	ND_LIT_FLOAT
	ND_LIT_STRING
	ND_LIT_NULL
	ND_CALL
	ND_ARGS
)

func (n NodeKind) String() string {
	k := [...]string{
		ND_UNKNOWN:       "UNKNOWN",
		ND_VARIABLE_DECL: "VARIABLE_DECL",
		ND_FUNCTION_DECL: "FUNCTION_DECL",
		ND_FUNCTION_DEFN: "FUNCTION_DEFN",
		ND_BLOCK:         "BLOCK",
		ND_RETURN:        "RETURN",
		ND_IF:            "IF",
		ND_WHILE:         "WHILE",
		ND_FOR:           "FOR",
		ND_EXPR_ADD:      "EXPR_ADD",
		ND_EXPR_SUB:      "EXPR_SUB",
		ND_EXPR_MUL:      "EXPR_MUL",
		ND_EXPR_DIV:      "EXPR_DIV",
		ND_EXPR_MOD:      "EXPR_MOD",
		ND_EXPR_EQL:      "EXPR_EQL",
		ND_EXPR_LSS:      "EXPR_LSS",
		ND_EXPR_GTE:      "EXPR_GTE",
		ND_EXPR_ASSIGN:   "EXPR_ASSIGN",
		ND_EXPR_NOT:      "EXPR_NOT",
		ND_EXPR_NEQ:      "EXPR_NEQ",
		ND_EXPR_LEQ:      "EXPR_LEQ",
		ND_EXPR_GEQ:      "EXPR_GEQ",
		ND_EXPR_AND:      "EXPR_AND",
		ND_EXPR_OR:       "EXPR_OR",
		ND_IDENT:         "IDENT",
		ND_LIT_INT:       "LIT_INT",
		ND_LIT_FLOAT:     "LIT_FLOAT",
		ND_LIT_STRING:    "LIT_STRING",
		ND_LIT_NULL:      "LIT_NULL",
		ND_CALL:          "CALL",
		ND_ARGS:          "ARGS",
	}
	return k[n]
}

type Node struct {
	Kind       NodeKind
	Next       interfaces.Node
	Literal    interfaces.Literal
	Identifier interfaces.Identifier
	Field      interfaces.Field
}

func (n *Node) String() string {
	switch n.GetKind().(NodeKind) {
	case ND_LIT_INT, ND_LIT_FLOAT, ND_LIT_STRING: // literal
		return fmt.Sprintf("NODE.LITERAL{%s}", n.GetLiteral().String())
	case ND_IDENT:
		return fmt.Sprintf("NODE.IDENT{%s}", n.GetIdentifier().String())
	default:
		return fmt.Sprintf("NODE{%s}", n.GetKind().String())
	}
}

func (n *Node) GetLiteral() interfaces.Literal {
	return n.Literal
}

func (n *Node) GetIdentifier() interfaces.Identifier {
	return n.Identifier
}

func (n *Node) GetField() interfaces.Field {
	return n.Field
}

func (n *Node) GetKind() interfaces.NodeKind {
	return n.Kind
}

func (n *Node) GetNext() interfaces.Node {
	return n.Next
}
