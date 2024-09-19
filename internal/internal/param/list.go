package param

type List struct {
	*File
}

func NewList() *List {
	return &List{
		File: NewFile(),
	}
}
