package gfx

var (
	_            = Pattern
	_ walkOption = (*optionPattern)(nil)
)

type optionPattern struct {
	pattern string
}

// Pattern 模式匹配
func Pattern(pattern string) *optionPattern {
	return &optionPattern{
		pattern: pattern,
	}
}

func (p *optionPattern) applyWalk(options *walkOptions) {
	options.matchable = patternMatchable(p.pattern)
}
