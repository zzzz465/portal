package cmd

import (
    "context"
    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/service/route53"
    "github.com/cockroachdb/errors"
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
    "github.com/zzzz465/portal/sd/internal/datasource/awsroute53"
    "github.com/zzzz465/portal/sd/internal/datasource/static"
    "github.com/zzzz465/portal/sd/internal/runner"
    "github.com/zzzz465/portal/sd/internal/store"
    "github.com/zzzz465/portal/sd/internal/web"
    "sync"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
    Use: "serve",
    Run: runServe,
}

func init() {
    rootCmd.AddCommand(serveCmd)

    // Here you will define your flags and configuration settings.
    serveCmd.Flags().Bool("route53", false, "enable query records from aws route53.")
    if err := viper.BindPFlag("datasource.AWSRoute53.enabled", serveCmd.Flags().Lookup("route53")); err != nil {
        log.Panic(err)
    }
    serveCmd.Flags().Bool("static", false, "enable query records from static config.")
    if err := viper.BindPFlag("datasource.static.enabled", serveCmd.Flags().Lookup("static")); err != nil {
        log.Panic(err)
    }
}

func runServe(cmd *cobra.Command, args []string) {
    // TODO: replace in-memory store to any store that given from argument.
    // TODO: support selecting (multiple) data source. maybe use config file?

    inMemoryStore := store.NewInMemoryStore()
    ctx, cancel := context.WithCancel(context.TODO())
    defer cancel()

    runners := make([]runner.Runner, 0)

    if viper.GetBool("datasource.AWSRoute53.enabled") {
        log.Debug("datasource route53 enabled.")
        r, err := initAWSRoute53Runner(inMemoryStore)
        if err != nil {
            log.Panic(err)
        }

        runners = append(runners, *r)
    }

    if viper.GetBool("datasource.static.enabled") {
        log.Debug("datasource static enabled.")
        r, err := initStaticRunner(inMemoryStore)
        if err != nil {
            log.Panic(err)
        }

        runners = append(runners, *r)
    }

    if len(runners) == 0 {
        log.Panic("no runners are enabled. at least 1 runner required.")
    }

    wg := sync.WaitGroup{}
    for _, r := range runners {
        errChan := r.Start(ctx)
        wg.Add(1)
        go func() {
            err := <-errChan
            if err != nil {
                log.Errorf("runner ended with error: %+v", err)
            }

            wg.Done()
        }()
    }

    server := web.NewHTTPServer(inMemoryStore)
    serverError := server.Start()

    if serverError != nil {
        log.Errorf("serverError: %+v", serverError)
    }

    wg.Wait()
}

func initAWSRoute53Runner(store store.Store) (*runner.Runner, error) {
    awsConfig, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
    if err != nil {
        return nil, errors.Wrap(err, "failed init aws config.")
    }

    client := route53.NewFromConfig(awsConfig)

    ds, err := awsroute53.NewDataSource(client, nil)
    if err != nil {
        return nil, errors.Wrap(err, "failed creating aws route53 datasource.")
    }

    r, err := runner.NewRunner(ds, store, nil)
    if err != nil {
        return nil, errors.Wrap(err, "failed creating route53 runner.")
    }

    return r, nil
}

func initStaticRunner(store store.Store) (*runner.Runner, error) {
    ds := static.NewDataSource(viper.GetViper())

    r, err := runner.NewRunner(ds, store, nil)
    if err != nil {
        return nil, errors.Wrap(err, "failed creating route53 runner.")
    }

    return r, nil
}
