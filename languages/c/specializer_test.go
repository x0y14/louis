package c

import (
	"github.com/google/go-cmp/cmp"
	"louis/interfaces"
	"louis/languages/ir"
	"testing"
)

func TestSpecializer_Specialize(t *testing.T) {
	specializer := Specializer{}

	expectedCNode := func() interfaces.Node {
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
		return nod
	}

	irNode := ir.Node{}
	cTok, err := specializer.Specialize(&irNode)
	if err != nil {
		t.Fatal(err)
	}

	want := expectedCNode()
	got := cTok

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}
