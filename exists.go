package gfx

import (
	`fmt`
	`os`
)

var _ = Exists

// Exists 判断文件是否存在
func Exists(path string, opts ...existsOption) (final string, exists bool) {
	_options := defaultExistsOptions()
	for _, opt := range opts {
		opt.applyExists(_options)
	}
	_options.paths = append(_options.paths, path)

	// 检查路径
	for _, _path := range _options.paths {

	}

	return
}

func existsWithExtensions(path string, typ checkType, extensions ...string) (final string, exists bool) {
	exists = true
	if 0 == len(extensions) {
		final = path
		exists = existsWithPath(path)
	} else {
		for _, ext := range extensions {
			path = fmt.Sprintf(filepathFormatter, path, ext)
			if existsWithPath(path) {
				if CheckTypeAny == typ {
					break
				}
			} else {

			}
		}
	}

	return
}

func existsWithPath(path string) (exists bool) {
	if _, err := os.Stat(path); nil != err && os.IsNotExist(err) {
		exists = false
	} else {
		exists = true
	}

	return
}
