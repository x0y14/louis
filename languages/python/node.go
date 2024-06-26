package python

import "louis/interfaces"

type NodeKind int

func (n NodeKind) String() string {
	return ""
}

type Node struct {
	Kind NodeKind
}

func (n *Node) String() string {
	return ""
}

func (n *Node) GetKind() interfaces.NodeKind {
	return n.Kind
}
