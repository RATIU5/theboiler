package item

import "fmt"

type Group struct {
	Items []Item
	Name  string
}

func (i *Group) Print() {
	fmt.Printf("%s:\n", i.Name)
	for _, itm := range i.Items {
		fmt.Print("  ")
		itm.Print()
	}
}
