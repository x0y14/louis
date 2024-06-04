package interfaces

type Commonizer interface {
	Commonize(n Node) (Node, error)
}
