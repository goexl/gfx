package builder

import (
	"strings"

	"github.com/goexl/gfx/internal/internal/constant"
	"github.com/goexl/gfx/internal/internal/core"
	"github.com/goexl/gfx/internal/internal/kernel"
	"github.com/goexl/gfx/internal/internal/param"
)

type Exists struct {
	params *param.Exists
}

func NewExists() (exists *Exists) {
	return &Exists{
		params: param.NewExists(),
	}
}

func (e *Exists) All() (exists *Exists) {
	e.params.Type = kernel.CheckTypeAll
	exists = e

	return
}

func (e *Exists) Any() (exists *Exists) {
	e.params.Type = kernel.CheckTypeAny

	return e
}

func (e *Exists) Dir(dir string, dirs ...string) *Exists {
	return e.Directory(dir, dirs...)
}

func (e *Exists) Directory(directory string, directories ...string) (exists *Exists) {
	e.params.Directories = append(e.params.Directories, append([]string{directory}, directories...))
	exists = e

	return
}

func (e *Exists) Filename(filename string, filenames ...string) (exists *Exists) {
	e.params.Filenames = append(e.params.Filenames, filename)
	e.params.Filenames = append(e.params.Filenames, filenames...)
	exists = e

	return
}

func (e *Exists) Extension(extension string, extensions ...string) (exists *Exists) {
	e.extension(append([]string{extension}, extensions...))
	exists = e

	return
}

func (e *Exists) Build() *core.Exists {
	// 检查扩展名是不是已经被设置过，如果被设置过去除默认配置
	if 1 < len(e.params.Extensions) {
		e.params.Extensions = e.params.Extensions[1:]
	}

	return core.NewExists(e.params)
}

func (e *Exists) extension(extensions []string) {
	for _, extension := range extensions {
		if !strings.HasPrefix(extension, constant.Dot) {
			builder := new(strings.Builder)
			builder.WriteString(constant.Dot)
			builder.WriteString(extension)
			extension = builder.String()
		}
		e.params.Extensions = append(e.params.Extensions, extensions...)
	}
}
