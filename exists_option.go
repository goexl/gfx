package gfx

var _ = NewExistsOptions

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

// NewExistsOptions 快捷方式，解决接口不对外暴露不能引用的问题
func NewExistsOptions(opts ...existsOption) []existsOption {
	return opts
}

func defaultExistsOptions() *existsOptions {
	return &existsOptions{
		paths: make([]string, 0),
		typ:   CheckTypeAny,
	}
}
