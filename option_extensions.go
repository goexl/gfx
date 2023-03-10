package gfx

import (
	"strings"
)

var (
	_            = Ext
	_            = Extension
	_            = Extensions
	_ nameOption = (*optionExtensions)(nil)
)

type optionExtensions struct {
	extensions []string
}

// Ext 文件扩展名
func Ext(ext string) *optionExtensions {
	return Extension(ext)
}

// Extension 文件扩展名
func Extension(extension string) *optionExtensions {
	return &optionExtensions{
		extensions: []string{convertExt(extension)},
	}
}

// Extensions 文件扩展名列表
func Extensions(extensions ...string) *optionExtensions {
	converts := make([]string, 0, len(extensions))
	for _, ext := range extensions {
		converts = append(converts, convertExt(ext))
	}

	return &optionExtensions{
		extensions: converts,
	}
}

func (e *optionExtensions) applyName(options *nameOptions) {
	options.ext = e.extensions[0]
}

func (e *optionExtensions) applyExists(options *existsOptions) {
	options.extensions = append(options.extensions, e.extensions...)
}

func convertExt(from string) (to string) {
	to = from
	if strings.HasPrefix(to, dot) {
		to = to[1:]
	}

	return
}
