package c

import (
	"fmt"
	"louis/interfaces"
	"strconv"
)

type TokenKind int

func (t TokenKind) String() string {
	k := [...]string{
		TK_ILLEGAL: "ILLEGAL",
		TK_EOF:     "EOF",
		TK_COMMENT: "COMMENT",

		TK_IDENT:  "IDENT",
		TK_INT:    "INT",
		TK_FLOAT:  "FLOAT",
		TK_CHAR:   "CHAR",
		TK_STRING: "STRING",

		TK_ADD: "+",
		TK_SUB: "-",
		TK_MUL: "*",
		TK_DIV: "/",
		TK_MOD: "%",

		TK_EQL:    "==",
		TK_LSS:    "<",
		TK_GTR:    ">",
		TK_ASSIGN: "=",
		TK_NOT:    "!",

		TK_NEQ: "!=",
		TK_LEQ: "<=",
		TK_GEQ: ">=",

		TK_LAND: "&&",
		TK_LOR:  "||",

		TK_LRB: "(",
		TK_LSB: "[",
		TK_LCB: "{",
		TK_RRB: ")",
		TK_RSB: "]",
		TK_RCB: "}",

		TK_DOT:   ".",
		TK_COMMA: ",",
		TK_COLON: ":",
		TK_SEMI:  ";",
	}
	return k[t]
}

// 参考: [Golang](https://github.com/golang/go/blob/master/src/go/token/token.go)
const (
	TK_ILLEGAL TokenKind = iota
	TK_EOF
	TK_COMMENT
	TK_IDENT  // i
	TK_INT    // 1
	TK_FLOAT  // 1.2
	TK_CHAR   // 'a'
	TK_STRING // "abc"
	TK_ADD    // +
	TK_SUB    // -
	TK_MUL    // *
	TK_DIV    // /
	TK_MOD    // %
	TK_EQL    // ==
	TK_LSS    // <
	TK_GTR    // >
	TK_ASSIGN // =
	TK_NOT    // !
	TK_NEQ    // !=
	TK_LEQ    // <=
	TK_GEQ    // >=
	TK_LAND   // &&
	TK_LOR    // ||
	TK_LRB    // (
	TK_LSB    // [
	TK_LCB    // {
	TK_RRB    // )
	TK_RSB    // ]
	TK_RCB    // }
	TK_DOT    // .
	TK_COMMA  // ,
	TK_COLON  // :
	TK_SEMI   // ;
)

type Token struct {
	Kind interfaces.TokenKind
	s    string
	i    int
	f    float64
	c    rune
	Next *Token
}

func (t *Token) String() string {
	return fmt.Sprintf("Token(%s)", t.Kind.String())
}

func (t *Token) GetKind() interfaces.TokenKind {
	return t.Kind
}

func (t *Token) GetNext() interfaces.Token {
	return t.Next
}

func symbolToTokenKind(s string) (TokenKind, bool) {
	k := map[string]TokenKind{
		"==": TK_EQL,

		"!=": TK_NEQ,
		"<=": TK_LEQ,
		">=": TK_GEQ,

		"&&": TK_LAND,
		"||": TK_LOR,

		"+": TK_ADD,
		"-": TK_SUB,
		"*": TK_MUL,
		"/": TK_DIV,
		"%": TK_MOD,

		"<": TK_LSS,
		">": TK_GTR,
		"=": TK_ASSIGN,
		"!": TK_NOT,

		"(": TK_LRB,
		"[": TK_LSB,
		"{": TK_LCB,
		")": TK_RRB,
		"]": TK_RSB,
		"}": TK_RCB,

		".": TK_DOT,
		",": TK_COMMA,
		":": TK_COLON,
		";": TK_SEMI,
	}

	v, ok := k[s]
	return v, ok
}

func numericTokenKind(num []rune) (TokenKind, bool) {
	// パースできるか? / 数値として正しいか?
	_, err := strconv.ParseFloat(string(num), 64)
	if err != nil {
		return TK_ILLEGAL, false
	}

	// dotが含まれるか?
	isIncludeDot := false
	for _, r := range num {
		if r == '.' {
			isIncludeDot = true
			break
		}
	}
	if isIncludeDot {
		return TK_FLOAT, true
	}
	return TK_INT, true
}
