package gfx_test

import (
	"path/filepath"
	"testing"

	"github.com/goexl/gfx"
)

type (
	in struct {
		dirs       [][]string
		filename   string
		extensions []string
	}

	expected struct {
		final  string
		exists bool
	}
)

func TestExists(t *testing.T) {
	tests := []struct {
		in       in
		expected expected
	}{{in: in{
		dirs: [][]string{{
			"testdata", "exists", "application",
		}},
		filename:   "application",
		extensions: []string{"yaml", "yml", "xml", "json", "toml"},
	}, expected: expected{
		final:  filepath.Clean("testdata/exists/application/application.yaml"),
		exists: true,
	}}, {in: in{
		dirs: [][]string{{
			"testdata/exists/application",
		}},
		filename:   "application",
		extensions: []string{"toml", "yaml", "yml", "xml", "json"},
	}, expected: expected{
		final:  filepath.Clean("testdata/exists/application/application.toml"),
		exists: true,
	}}, {in: in{
		dirs: [][]string{{
			"testdata", "exists", "application", "application",
		}},
		filename:   "application",
		extensions: []string{"toml", "yaml", "yml", "xml", "json"},
	}, expected: expected{
		exists: false,
	}}, {in: in{
		dirs: [][]string{{
			"testdata", "exists", "application",
		}, {
			"testdata", "exists", "application", "application",
		}},
		filename:   "application",
		extensions: []string{"gfx", "gex"},
	}, expected: expected{
		exists: false,
	}}}

	for index, test := range tests {
		exists := gfx.Exists().Filename(test.in.filename).Extension(test.in.extensions[0], test.in.extensions[1:]...)
		for _, dir := range test.in.dirs {
			exists.Directory(dir[0], dir[1:]...)
		}
		final, checked := exists.Build().Check()
		if true == checked && true == test.expected.exists {
			if final != test.expected.final {
				t.Fatalf(
					"第%d个测试不通过，期望：{final=%v, checked=%v}，实际：{final=%v, exist=%v}",
					index+1,
					test.expected.final, test.expected.exists,
					final, checked,
				)
			}
		}
	}
}
