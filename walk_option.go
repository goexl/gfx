package gfx

var _ = NewWalkOptions

type (
	walkOption interface {
		applyWalk(options *walkOptions)
	}

	walkOptions struct {
		typ       _type
		matchable matchable
	}
)

// NewWalkOptions 暴露给外部使用的快捷方法
func NewWalkOptions(opts ...walkOption) []walkOption {
	return opts
}

func defaultWalkOptions() *walkOptions {
	return &walkOptions{
		typ:       TypeFile,
		matchable: patternMatchable(all),
	}
}
