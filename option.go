package gfx

import (
	"io/fs"
	"os"
)

var _ = NewOptions

type (
	option interface {
		apply(options *options)
	}

	options struct {
		typ       _type
		fileMode  fs.FileMode
		writeMode writeMode
		owner     *owner
	}
)

// NewOptions 暴露给外部使用的快捷方法
func NewOptions(opts ...option) []option {
	return opts
}

func defaultOptions() *options {
	return &options{
		typ:      TypeFile,
		fileMode: os.ModePerm,
	}
}
