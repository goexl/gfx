package param

import (
	"github.com/goexl/gfx/internal/internal/kernel"
)

type File struct {
	Type        kernel.FileType
	Directories [][]string
	Filenames   []string
	Extensions  []string
}

func NewFile() *File {
	return &File{
		Type: kernel.FileTypeAll,
		Directories: [][]string{{
			".",
		}},
		Extensions: []string{
			".*",
		},
	}
}
