package gfx

import (
	`fmt`
	`path/filepath`
	`strings`
)

var (
	_ = Name
	_ = NewName
)

// Name 获得文件名称
func Name(path string, opts ...nameOption) (filename string) {
	_options := defaultNameOptions()
	for _, opt := range opts {
		opt.applyName(_options)
	}

	switch _options.typ {
	case TypeDir:
		filename = dir(path)
	case TypeFile:
		fallthrough
	default:
		filename = name(path, _options.ext)
	}

	return
}

// NewName 新文件名，在避免文件名冲突的情况下
func NewName(original string) (new string) {
	for {
		index := 1
		new = filepath.Join(filepath.Dir(original), name(original, fmt.Sprintf(`%d.%s`, index, filepath.Ext(original))))
		if !existsWithPath(new) {
			break
		}
	}

	return
}

func name(path string, ext string) (filename string) {
	base := filepath.Base(path)
	_ext := filepath.Ext(path)
	_name := base[:len(base)-len(_ext)]

	if `` == strings.TrimSpace(ext) {
		filename = _name
	} else {
		filename = fmt.Sprintf(`%s.%s`, _name, ext)
	}

	return
}

func dir(path string) string {
	return filepath.Join(filepath.Dir(path), name(filepath.Base(path), ``))
}
