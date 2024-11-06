package param

import (
	"github.com/goexl/gfx/internal/internal/kernel"
)

type Limit struct {
	Type kernel.LimitType
}

func NewLimit() *Limit {
	return &Limit{
		Type: kernel.LimitTypeAll,
	}
}
