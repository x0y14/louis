package main

import (
	"fmt"
	"log"
	"louis"
	"louis/interfaces"
	"louis/languages/c"
	"louis/languages/python"
)

var l louis.Louis

const (
	ILLEGAL louis.SupportedLanguage = iota
	C
	PYTHON
)

func init() {
	// language setup
	l = louis.Louis{Transpilers: make(map[louis.SupportedLanguage]interfaces.Transpiler)}

	l.Transpilers[C] = &c.Transpiler{
		Tokenizer:   &c.Tokenizer{},
		Parser:      &c.Parser{},
		Commonizer:  &c.Commonizer{},
		Specializer: &c.Specializer{},
		Generator:   &c.Generator{},
	}

	l.Transpilers[PYTHON] = &python.Transpiler{
		Tokenizer:   &python.Tokenizer{},
		Parser:      &python.Parser{},
		Commonizer:  &python.Commonizer{},
		Specializer: &python.Specializer{},
		Generator:   &python.Generator{},
	}
}

func main() {
	// settings from cmd args
	userInput := `int main() { return 0; }`
	inputLang := C
	outputLang := PYTHON

	// input -> ir
	inT, err := l.Transpilers[inputLang].Tokenize(userInput)
	if err != nil {
		log.Fatal(err)
	}
	inN, err := l.Transpilers[inputLang].Parse(inT)
	if err != nil {
		log.Fatal(err)
	}
	irN, err := l.Transpilers[inputLang].Commonize(inN)
	if err != nil {
		log.Fatal(err)
	}

	// ir -> output
	outN, err := l.Transpilers[outputLang].Specialize(irN)
	if err != nil {
		log.Fatal(err)
	}
	result, err := l.Transpilers[outputLang].Generate(outN)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", result)
}
