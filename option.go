package gfx

import (
	`io/fs`
	`os`
)

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

func defaultOptions() *options {
	return &options{
		typ:      TypeFile,
		fileMode: os.ModePerm,
	}
}
