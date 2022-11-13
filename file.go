package gfx

import (
	`os`
	"path/filepath"
	`syscall`
)

var (
	_ = Create
	_ = Rename
	_=Move
	_ = Delete
)

// Create 创建文件或者目录
// 默认创建文件
// 默认权限是0777
func Create(path string, opts ...option) (err error) {
	_options := defaultOptions()
	for _, opt := range opts {
		opt.apply(_options)
	}

	if existsWithPath(path) {
		switch _options.writeMode {
		case WriteModeError:
			err = errFileExists
		case WriteModeSkip:
			return
		case WriteModeOverride:
			err = Delete(path)
		case WriteModeRename:
			path = NewName(path)
		}
	}
	if nil != err {
		return
	}

	// 创建文件或者目录
	switch _options.typ {
	case TypeDir:
		err = os.MkdirAll(path, _options.fileMode)
	default:
		_, err = os.Create(path)
	}
	if nil != err {
		return
	}

	// 改变文件的拥有者
	if nil != _options.owner {
		err = os.Chown(path, _options.owner.uid, _options.owner.gid)
	}

	return
}

// Rename 重命名文件或者目录
func Rename(from string, to string) (err error) {
	if !existsWithPath(filepath.Dir(to)) {
		err = os.MkdirAll(to, os.ModePerm)
	} else {
		err = syscall.Rename(from, to)
	}

	return
}

// Move 移动文件
func Move(from string, to string) error{
	return Rename(from,to)
}

// Delete 删除文件或者目录
func Delete(filename string) error {
	return os.RemoveAll(filename)
}

// Is 判断所给路径是否为文件或者目录
func Is(path string, opts ...isOption) (is bool, err error) {
	_options := defaultIsOptions()
	for _, opt := range opts {
		opt.applyCheck(_options)
	}

	var stat os.FileInfo
	if stat, err = os.Stat(path); nil != err {
		return
	}

	switch _options.typ {
	case TypeDir:
		is = stat.IsDir()
	case TypeFile:
		is = !stat.IsDir()
	default:
		is = true
	}
	stat = nil

	return
}
