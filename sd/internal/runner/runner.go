package runner

import (
    "context"
    "github.com/cockroachdb/errors"
    "github.com/hashicorp/go-multierror"
    "github.com/zzzz465/portal/sd/internal/measure"
    "github.com/zzzz465/portal/sd/internal/store"
    "github.com/zzzz465/portal/sd/internal/timer"
    "go.uber.org/zap"
    "time"
)

type Runner struct {
    datasource datasource
    store      store.Store
    log        *zap.SugaredLogger
    running    bool

    jobChan chan any

    debounceTime time.Duration
    debounce     timer.Debounce
}

func NewRunner(datasource datasource, store store.Store, log *zap.SugaredLogger) (*Runner, error) {
    runner := &Runner{datasource: datasource, store: store, log: log, jobChan: make(chan any, 1), debounceTime: time.Second}

    if log == nil {
        log, err := zap.NewDevelopment()
        if err != nil {
            return nil, err
        }

        runner.log = log.Sugar()
    }

    return runner, nil
}

func (r *Runner) Start(ctx context.Context) <-chan error {
    errChan := make(chan error)

    if r.running {
        errChan <- errors.New("runner is already running!")
    }

    var cancelFunc context.CancelFunc = nil

    TTL := r.datasource.TTL()
    if TTL > 0 {
        r.log.Infof("starting runner. interval: %s", r.datasource.TTL())
        var intervalCtx context.Context

        intervalCtx, cancelFunc = context.WithCancel(ctx)
        Interval(intervalCtx, r.datasource.TTL(), r.jobChan)
    } else {
        r.jobChan <- struct{}{}
    }

    if updatable, ok := r.datasource.(updatable); ok {
        updatable.OnDatasourceUpdated(r.onUpdatedCallback)
    }

    go func() {
        errChan <- r.startJobReceiver(ctx)

        if cancelFunc != nil {
            cancelFunc()
        }
    }()

    return errChan
}

func (r *Runner) startJobReceiver(ctx context.Context) error {
    for {
        <-r.jobChan

        r.debounce.Invoke(func() { r.executeTask(ctx) }, r.debounceTime)

        select {
        case <-ctx.Done():
            return nil
        default:
        }
    }
}

func (r *Runner) executeTask(ctx context.Context) {
    r.log.Infof("start updating records...")
    defer measure.Elapsed(r.log.Infof, "update took: %v")()

    err := r.updateRecords(ctx)
    if err != nil {
        r.log.Errorf("failed update records. err: %+v", err)
    }
}

func (r *Runner) updateRecords(ctx context.Context) error {
    var err error

    records, err := r.datasource.FetchRecords()
    if err != nil {
        return err
    }

    for _, record := range records {
        err2 := r.store.WriteRecord(record)
        if err2 != nil {
            err = multierror.Append(err, err2)
        }
    }

    return err
}

func (r *Runner) onUpdatedCallback() {
    r.jobChan <- struct{}{}
}
