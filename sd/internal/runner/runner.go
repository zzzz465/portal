package runner

type Runner struct {
	datasource datasource
}

func NewRunner(datasource datasource) *Runner {
	runner := &Runner{datasource: datasource}

	return runner
}

func (r *Runner) Run() chan<- error {
	exit := make(chan<- error)

	go func() {
		exit <- r.run()
	}()

	return exit
}

func (r *Runner) run() error {
	return nil
}
