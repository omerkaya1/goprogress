package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/omerkaya1/goprogress"
)

func main() {
	b := goprogress.NewBar(os.Stdout)

	b.SetTotal(100)

	ctx, _ := context.WithTimeout(context.Background(), time.Second*15)
	b.Start(ctx)

	for i := 1; i <= 100; i++ {
		if !b.AdvanceProgress(int64(i)) {
			log.Println(b.Err())
			break
		}
		time.Sleep(time.Millisecond * 100)
	}
	b.Finish()
	os.Exit(0)
}
