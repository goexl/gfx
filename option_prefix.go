package gfx

var (
	_           = Prefix
	_ zipOption = (*optionPrefix)(nil)
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
