package gfx

import (
	"errors"
)

var (
	errFileExists     = errors.New(`文件已存在`)
	errSourceNotfound = errors.New(`源文件不存在`)
	errDestExists     = errors.New(`目的文件已存在`)
)
