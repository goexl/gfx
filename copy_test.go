package gfx_test

import (
	`testing`

	`github.com/goexl/gfx`
)

func TestCopy(t *testing.T) {
	if err := gfx.Copy(`testdata/copy`, `testdata/test`); nil != err {
		t.FailNow()
	}
}
