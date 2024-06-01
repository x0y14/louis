package python

type Transpiler struct {
	*Tokenizer
	*Parser
	*Commonizer
	*Specializer
	*Generator
}
