package item

import "fmt"

type Group struct {
	Items []Item
	Name  string
}

func (i *Group) Print() {
	fmt.Printf("%s")
}
