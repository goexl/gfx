package gfx

type copyBuilder struct {
	params *copyParams
}

func newCopyBuilder(from string,to string) *copyBuilder {
	return &copyBuilder{
		params: newCopyParams(from, to),
	}
}
