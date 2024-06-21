package ir

type ReturnField struct {
	Values *Node
}

type BlockField struct {
	Stmts *Node
}

type VariableDeclField struct {
}

type VariableDefField struct {
}

type FunctionDefField struct {
	ReturnType *Node
	Ident      *Node
	Params     *Node
	Block      *Node
}
