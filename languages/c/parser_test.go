package c

import (
	"testing"
)

func TestParser_Parse(t *testing.T) {
	tokenizer := Tokenizer{}
	tok, err := tokenizer.Tokenize("int main() { return 0; }")
	if err != nil {
		t.Fatal(err)
	}
	parser := Parser{}
	nod, err := parser.Parse(tok)
	if err != nil {
		t.Fatal(err)
	}
	_ = nod
}
