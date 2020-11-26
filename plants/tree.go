package plants

import "fmt"

type Tree struct {
	Name string
	Height int
	Width int
}

func (t *Tree) String() string {
	return fmt.Sprintf("Tree \"%s\" standing %d inches tall and %d wide.", t.Name, t.Height, t.Width)
}

func NewTree(name string, height int) Tree {
	return Tree{Name: name, Height: height, Width: 12}
}

