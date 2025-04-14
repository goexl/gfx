package builder

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/goexl/gfx/internal/internal/constant"
	"github.com/goexl/gfx/internal/internal/param"
)

type file[T any] struct {
	params *param.File
	limit  *param.Limit
	from   *T

	reset bool
}

func newFile[T any](params *param.File, from *T) *file[T] {
	return &file[T]{
		params: params,
		limit:  param.NewLimit(),
		from:   from,
	}
}

func (f *file[T]) Limit() *limit[file[T]] {
	return newLimit(f, f.limit)
}

func (f *file[T]) Reset() (t *T) {
	f.reset = true
	t = f.from

	return
}

func (f *file[T]) Dir(required string, optionals ...string) *T {
	return f.Directory(required, optionals...)
}

func (f *file[T]) Directory(required string, optionals ...string) (t *T) {
	f.params.Directories = append(f.params.Directories, append([]string{required}, optionals...))
	t = f.from

	return
}

func (f *file[T]) Filepath(required string, optionals ...string) (t *T) {
	for _, path := range append([]string{required}, optionals...) {
		if "" == path {
			continue
		}

		if info, se := os.Stat(path); (nil == se && info.IsDir()) || "" == filepath.Ext(path) {
			// 如果是目录或者没有扩展名
			// !认为是目录，因为有可能配置成为一个不存在路径，为后续动态加载文件提供可能
			f.params.Directories = append(f.params.Directories, []string{path})
		} else {
			dir, filename := filepath.Split(path)
			name := filepath.Base(filename)
			ext := filepath.Ext(name)
			f.params.Directories = append(f.params.Directories, []string{dir})
			f.params.Filenames = append(f.params.Filenames, name)
			f.params.Extensions = append(f.params.Extensions, ext)
		}
	}
	t = f.from

	return
}

func (f *file[T]) Filename(required string, optionals ...string) (t *T) {
	f.params.Filenames = append(f.params.Filenames, required)
	f.params.Filenames = append(f.params.Filenames, optionals...)
	t = f.from

	return
}

func (f *file[T]) Extension(required string, optionals ...string) (t *T) {
	f.extension(append([]string{required}, optionals...))
	t = f.from

	return
}

func (f *file[T]) build() {
	// 检查是否重置，如果被设置过去除默认配置
	if f.reset {
		f.params.Extensions = f.params.Extensions[1:]
	}
	// 检查是否重置，如果被设置过去除默认配置
	if f.reset {
		f.params.Directories = f.params.Directories[1:]
	}
}

func (f *file[T]) extension(extensions []string) {
	for _, extension := range extensions {
		if !strings.HasPrefix(extension, constant.Dot) {
			builder := new(strings.Builder)
			builder.WriteString(constant.Dot)
			builder.WriteString(extension)
			extension = builder.String()
		}
		f.params.Extensions = append(f.params.Extensions, extension)
	}
}
