package gfx

type (
	isOption interface {
		applyCheck(options *isOptions)
	}

	isOptions struct {
		typ _type
	}
)

func defaultIsOptions() *isOptions {
	return &isOptions{
		typ: TypeDir,
	}
}
