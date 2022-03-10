package gfx

var (
	_            = Type
	_            = File
	_            = Dir
	_ option     = (*optionType)(nil)
	_ nameOption = (*optionType)(nil)
)

type optionType struct {
	typ _type
}

// Type 配置类型
func Type(_type _type) *optionType {
	return &optionType{
		typ: _type,
	}
}

// File 配置类型为文件
func File() *optionType {
	return Type(TypeFile)
}

func Dir() *optionType {
	return Type(TypeDir)
}

func (t *optionType) apply(options *options) {
	options.typ = t.typ
}

func (t *optionType) applyName(options *nameOptions) {
	options.typ = t.typ
}

func (t *optionType) applyCheck(options *isOptions) {
	options.typ = t.typ
}
