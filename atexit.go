// Package atexit provides a way to set a function that will be called at normal process termination (by ^C = SIGINT = "kill -2").
package atexit

import (
	"os"
	"os/signal"
)

var s chan os.Signal

// Use registers the atexit function.
// It can be called only once per program.
func Use(f func()) {
	if s != nil {
		panic("atexit can only be called a single time")
	}

	s = make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt)

	go func() {
		<-s
		f()
		os.Exit(0)
	}()
}
