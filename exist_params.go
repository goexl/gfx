package gfx

import (
	"path/filepath"
	"strings"
)

type existsParams struct {
	dirs       []string
	filenames  []string
	extensions []string
	typ        checkType
}

func newExistsParams(paths []string) (params *existsParams) {
	params = new(existsParams)
	params.dirs=make([]string, 0, len(paths))
		params.filenames=make([]string, 0, len(paths))
			params.extensions=make([]string, 0, len(paths))
	params.typ = checkTypeAny

	for _,path:=range paths{
	dir, file := filepath.Split(path)
	index := strings.Index(file, dot)
	params.dirs = append(params.dirs, dir)
	params.filenames=append(params.filenames, file[:index])
	params.extensions = append(params.extensions, file[index:])
	}

	return
}
