package param

type File struct {
	Directories [][]string
	Filenames   []string
	Extensions  []string
}

func NewFile() *File {
	return &File{
		Directories: [][]string{{
			".",
		}},
		Extensions: []string{
			".*",
		},
	}
}
