package builder

import (
	"github.com/goexl/gfx/internal/internal/kernel"
	"github.com/goexl/gfx/internal/internal/param"
)

type limit[F any] struct {
	from *F

	current *param.Limit
	params  *param.Limit
}

func newLimit[F any](from *F, params *param.Limit) *limit[F] {
	return &limit[F]{
		from: from,

		current: param.NewLimit(),
		params:  params,
	}
}

func (l *limit[F]) Dir() *limit[F] {
	return l.Directory()
}

func (l *limit[F]) Directory() (limit *limit[F]) {
	l.current.Type = kernel.LimitTypeDirectory
	limit = l

	return
}

func (l *limit[F]) File() (limit *limit[F]) {
	l.current.Type = kernel.LimitTypeFile
	limit = l

	return
}

func (l *limit[F]) All() (limit *limit[F]) {
	l.current.Type = kernel.LimitTypeAll
	limit = l

	return
}

func (l *limit[F]) Build() (from *F) {
	*l.params = *l.current
	from = l.from

	return
}
