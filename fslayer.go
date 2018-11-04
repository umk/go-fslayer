package fslayer

import (
	"os"

	"gopkg.in/src-d/go-billy.v4"
	"gopkg.in/src-d/go-billy.v4/memfs"
	"gopkg.in/src-d/go-billy.v4/osfs"
)

type (
	// Extends the file system abstraction with new methods.
	storage interface {
		billy.Filesystem
		// RemoveAll removes path and any children it contains. It removes
		// everything it can but returns the first error it encounters. If the path
		// does not exist, RemoveAll returns nil (no error).
		RemoveAll(path string) error
	}

	// A storage that keeps everything in memory.
	memoryStorage struct {
		billy.Filesystem
	}

	// A storage that uses persistent storage to storage data.
	storageDevice struct {
		billy.Filesystem
	}
)

// -----------------------------------------------------------------------------
// Storage initialization

var root storage

// Fs gets the root file system object.
func Fs() storage {
	if root == nil {
		panic("file system has not been initialized")
	}
	return root
}

// setRoot checks that root file system object has not been initialized yet, and
// then assigns provided object as root.
func setRoot(fs storage) {
	if root != nil {
		panic("file system has already been initialized")
	}
	root = fs
}

// UseMemoryStorage instructs the application to use RAM for storing the
// configuration. The initialization must be performed exactly once during
// startup.
func UseMemoryStorage() {
	setRoot(memoryStorage{Filesystem: memfs.New()})
}

// UseStorageDevice instructs the application to use storage device for storing
// the configuration. The initialization must be performed exactly once during
// startup.
func UseStorageDevice() {
	setRoot(storageDevice{Filesystem: osfs.New("")})
}

// -----------------------------------------------------------------------------
// Storage implementation

func (ms memoryStorage) RemoveAll(path string) error {
	return ms.Remove(path)
}

func (sd storageDevice) RemoveAll(path string) error {
	return os.RemoveAll(path)
}
