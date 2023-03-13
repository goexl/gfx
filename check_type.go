package gfx

const (
	checkTypeUnknown checkType = iota
	checkTypeAll
	checkTypeAny
)

var _ = checkTypeUnknown

type checkType uint8
