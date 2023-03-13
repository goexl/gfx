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

func newExistsParams(path string) (params *existsParams) {
	params = new(existsParams)
	params.typ = checkTypeAny

	dir, file := filepath.Split(path)
	index := strings.Index(file, dot)
	params.dirs = []string{
		dir,
	}
	params.filenames = []string{
		file[:index],
	}
	params.extensions = []string{
		file[index:],
	}

	return
}
