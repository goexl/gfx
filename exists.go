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
	// 默认的路径必须在最前
	_options.paths = append([]string{path}, _options.paths...)

	// 检查路径
	exists = true
	for _, _path := range _options.paths {
		typ := _options.typ
		final, exists = existsWithExtensions(_path, typ, _options.extensions...)
		if CheckTypeAny == typ && exists || CheckTypeAll == typ && !exists {
			break
		}
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
			final = fmt.Sprintf(filepathFormatter, path, ext)
			exists = existsWithPath(final)

			if CheckTypeAny == typ && exists || CheckTypeAll == typ && !exists {
				break
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
