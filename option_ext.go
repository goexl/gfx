package gfx

var _ nameOption = (*optionExt)(nil)

type optionExt struct {
	ext string
}

// Ext 配置文件扩展名
func Ext(ext string) *optionExt {
	return &optionExt{
		ext: ext,
	}
}

func (e *optionExt) applyName(options *nameOptions) {
	options.ext = e.ext
}
