# AtExit

AtExit registers a function to be called at normal process termination (by `^C` = `SIGINT` = `kill -2`).

## Installation

```Shell
$ go get github.com/whitedevops/atexit
```

## Usage [![GoDoc](https://godoc.org/github.com/whitedevops/atexit?status.svg)](https://godoc.org/github.com/whitedevops/atexit)

```Go
package main

import (
	"github.com/whitedevops/atexit"
)

func main() {
	atexit.Use(func() {
		// Do something before exit…
	})
}
```
