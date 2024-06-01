package c

import "louis/interfaces"

type Tokenizer struct {
}

func (t *Tokenizer) Tokenize(s string) (interfaces.Token, error) {
	return &Token{}, nil
}
