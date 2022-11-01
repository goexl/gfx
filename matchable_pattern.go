package gfx

import (
	"path"
)

func patternMatchable(pattern string) matchable {
	return func(filepath string) (matched bool, err error) {
		return path.Match(pattern, filepath)
	}
}
