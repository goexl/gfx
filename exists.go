package gfx

import (
	"os"
	"path/filepath"
)

var _ = Exists

type exists struct {
	params *existsParams
}

func Exists(path string) *existsBuilder {
	return newExistsBuilder(path)
}

func newExists(params *existsParams) *exists {
	return &exists{
		params: params,
	}
}

func (e *exists) Check() (final string, exists bool) {
	if 0==len(e.params.filenames){
		e.params.filenames=append(e.params.filenames, empty)
	}
		if 0==len(e.params.extensions){
		e.params.extensions=append(e.params.extensions, empty)
	}
	// 检查路径
	exists = true
	for _, dir := range e.params.dirs {
		for _, filename := range e.params.filenames {
			for _, ext := range e.params.extensions {
				final = filepath.Join(dir, filename, ext)
				exists = e.exists(final)
				if checkTypeAny == e.params.typ && exists || checkTypeAll == e.params.typ && !exists {
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

func (e *exists) exists(path string) (exists bool) {
	if _, err := os.Stat(path); nil != err && os.IsNotExist(err) {
		exists = false
	} else {
		exists = true
	}

	return
}
