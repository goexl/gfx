package param

import (
	"github.com/goexl/gfx/internal/internal/kernel"
)

type Exists struct {
	*File

	Type kernel.CheckType
}

func NewExists() *Exists {
	return &Exists{
		File: NewFile(),

		Type: kernel.CheckTypeAny,
	}
}
