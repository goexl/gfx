package gfx

import (
	"strings"
)

type existsBuilder struct {
	params *existsParams
}

func newExistsBuilder(paths ...string) *existsBuilder {
	return &existsBuilder{
		params: newExistsParams(paths),
	}
}

func (eb *existsBuilder) All() *existsBuilder {
	eb.params.typ = checkTypeAll

	return eb
}

func (eb *existsBuilder) Any() *existsBuilder {
	eb.params.typ = checkTypeAny

	return eb
}

func (eb *existsBuilder) Dir(dirs ...string) *existsBuilder {
	eb.params.dirs = append(eb.params.dirs, dirs...)

	return eb
}

func (eb *existsBuilder) Filename(filenames ...string) *existsBuilder {
	eb.params.filenames = append(eb.params.filenames, filenames...)

	return eb
}

func (eb *existsBuilder) Ext(extensions ...string) *existsBuilder {
	for _, ext := range extensions {
		if !strings.HasPrefix(ext, dot) {
			builder:=new(strings.Builder)
			builder.WriteString(dot)
			builder.WriteString(ext)
			ext=builder.String()
		}
		eb.params.extensions = append(eb.params.extensions, extensions...)
	}

	return eb
}

func (eb *existsBuilder) Build() *exists {
	return newExists(eb.params)
}
