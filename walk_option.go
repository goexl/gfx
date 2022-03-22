package gfx

type (
	walkOption interface {
		applyWalk(options *walkOptions)
	}

	walkOptions struct {
		typ       _type
		matchable matchable
	}
)

func defaultWalkOptions() *walkOptions {
	return &walkOptions{
		typ:       TypeFile,
		matchable: patternMatchable(all),
	}
}
