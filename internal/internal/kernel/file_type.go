package kernel

const (
	FileTypeFile FileType = iota + 1
	FileTypeDirectory
	FileTypeAll
)

type FileType uint8
