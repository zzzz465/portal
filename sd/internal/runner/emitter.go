package runner

import (
	"context"
	"time"
)

func Interval(ctx context.Context, interval time.Duration, target chan<- any) {
	go func() {
		for {
			target <- struct {}{}
			select {
			case <-ctx.Done():
				return

			default:
			}

			time.Sleep(interval)
		}
	}()
}
