package goprogress

import (
	"io"
)

type (
	// Progressbar .
	Bar struct {
		dst    io.WriteSeeker
		target string
		err    error
	}
)

// NewProgressbar .
func NewBar(w io.WriteSeeker, targetName string) *Bar {
	return &Bar{
		dst:    w,
		target: targetName,
	}
}

// ReportProgress .
func (b *Bar) ReportProgress(i int) bool {
	return true
}

// Err .
func (b *Bar) Err() error {
	return b.err
}
