package gfx

var (
	_            = Pattern
	_ walkOption = (*optionPattern)(nil)
)

type optionPattern struct {
	pattern string
}

// Pattern 配置文件匹配模式
func Pattern(pattern string) *optionPattern {
	return &optionPattern{
		pattern: pattern,
	}
}

func (p *optionPattern) applyWalk(options *walkOptions) {
	options.pattern = p.pattern
}
