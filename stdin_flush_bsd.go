//go:build darwin || freebsd || netbsd || openbsd || dragonfly

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
	// TIOCFLUSH expects a pointer to an int bitmask: FREAD=1, FWRITE=2.
	// FREAD is not exported by x/sys/unix on all BSD targets, so use the
	// literal value to flush the input queue only.
	const fread = 1
	_ = unix.IoctlSetPointerInt(int(f.Fd()), unix.TIOCFLUSH, fread)
}
