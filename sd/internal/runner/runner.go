package runner

import (
	"context"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"time"
)

type Runner struct {
	datasource datasource
	store      *store
	log        *zap.SugaredLogger
}

func NewRunner(datasource datasource, store *store, log *zap.SugaredLogger) (*Runner, error) {
	runner := &Runner{datasource: datasource, store: store, log: log}

	if log == nil {
		log, err := zap.NewDevelopment()
		if err != nil {
			return nil, err
		}

		runner.log = log.Sugar()
	}

	return runner, nil
}

func (r *Runner) Run(ctx context.Context, errChan chan<- error) {
	go func() {
		err := r.run(ctx)
		if err != nil {
			errChan <- err
		}
	}()
}

func (r *Runner) run(ctx context.Context) error {
	for {
		err := r.updateRecords(ctx)
		if err != nil {
			r.log.Errorf("failed update records. %v", err)
		}

		select {
		case <-ctx.Done():
			return nil

		default:
			break
		}

		wait := time.Duration(r.datasource.TTL()) * time.Second
		<-time.After(wait)
	}
}

func (r *Runner) updateRecords(ctx context.Context) error {
	// TODO: impl this
	return errors.New("method not implemented.")
}
