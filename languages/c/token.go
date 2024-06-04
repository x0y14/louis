package c

import (
	"fmt"
	"louis/interfaces"
	"strconv"
)

type TokenKind int

func (t TokenKind) String() string {
	k := [...]string{
		ILLEGAL: "ILLEGAL",
		EOF:     "EOF",
		COMMENT: "COMMENT",

		IDENT:  "IDENT",
		INT:    "INT",
		FLOAT:  "FLOAT",
		CHAR:   "CHAR",
		STRING: "STRING",

		ADD: "+",
		SUB: "-",
		MUL: "*",
		DIV: "/",
		MOD: "%",

		EQL:    "==",
		LSS:    "<",
		GTR:    ">",
		ASSIGN: "=",
		NOT:    "!",

		NEQ: "!=",
		LEQ: "<=",
		GEQ: ">=",

		LAND: "&&",
		LOR:  "||",

		LRB: "(",
		LSB: "[",
		LCB: "{",
		RRB: ")",
		RSB: "]",
		RCB: "}",

		DOT:   ".",
		COMMA: ",",
		COLON: ":",
		SEMI:  ";",
	}
	return k[t]
}

// 参考: [Golang](https://github.com/golang/go/blob/master/src/go/token/token.go)
const (
	ILLEGAL TokenKind = iota
	EOF
	COMMENT

	IDENT  // i
	INT    // 1
	FLOAT  // 1.2
	CHAR   // 'a'
	STRING // "abc"

	ADD // +
	SUB // -
	MUL // *
	DIV // /
	MOD // %

	EQL    // ==
	LSS    // <
	GTR    // >
	ASSIGN // =
	NOT    // !

	NEQ // !=
	LEQ // <=
	GEQ // >=

	LAND // &&
	LOR  // ||

	LRB // (
	LSB // [
	LCB // {
	RRB // )
	RSB // ]
	RCB // }

	DOT   // .
	COMMA // ,
	COLON // :
	SEMI  // ;
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
		"==": EQL,

		"!=": NEQ,
		"<=": LEQ,
		">=": GEQ,

		"&&": LAND,
		"||": LOR,

		"+": ADD,
		"-": SUB,
		"*": MUL,
		"/": DIV,
		"%": MOD,

		"<": LSS,
		">": GTR,
		"=": ASSIGN,
		"!": NOT,

		"(": LRB,
		"[": LSB,
		"{": LCB,
		")": RRB,
		"]": RSB,
		"}": RCB,

		".": DOT,
		",": COMMA,
		":": COLON,
		";": SEMI,
	}

	v, ok := k[s]
	return v, ok
}

func numericTokenKind(num []rune) (TokenKind, bool) {
	// パースできるか? / 数値として正しいか?
	_, err := strconv.ParseFloat(string(num), 64)
	if err != nil {
		return ILLEGAL, false
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
		return FLOAT, true
	}
	return INT, true
}
