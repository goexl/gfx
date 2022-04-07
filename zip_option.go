package gfx

type (
	zipOption interface {
		applyZip(option *zipOptions)
	}

	zipOptions struct {
		prefix  string
		sources []string
	}
)

func defaultZipOptions(source string) *zipOptions {
	return &zipOptions{
		prefix: ``,
		sources: []string{
			source,
		},
	}
}
