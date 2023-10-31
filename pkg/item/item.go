package item

import "fmt"

type Item struct {
	IsDir     bool
	Path      string
	Value     string
	Workspace string
}

func New(isDir bool, path string, value string) *Item {
	return &Item{IsDir: isDir, Path: path, Value: value}
}

func (i *Item) Print() {
	fmt.Printf("%s:\n", i.Workspace)
	if i.IsDir {
		fmt.Printf("%s/\n", i.Path)
	} else {
		fmt.Printf("%s: %s\n", i.Path, i.Value)
	}
}
