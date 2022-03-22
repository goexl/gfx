package gfx

import (
	`regexp`
)

var (
	_            = Regex
	_ walkOption = (*optionRegex)(nil)
)

type optionRegex struct {
	regex string
}

// Regex 正则匹配
func Regex(regex string) *optionRegex {
	return &optionRegex{
		regex: regex,
	}
}

func (r *optionRegex) applyWalk(options *walkOptions) {
	options.matchable = func(path string) (matched bool, err error) {
		return regexp.MatchString(r.regex, path)
	}
}
