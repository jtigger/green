package plants

type Tree struct {
	Name string
}

func NewTree(name string) Tree {
	return Tree{Name: name}
}

