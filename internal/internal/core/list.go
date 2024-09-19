package core

import (
	"path/filepath"

	"github.com/goexl/gfx/internal/internal/core/internal"
	"github.com/goexl/gfx/internal/internal/param"
)

type List struct {
	params *param.List
	file   *internal.File
}

func NewList(params *param.List) *List {
	return &List{
		params: params,
		file:   internal.NewFile(params.File),
	}
}

func (l *List) All() (files []string) {
	files = make([]string, 0, 8)
	for _, directories := range l.params.Directories {
		directory := filepath.Join(directories...)
		patterns := l.file.Patterns(directory)
		for _, pattern := range patterns {
			if all, ge := filepath.Glob(pattern); nil == ge && 0 != len(all) {
				files = append(files, all...)
			}
		}
	}

	return
}

func (l *List) check(pattern string) (final string, exists bool) {
	if files, _ := filepath.Glob(pattern); 0 != len(files) {
		final = files[0]
		exists = true
	}

	return
}
