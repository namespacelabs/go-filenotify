// Package filenotify provides a mechanism for watching file(s) for changes.
// Generally leans on fsnotify, but provides a poll-based notifier which fsnotify does not support.
// These are wrapped up in a common interface so that either can be used interchangeably in your code.
package filenotify

import (
	"io"

	"github.com/fsnotify/fsnotify"
)

// FileWatcher is an interface for implementing file notification watchers
type FileWatcher interface {
	Events() <-chan fsnotify.Event
	Errors() <-chan error
	Add(name string) error
	Remove(name string) error
	Close() error
}

// NewPollingWatcher returns a poll-based file watcher
func NewPollingWatcher(debug io.Writer) FileWatcher {
	return &filePoller{
		events: make(chan fsnotify.Event),
		errors: make(chan error),
		debug:  debug,
	}
}

// NewEventWatcher returns an fs-event based file watcher
func NewEventWatcher() (FileWatcher, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}
	return &fsNotifyWatcher{watcher}, nil
}
