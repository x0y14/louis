package interfaces

type NodeKind interface {
	String() string
}

type Node interface {
	String() string
	GetKind() NodeKind
	GetLiteral() Literal
	GetIdentifier() Identifier
	GetField() Field
}
