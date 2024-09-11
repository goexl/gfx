package gfx

import (
	"github.com/goexl/gfx/internal/builder"
)

// Exist 判断是否存在
func Exist() *builder.Exists {
	return builder.NewExists()
}
