package gfx

var (
	_              = Path
	_              = Paths
	_ existsOption = (*optionPaths)(nil)
)

type optionPaths struct {
	paths []string
}

// Path 路径
func Path(path string) *optionPaths {
	return &optionPaths{
		paths: []string{path},
	}
}

// Paths 路径列表
func Paths(paths ...string) *optionPaths {
	return &optionPaths{
		paths: paths,
	}
}

func (p *optionPaths) applyExists(options *existsOptions) {
	options.paths = append(options.paths, p.paths...)
}
