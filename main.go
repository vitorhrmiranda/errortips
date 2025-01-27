package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"
)

func main() {
	_, err := os.ReadFile("nonexistent.txt")
	if err == nil {
		return // no error
	}

	// #1 basic error check using string or substring
	if strings.Contains(err.Error(), "no such file or directory") {
		log.Printf("#1 - check error string: (%T) %v\n", err, err)
	}

	// #2 -> https://go.dev/doc/effective_go#errors
	/*
		type PathError struct {
			Op   string
			Path string
			Err  error
		}

		func (e *PathError) Error() string { return e.Op + " " + e.Path + ": " + e.Err.Error() }
	*/
	if e, ok := err.(*os.PathError); ok && e.Err == syscall.ENOENT {
		log.Printf("#2 - Check error type assertion: (%T) %v\n", err, err)
	}

	// #3 -> https://go.dev/wiki/ErrorValueFAQ#how-should-i-change-my-error-handling-code-to-work-with-the-new-features
	// https://go.googlesource.com/proposal/+/master/design/29934-error-values.md
	if errors.Is(err, os.ErrNotExist) {
		log.Printf("#3 - Check error value: (%T) %v\n", err, err)
	}

	// #4 -> encapsulate error
	// https://go.dev/blog/go1.13-errors
	wrappedErr := fmt.Errorf("read file in main: %w", err)
	if errors.Is(wrappedErr, os.ErrNotExist) {
		log.Printf("#4 - Check wrapped err value: (%T) %v\n", err, err)
	}
}
