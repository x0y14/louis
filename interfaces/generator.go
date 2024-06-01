package interfaces

type Generator interface {
	Generate(n Node) (string, error)
}
