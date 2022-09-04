package gfx

// Watcher 文件监控
type Watcher interface {
	// OnCreated 当文件被创建
	OnCreated(path string)

	// OnChanged 文件发生改变时
	OnChanged(path string, content []byte)

	// OnDeleted 文件被删除时
	OnDeleted(path string)

	// OnRenamed 文件被重命名时
	OnRenamed(path string)

	// OnPermissionChanged 文件权限改变时
	OnPermissionChanged(path string)

	// OnError 当发生错误时
	OnError(err error)
}
