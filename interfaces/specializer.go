package interfaces

type Specializer interface {
	Specialize(n Node) (Node, error)
}
