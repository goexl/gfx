package gfx

type copyParams struct {
	from string
	to string
}

func newCopyParams(from string, to string) *copyParams {
	return &copyParams{
		from: from,
		to:to,
	}
}
