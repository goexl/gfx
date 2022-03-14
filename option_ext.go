package gfx

var (
	_            = Ext
	_ nameOption = (*optionExt)(nil)
)

type optionExt struct {
	extensions []string
}

// Ext 配置文件扩展名
func Ext(ext string, others ...string) *optionExt {
	return &optionExt{
		extensions: append([]string{
			ext,
		}, others...),
	}
}

func (e *optionExt) applyName(options *nameOptions) {
	options.ext = e.extensions[0]
}

func (e *optionExt) applyExists(options *existsOptions) {
	options.extensions = append(options.extensions, e.extensions...)
}
