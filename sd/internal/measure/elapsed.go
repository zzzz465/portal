package measure

import (
	"go.uber.org/zap"
	"time"
)

func Elapsed(log *zap.SugaredLogger, format string) func() {
	now := time.Now()
	return func() {
		log.Debugf(format, time.Since(now))
	}
}
