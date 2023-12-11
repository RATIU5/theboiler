package files

type FileContent struct {
	Path    string
	IsDir   bool
	Content []string
}

func (fc FileContent) String() string {
	str := "------------------\n"
	if fc.IsDir {
		str += fc.Path + "\n"
	} else {
		str += fc.Path + ":\n"
	}
	str += "------------------\n"
	for _, content := range fc.Content {
		str += content + "\n"
	}
	return str
}
