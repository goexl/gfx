package core

import (
	"path/filepath"

	"github.com/goexl/gfx/internal/internal/kernel"
	"github.com/goexl/gfx/internal/internal/param"
)

type Exists struct {
	params *param.Exists
}

func NewExists(params *param.Exists) *Exists {
	return &Exists{
		params: params,
	}
}

func (e *Exists) Check() (final string, exists bool) {
	// 检查路径
	exists = true
	for _, directories := range e.params.Directories {
		directory := filepath.Join(directories...)
		for _, filename := range e.params.Filenames {
			for _, extension := range e.params.Extensions {
				pattern := filepath.Join(directory, filename, extension)
				final, exists = e.check(pattern)
				if kernel.CheckTypeAny == e.params.Type && exists || kernel.CheckTypeAll == e.params.Type && !exists {
					break
				}
			}
		}
	}

	// 如果不存在，需要清空最终路径
	if !exists {
		final = ""
	}

	return
}

func (e *Exists) check(pattern string) (final string, exists bool) {
	if files, _ := filepath.Glob(pattern); 0 != len(files) {
		final = files[0]
		exists = true
	}

	return
}
