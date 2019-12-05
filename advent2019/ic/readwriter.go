package ic

import (
	"bytes"
	"io"
	"os"
)

// The primary motivation for this is to be able to embed a
// bytes.Buffer with stdout; stdout and the buffer are both io.Writers
// and I want to make it explicit which is being used for what.
type readWriter struct {
	r io.Reader
	w io.Writer
}

func (rw *readWriter) Read(p []byte) (int, error) {
	return rw.r.Read(p)
}

func (rw *readWriter) Write(p []byte) (int, error) {
	return rw.w.Write(p)
}

func Console() io.ReadWriter {
	return &readWriter{
		r: os.Stdin,
		w: os.Stderr,
	}
}

// Canned provides a canned response to anything reading while still
// emitting to stdout.
func Canned(s string) io.ReadWriter {
	return &readWriter{
		r: bytes.NewBufferString(s),
		w: os.Stdout,
	}
}
