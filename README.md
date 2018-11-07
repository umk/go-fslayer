## Go file system abstraction

[![GoDoc](https://godoc.org/github.com/umk/go-fslayer?status.svg)](https://godoc.org/github.com/umk/go-fslayer)
[![Go Report Card](https://goreportcard.com/badge/github.com/umk/go-fslayer)](https://goreportcard.com/report/github.com/umk/go-fslayer)

The package provides an abstraction over the file system using [go-billy](https://github.com/src-d/go-billy) package. It exposes a single function `Fs()`, which provides the I/O methods for manipulating the files either on a persistent storage or in a memory, depending on how the package has been initialized.

The initialization should be performed during the application startup or the test initialization the following way: 

```go
import "github.com/umk/go-fslayer"

func main() {
	fslayer.UseStorageDevice(); // for persistent storage
	fslayer.UseMemoryStorage(); // to save files in memory
}
```
Either of these two functions may be called in a single process, and may be called just once. When called twice, the application will panic.

When trying to access the file system, use the following:

```go
import . "github.com/umk/go-fslayer"

func DoSmth() {
	f, err := Fs().Open("/foo/bar.txt")
	// ...
}
```
If `Fs()` is used without the storage type initialized, like shown before, the application will panic.

See this package's GoDoc reference and documentation for go-billy's [osfs](https://godoc.org/gopkg.in/src-d/go-billy.v4/osfs) and [memfs](https://godoc.org/gopkg.in/src-d/go-billy.v4/memfs) for the list of functions available.
