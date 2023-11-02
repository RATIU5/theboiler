package item

import "fmt"

type Item struct {
	IsDir bool
	Path  string
	Value string
}

func New(isDir bool, path string, value string, name string) *Item {
	return &Item{IsDir: isDir, Path: path, Value: value, Name: name}
}

func (i *Item) Print() {
	fmt.Printf("%s:\n", i.Name)
	if i.IsDir {
		fmt.Printf("dir: %s\n", i.Path)
	} else {
		fmt.Printf("file: %s\n", i.Path)
	}
}
