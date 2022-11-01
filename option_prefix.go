package gfx

import (
	"strings"
)

var (
	_           = Prefix
	_ zipOption = (*optionPrefix)(nil)
	_ walkOption = (*optionPrefix)(nil)
)

type optionPrefix struct {
	prefix string
}

// Prefix 前缀
func Prefix(prefix string) *optionPrefix {
	return &optionPrefix{
		prefix: prefix,
	}
}

func (p *optionPrefix) applyZip(options *zipOptions) {
	options.prefix = p.prefix
}

func (p *optionPrefix) applyWalk(options *walkOptions) {
	options.matchable = func(path string) (matched bool, err error) {
		matched = strings.HasSuffix(path, p.prefix)

		return
	}
}
