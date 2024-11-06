package kernel

const (
	LimitTypeFile LimitType = iota + 1
	LimitTypeDirectory
	LimitTypeAll
)

type LimitType uint8
