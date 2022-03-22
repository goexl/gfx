package gfx

import (
	`path/filepath`
)

func patternMatchable(pattern string) matchable {
	return func(path string) (matched bool, err error) {
		return filepath.Match(pattern, path)
	}
}
