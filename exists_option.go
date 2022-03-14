package gfx

type (
	existsOption interface {
		applyExists(options *existsOptions)
	}

	existsOptions struct {
		paths      []string
		typ        checkType
		extensions []string
	}
)

func defaultExistsOptions() *existsOptions {
	return &existsOptions{
		paths: make([]string, 0),
		typ:   CheckTypeAny,
	}
}
