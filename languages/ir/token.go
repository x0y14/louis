package ir

import "louis/interfaces"

type Token struct {
}

func (t *Token) String() string {
	return ""
}

func (t *Token) GetKind() int {
	return 0
}

func (t *Token) GetNext() interfaces.Token {
	return nil
}
