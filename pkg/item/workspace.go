package item

import "fmt"

type Workspace struct {
	Groups []Group
	Name   string
}

func (w *Workspace) Print() {
	fmt.Printf("%s:\n", w.Name)
	for _, grp := range w.Groups {
		fmt.Printf("  %s", grp.Name)
	}
}
