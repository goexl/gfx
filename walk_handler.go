package gfx

import (
	`os`
)

type walkHandler func(path string, info os.FileInfo)
