package lw

import (
	"bytes"
	"testing"
)

func TestLimitedWriter(t *testing.T) {
	b := new(bytes.Buffer)
	l := LimitWriter(b, 10)
	n, err := l.Write([]byte{0, 1, 2, 3})
	if err != nil {
		t.Fatalf("error received - %s", err)
	} else if n != 4 {
		t.Fatalf("expected 4 bytes written, got %d", n)
	}
	n, err = l.Write([]byte{4, 5, 6})
	if err != nil {
		t.Fatalf("error received - %s", err)
	} else if n != 3 {
		t.Fatalf("expected 3 bytes written, got %d", n)
	}
	n, err = l.Write([]byte{7, 8, 9, 10, 11, 12})
	if _, ok := err.(*NoSpace); !ok {
		t.Fatalf("expected error of type NoSpace, got %s", err)
	} else if n != 3 {
		t.Fatalf("expected 3 bytes written, got %d", n)
	}

}
