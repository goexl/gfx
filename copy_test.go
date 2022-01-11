package gfx_test

import (
	`testing`

	`github.com/storezhang/gox/file`
)

func TestCopy(t *testing.T) {
	if err := file.Copy(`../testdata/copy`, `../testdata/test`); nil != err {
		t.FailNow()
	}
}
