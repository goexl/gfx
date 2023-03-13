package gfx

import (
	"strings"
)

type existsBuilder struct {
	params *existsParams
}

func newExistsBuilder(path string) *existsBuilder {
	return &existsBuilder{
		params: newExistsParams(path),
	}
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
			// TODO
		}
		eb.params.extensions = append(eb.params.extensions, extensions...)
	}

	return eb
}

func (eb *existsBuilder) Build() *exists {
	return newExists(eb.params)
}
