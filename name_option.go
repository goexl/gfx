package gfx

type (
	nameOption interface {
		applyName(options *nameOptions)
	}

	nameOptions struct {
		typ _type
		ext string
	}
)

func defaultNameOptions() *nameOptions {
	return &nameOptions{
		typ: TypeFile,
	}
}
