package gfx

var (
	_ = Path
	_ = Paths
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
