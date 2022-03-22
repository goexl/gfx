package gfx

import (
	`regexp`
)

var (
	_            = Regexp
	_ walkOption = (*optionRegexp)(nil)
)

type optionRegexp struct {
	regexp string
}

// Regexp 正则匹配
func Regexp(regex string) *optionRegexp {
	return &optionRegexp{
		regexp: regex,
	}
}

func (r *optionRegexp) applyWalk(options *walkOptions) {
	options.matchable = func(path string) (matched bool, err error) {
		return regexp.MatchString(r.regexp, path)
	}
}
