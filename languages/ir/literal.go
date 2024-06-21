package ir

type LiteralType interface {
	int | float64 | string
}

type Literal[T LiteralType] struct {
	Data T
}

func (l *Literal[T]) GetData() T {
	return l.Data
}

func (l *Literal[T]) String() string {
	return ""
}
