package lw

import (
	"errors"
	"io"
)

type NoSpace struct{}

func (n *NoSpace) Error() string {
	return "no space left"
}

var noSpace error = &NoSpace{}

type LimitedWriter struct {
	W io.Writer
	N int64
}

func LimitWriter(w io.Writer, n int64) io.Writer {
	return &LimitedWriter{w, n}
}

func (l *LimitedWriter) Write(p []byte) (n int, err error) {
	if l.N <= 0 {
		err = noSpace
	} else if lp := int64(len(p)); lp <= l.N {
		n, err = l.W.Write(p)
	} else if n, err = l.W.Write(p[:l.N]); err == nil {
		err = noSpace
	}
	l.N -= int64(n)
	return
}
