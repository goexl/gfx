package gfx_test

import (
	"testing"

	"github.com/goexl/gfx"
)

type (
	in struct {
		path       string
		paths      []string
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
		path:       `testdata/exists/application.yml`,
		paths:      make([]string, 0),
		extensions: make([]string, 0),
	}, expected: expected{
		final:  `testdata/exists/application.yml`,
		exists: true,
	}}, {in: in{
		path:       `testdata/exists/application.test`,
		paths:      []string{`testdata/exists/application`},
		extensions: []string{`yaml`, `yml`, `xml`, `json`, `toml`},
	}, expected: expected{
		final:  `testdata/exists/application.yaml`,
		exists: true,
	}}, {in: in{
		path:       `testdata/exists/application.test`,
		paths:      []string{`testdata/exists/application`},
		extensions: []string{`toml`, `yaml`, `yml`, `xml`, `json`},
	}, expected: expected{
		final:  `testdata/exists/application.toml`,
		exists: true,
	}}, {in: in{
		path:       `testdata/exists/application.test`,
		paths:      []string{`testdata/exists/application/application`},
		extensions: []string{`toml`, `yaml`, `yml`, `xml`, `json`},
	}, expected: expected{
		final:  `testdata/exists/application/application.toml`,
		exists: true,
	}}, {in: in{
		path:       `testdata/exists/application.test`,
		paths:      []string{`testdata/exists/application`, `testdata/exists/application/application`},
		extensions: []string{`gfx`, `gex`},
	}, expected: expected{
		final:  `testdata/exists/application/application.toml`,
		exists: false,
	}}}

	for _, test := range tests {
		final, exists := gfx.Exists(test.in.path, gfx.Paths(test.in.paths...), gfx.Extensions(test.in.extensions...))
		if true == exists && true == test.expected.exists {
			if final != test.expected.final {
				t.Fatalf(
					`期望：{final=%v, exists=%v}，实际：{final=%v, exist=%v}`,
					test.expected.final, test.expected.exists,
					final, exists,
				)
			}
		}
	}
}
