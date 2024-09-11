package kernel

const (
	CheckTypeAll CheckType = iota + 1
	CheckTypeAny
)

type CheckType uint8
