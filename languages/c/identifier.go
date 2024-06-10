package c

type Identifier struct {
	Name string
}

func (i Identifier) String() string {
	return i.Name
}
