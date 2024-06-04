package interfaces

type TokenKind interface {
	String() string
}

type Token interface {
	String() string
	GetKind() TokenKind
	GetNext() Token
}
