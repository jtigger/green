package plants

type Tree struct {
	Name string
	Height int
}

func NewTree(name string) Tree {
	return Tree{Name: name}
}

