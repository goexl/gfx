package gfx

import (
	`os`
	`path/filepath`
)

var (
	_ = All
	_ = Walk
)

// All 取得目录下所有文件，包含子目录
// 默认文件匹配所有文件
func All(dir string, opts ...walkOption) (filenames []string, err error) {
	filenames = make([]string, 0)
	err = Walk(dir, func(path string, info os.FileInfo) {
		filenames = append(filenames, path)
	}, opts...)

	return
}

// Walk 浏览目录
// 默认文件匹配所有文件
func Walk(dir string, handler walkHandler, opts ...walkOption) (err error) {
	_options := defaultWalkOptions()
	for _, opt := range opts {
		opt.applyWalk(_options)
	}

	err = filepath.Walk(dir, func(path string, info os.FileInfo, walkErr error) (err error) {
		if err = walkErr; nil != err {
			return
		}
		if TypeDir == _options.typ && !info.IsDir() || TypeFile == _options.typ && info.IsDir() {
			return
		}

		if nil==_options.matchable{
			handler(path,info)
		}else if matched, me := _options.matchable(path); me != nil {
			err = me
		} else if matched {
			handler(path, info)
		}

		return
	})

	return
}
