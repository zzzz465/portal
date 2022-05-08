package runner

import (
	"context"
	"github.com/hashicorp/go-multierror"
	"go.uber.org/zap"
	"time"
)

type Runner struct {
	datasource datasource
	store      store
	log        *zap.SugaredLogger
}

func NewRunner(datasource datasource, store store, log *zap.SugaredLogger) (*Runner, error) {
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
		wait := time.Duration(r.datasource.TTL()) * time.Second
		nextTime := time.Now().Add(wait)

		r.log.Infof("start updating records...")
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

		r.log.Debugf("wait until %v...", time.Until(nextTime))
		<-time.After(time.Until(nextTime))
	}
}

func (r *Runner) updateRecords(ctx context.Context) error {
	var err error

	records, err := r.datasource.FetchRecords()
	if err != nil {
		return err
	}

	for _, record := range records {
		err2 := r.store.Write(record.Key, record)
		if err2 != nil {
			err = multierror.Append(err, err2)
		}
	}

	return err
}
