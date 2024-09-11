package gfx

import (
	"github.com/goexl/gfx/internal/builder"
)

// Exists 判断是否存在
func Exists() *builder.Exists {
	return builder.NewExists()
}
