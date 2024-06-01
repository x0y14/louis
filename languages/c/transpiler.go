package c

type Transpiler struct {
	*Tokenizer
	*Parser
	*Commonizer
	*Specializer
	*Generator
}
