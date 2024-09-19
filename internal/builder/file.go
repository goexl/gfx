package builder

import (
	"path/filepath"
	"strings"

	"github.com/goexl/gfx/internal/internal/constant"
	"github.com/goexl/gfx/internal/internal/param"
)

type file[T any] struct {
	params *param.File
	from   *T
}

func newFile[T any](params *param.File, from *T) *file[T] {
	return &file[T]{
		params: params,
		from:   from,
	}
}

func (f *file[T]) Dir(dir string, dirs ...string) *T {
	return f.Directory(dir, dirs...)
}

func (f *file[T]) Directory(directory string, directories ...string) (t *T) {
	f.params.Directories = append(f.params.Directories, append([]string{directory}, directories...))
	t = f.from

	return
}

func (f *file[T]) Filepath(required string, paths ...string) (t *T) {
	for _, path := range append([]string{required}, paths...) {
		if "" == path {
			continue
		}

		dir, filename := filepath.Split(path)
		name := filepath.Base(filename)
		ext := filepath.Ext(name)
		f.params.Directories = append(f.params.Directories, filepath.SplitList(dir))
		f.params.Filenames = append(f.params.Filenames, name)
		f.params.Extensions = append(f.params.Extensions, ext)
	}
	t = f.from

	return
}

func (f *file[T]) Filename(filename string, filenames ...string) (t *T) {
	f.params.Filenames = append(f.params.Filenames, filename)
	f.params.Filenames = append(f.params.Filenames, filenames...)
	t = f.from

	return
}

func (f *file[T]) Extension(extension string, extensions ...string) (t *T) {
	f.extension(append([]string{extension}, extensions...))
	t = f.from

	return
}

func (f *file[T]) build() {
	// 检查扩展名是不是已经被设置过，如果被设置过去除默认配置
	if 1 < len(f.params.Extensions) {
		f.params.Extensions = f.params.Extensions[1:]
	}
	// 检查目录是不是已经被设置过，如果被设置过去除默认配置
	if 1 < len(f.params.Directories) {
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
