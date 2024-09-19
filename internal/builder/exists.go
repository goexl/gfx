package builder

import (
	"github.com/goexl/gfx/internal/internal/core"
	"github.com/goexl/gfx/internal/internal/kernel"
	"github.com/goexl/gfx/internal/internal/param"
)

type Exists struct {
	*file[Exists]

	params *param.Exists
}

func NewExists() (exists *Exists) {
	exists = new(Exists)
	params := param.NewExists()

	exists.file = newFile(params.File, exists)
	exists.params = params

	return
}

func (e *Exists) All() (exists *Exists) {
	e.params.Type = kernel.CheckTypeAll
	exists = e

	return
}

func (e *Exists) Any() (exists *Exists) {
	e.params.Type = kernel.CheckTypeAny
	exists = e

	return
}

func (e *Exists) Build() (exists *core.Exists) {
	e.file.build()
	exists = core.NewExists(e.params)

	return
}
