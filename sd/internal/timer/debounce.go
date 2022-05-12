package timer

import (
    "context"
    "time"
)

type Debounce struct {
    cancel context.CancelFunc
}

func (d *Debounce) Invoke(cb func(), delay time.Duration) {
    ctx, cancel := context.WithCancel(context.Background())

    if d.cancel != nil {
        d.cancel()
    }

    d.cancel = cancel
    d.invoke(cb, delay, ctx)
}

func (d *Debounce) invoke(cb func(), delay time.Duration, ctx context.Context) {
    time.Sleep(delay)

    select {
    case <-ctx.Done():
        return
    default:
        cb()
        d.cancel = nil
    }
}
