package gfx

import (
	"os"
	"sync"

	"github.com/fsnotify/fsnotify"
)

var (
	_      = Watch
	once   sync.Once
	notify *fsnotify.Watcher
)

func Watch(path string, watcher Watcher) (err error) {
	if nil == notify {
		notify, err = fsnotify.NewWatcher()
	}
	if nil != err {
		return
	}

	// 只能被调用一次
	once.Do(func() {
		go watch(watcher)
	})
	err = notify.Add(path)

	return
}

func watch(watcher Watcher) {
	for {
		select {
		case event, ok := <-notify.Events:
			onEvent(watcher, event, ok)
		case err, ok := <-notify.Errors:
			onError(watcher, err, ok)
		}
	}
}

func onEvent(watcher Watcher, event fsnotify.Event, ok bool) {
	if !ok {
		return
	}

	path := event.Name
	switch event.Op {
	case fsnotify.Write:
		if data, err := os.ReadFile(path); nil == err {
			watcher.OnChanged(path, data)
		}
	case fsnotify.Remove:
		watcher.OnDeleted(path)
	case fsnotify.Rename:
		watcher.OnRenamed(path)
	case fsnotify.Create:
		watcher.OnCreated(path)
	case fsnotify.Chmod:
		watcher.OnPermissionChanged(path)
	}
}

func onError(watcher Watcher, err error, ok bool) {
	if !ok {
		return
	}

	watcher.OnError(err)
}
