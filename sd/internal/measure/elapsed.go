package measure

import (
	"time"
)

type logFunc func(format string, args ...any)

func Elapsed(logFunc logFunc, format string) func() {
	now := time.Now()
	return func() {
		logFunc(format, time.Since(now))
	}
}
