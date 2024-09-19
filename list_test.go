package gfx_test

import (
	"path/filepath"
	"reflect"
	"testing"

	"github.com/goexl/gfx"
)

func TestList(t *testing.T) {
	tests := []struct {
		in       in
		expected []string
	}{{in: in{
		dirs: [][]string{{
			"testdata", "exists", "application",
		}},
		filename:   "application",
		extensions: []string{"yaml", "yml", "xml", "json", "toml"},
	}, expected: []string{
		filepath.Clean("testdata/exists/application/application.yaml"),
		filepath.Clean("testdata/exists/application/application.yml"),
		filepath.Clean("testdata/exists/application/application.xml"),
		filepath.Clean("testdata/exists/application/application.json"),
		filepath.Clean("testdata/exists/application/application.toml"),
	}}, {in: in{
		dirs: [][]string{{
			"testdata/exists/application",
		}},
		filename:   "application",
		extensions: []string{"toml", "yaml", "xml", "json"},
	}, expected: []string{
		filepath.Clean("testdata/exists/application/application.toml"),
		filepath.Clean("testdata/exists/application/application.yaml"),
		filepath.Clean("testdata/exists/application/application.xml"),
		filepath.Clean("testdata/exists/application/application.json"),
	}}, {in: in{
		dirs: [][]string{{
			"testdata", "exists", "application", "application",
		}},
		filename:   "application",
		extensions: []string{"toml", "yaml", "yml", "xml", "json"},
	}, expected: []string{
		// 无
	}}, {in: in{
		dirs: [][]string{{
			"testdata", "exists", "application",
		}, {
			"testdata", "exists", "application", "application",
		}},
		filename:   "application",
		extensions: []string{"gfx", "gex"},
	}, expected: []string{
		// 无
	}}}

	for index, test := range tests {
		list := gfx.List().Filename(test.in.filename).Extension(test.in.extensions[0], test.in.extensions[1:]...)
		for _, dir := range test.in.dirs {
			list.Directory(dir[0], dir[1:]...)
		}
		files := list.Build().All()
		if !reflect.DeepEqual(files, test.expected) {
			t.Fatalf("第%d个测试出错，期望：%v，实际：%v", index+1, test.expected, files)
		}
	}
}
