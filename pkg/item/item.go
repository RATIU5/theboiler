package item

import "fmt"

type Item struct {
	IsDir bool
	Path  string
	Value string
}

func New(isDir bool, path string, value string) *Item {
	return &Item{IsDir: isDir, Path: path, Value: value}
}

func (i *Item) Print() {
	if i.IsDir {
		fmt.Printf("dir: %s\n", i.Path)
	} else {
		fmt.Printf("file: %s\n", i.Path)
	}
}
