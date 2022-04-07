package gfx

var (
	_           = Source
	_ zipOption = (*optionSource)(nil)
)

type optionSource struct {
	source string
}

// Source 源文件
func Source(source string) *optionSource {
	return &optionSource{
		source: source,
	}
}

func (s *optionSource) applyZip(options *zipOptions) {
	options.sources = append(options.sources, s.source)
}
