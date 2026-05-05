//go:build windows

package main

import (
	"io"
	"os"

	"golang.org/x/sys/windows"
)

// flushStdin discards any pending bytes in the console input buffer for the
// given reader. If the reader is not an *os.File backed by a console, this is
// a no-op, so pipe-based stdin (used by tests) is unaffected.
//
// This prevents pre-restart, unsubmitted input from leaking into the next
// command session when -stdin is enabled. See issue #24.
func flushStdin(r io.Reader) {
	f, ok := r.(*os.File)
	if !ok {
		return
	}
	_ = windows.FlushConsoleInputBuffer(windows.Handle(f.Fd()))
}
