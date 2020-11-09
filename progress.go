package goprogress

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"
)

type (
	// Progressbar .
	Bar struct {
		dst     io.Writer
		total   int64
		current int64
		target  string
		err     error
		data    chan int64
	}
)

// NewProgressbar .
func NewBar(w io.Writer) *Bar {
	if w == nil {
		w = os.Stdout
	}
	return &Bar{
		dst:  w,
		data: make(chan int64, 2),
	}
}

// SetTargetName .
func (b *Bar) SetTargetName(name string) {
	b.target = name
}

// SetTotal .
func (b *Bar) SetTotal(t int64) {
	b.total = t
}

// AdvanceProgress .
func (b *Bar) AdvanceProgress(i int64) bool {
	if b.err != nil {
		return false
	}
	b.data <- i
	return true
}

func (b *Bar) Start(ctx context.Context) {
	go func() {
	WC:
		for {
			select {
			case <-ctx.Done():
				b.err = ctx.Err()
				break WC
			case v, ok := <-b.data:
				if ok {
					b.current = v
					_, err := b.dst.Write([]byte(fmt.Sprintf("\rProgress: [%-50s]%3d%% %8d/%d", strings.Repeat("▮", int(v)/2), b.current, b.current, b.total)))
					if err != nil {
						b.err = err
						break WC
					}
				} else {
					break WC
				}
			}
		}
	}()
}

// Err .
func (b *Bar) Err() error {
	return b.err
}

// Finish .
func (b *Bar) Finish() {
	close(b.data)
}
