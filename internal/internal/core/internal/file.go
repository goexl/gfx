package internal

import (
	"path/filepath"
	"strings"

	"github.com/goexl/gfx/internal/internal/param"
)

type File struct {
	params *param.File
}

func NewFile(params *param.File) *File {
	return &File{
		params: params,
	}
}

func (f *File) Patterns(dir string) (patterns []string) {
	if 0 == len(f.params.Filenames) {
		patterns = []string{dir}
	} else {
		patterns = f.patterns(dir)
	}

	return
}

func (f *File) patterns(dir string) (patterns []string) {
	patterns = make([]string, 0, len(f.params.Filenames)*len(f.params.Extensions))
	for _, filename := range f.params.Filenames {
		for _, extension := range f.params.Extensions {
			name := new(strings.Builder)
			name.WriteString(filename)
			name.WriteString(extension)
			patterns = append(patterns, filepath.Join(dir, name.String()))
		}
	}

	return
}
