package plants

import "fmt"

type Arbor struct {
	Name string
	Height int
	Width int
}

func (t *Arbor) String() string {
	return fmt.Sprintf("Arbor \"%s\" standing %d inches tall and %d wide.", t.Name, t.Height, t.Width)
}

func NewArbor(name string, height int) Arbor {
	return Arbor{Name: name, Height: height, Width: 12}
}

