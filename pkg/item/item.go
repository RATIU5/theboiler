package item

import "fmt"

type Item struct {
	isDir     bool
	path      string
	value     string
	workspace string
}

func New(isDir bool, path string, value string) *Item {
	return &Item{isDir: isDir, path: path, value: value}
}

func (i *Item) Print() {
	if i.isDir {
		fmt.Printf("%s/\n", i.path)
	} else {
		fmt.Printf("%s: %s\n", i.path, i.value)
	}
}
