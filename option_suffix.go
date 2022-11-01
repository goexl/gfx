package gfx

import (
	"strings"
)

var (
	_            = Suffix
	_ walkOption = (*optionSuffix)(nil)
)

type optionSuffix struct {
	suffix string
}

// Suffix 以什么结束
func Suffix(suffix string) *optionSuffix {
	return &optionSuffix{
		suffix: suffix,
	}
}

func (s *optionSuffix) applyWalk(options *walkOptions) {
	options.matchable = func(path string) (matched bool, err error) {
		matched = strings.HasSuffix(path, s.suffix)

		return
	}
}
