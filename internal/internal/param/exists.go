package param

import (
	"github.com/goexl/gfx/internal/internal/kernel"
)

type Exists struct {
	Directories [][]string
	Filenames   []string
	Extensions  []string
	Type        kernel.CheckType
}

func NewExists() *Exists {
	return &Exists{
		Directories: [][]string{{
			".",
		}},
		Extensions: []string{
			".*",
		},
		Type: kernel.CheckTypeAny,
	}
}
