package builder

import (
	"github.com/goexl/gfx/internal/internal/core"
	"github.com/goexl/gfx/internal/internal/param"
)

type List struct {
	*file[List]

	params *param.List
}

func NewList() (list *List) {
	list = new(List)
	params := param.NewList()

	list.file = newFile(params.File, list)
	list.params = params

	return
}

func (l *List) Build() (list *core.List) {
	l.file.build()
	list = core.NewList(l.params)

	return
}
