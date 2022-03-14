package gfx

import (
	`strings`
)

var (
	_            = Ext
	_ nameOption = (*optionExt)(nil)
)

type optionExt struct {
	extensions []string
}

// Ext 配置文件扩展名
func Ext(ext string, others ...string) *optionExt {
	extensions := make([]string, len(others)+1)
	extensions = append(extensions, convertExt(ext))
	for _, other := range others {
		extensions = append(extensions, convertExt(other))
	}

	return &optionExt{
		extensions: extensions,
	}
}

func (e *optionExt) applyName(options *nameOptions) {
	options.ext = e.extensions[0]
}

func (e *optionExt) applyExists(options *existsOptions) {
	options.extensions = append(options.extensions, e.extensions...)
}

func convertExt(from string) (to string) {
	to = from
	if strings.HasPrefix(to, dot) {
		to = to[1:]
	}

	return
}
