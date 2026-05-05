//go:build linux

package main

import (
	"io"
	"os"

	"golang.org/x/sys/unix"
)

// flushStdin discards any pending bytes in the kernel TTY input queue for the
// given reader. If the reader is not an *os.File backed by a TTY, this is a
// no-op (ENOTTY is ignored), so pipe-based stdin (used by tests) is unaffected.
//
// This prevents pre-restart, unsubmitted input from leaking into the next
// command session when -stdin is enabled. See issue #24.
func flushStdin(r io.Reader) {
	f, ok := r.(*os.File)
	if !ok {
		return
	}
	_ = unix.IoctlSetInt(int(f.Fd()), unix.TCFLSH, unix.TCIFLUSH)
}
