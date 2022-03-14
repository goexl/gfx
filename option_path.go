package gfx

var (
	_              = Path
	_ existsOption = (*optionPath)(nil)
)

type optionPath struct {
	paths []string
}

// Path 配置路径
func Path(path string, others ...string) *optionPath {
	return &optionPath{
		paths: append([]string{
			path,
		}, others...),
	}
}

func (p *optionPath) applyExists(options *existsOptions) {
	options.paths = append(options.paths, p.paths...)
}
