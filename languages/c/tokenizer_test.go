package c

import (
	"fmt"
	"log"
	"testing"
)

func TestTokenizer_Tokenize(t *testing.T) {
	tokenizer := Tokenizer{}
	tok, err := tokenizer.Tokenize("int main() { return 0; }")
	if err != nil {
		log.Fatal(err)
	}
	for {
		fmt.Println(tok.String())
		if tok.GetKind() == TK_EOF {
			break
		}
		tok = tok.GetNext()
	}
}
