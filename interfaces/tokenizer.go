package interfaces

type Tokenizer interface {
	Tokenize(s string) (Token, error)
}
