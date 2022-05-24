package cmd

import (
    "context"
    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/service/route53"
    "github.com/cockroachdb/errors"
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
    "github.com/zzzz465/portal/sd/internal/datasource"
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

    serveCmd.Flags().String("address", ":4000", "configure host and port")
    if err := viper.BindPFlag("address", serveCmd.Flags().Lookup("address")); err != nil {
        log.Panic(err)
    }
    serveCmd.Flags().Bool("route53", false, "enable query records from aws route53.")
    if err := viper.BindPFlag("datasource.AWSRoute53.enabled", serveCmd.Flags().Lookup("route53")); err != nil {
        log.Panic(err)
    }
    serveCmd.Flags().Bool("static", false, "enable query records from static config.")
    if err := viper.BindPFlag("datasource.static.enabled", serveCmd.Flags().Lookup("static")); err != nil {
        log.Panic(err)
    }
    serveCmd.Flags().String("static-datasource-config", "", "configuration file for static datasource")
    if err := viper.BindPFlag("datasource.static.valueFile", serveCmd.Flags().Lookup("static-datasource-config")); err != nil {
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
    dsMap := make(map[string]datasource.Datasource)

    if viper.GetBool("datasource.AWSRoute53.enabled") {
        log.Debug("datasource route53 enabled.")
        r, ds, err := initAWSRoute53Runner(inMemoryStore)
        if err != nil {
            log.Panic(err)
        }

        runners = append(runners, *r)
        dsMap[ds.Identifier()] = ds
    }

    if viper.GetBool("datasource.static.enabled") {
        log.Debug("datasource static enabled.")
        r, ds, err := initStaticRunner(inMemoryStore)
        if err != nil {
            log.Panic(err)
        }

        runners = append(runners, *r)
        dsMap[ds.Identifier()] = ds
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

    addr := viper.GetString("address")
    log.Infof("starting server with addr %s", addr)

    server := web.NewHTTPServer(inMemoryStore, dsMap)
    serverError := server.Start(addr)

    if serverError != nil {
        log.Errorf("serverError: %+v", serverError)
    }

    wg.Wait()
}

func initAWSRoute53Runner(store store.Store) (*runner.Runner, datasource.Datasource, error) {
    awsConfig, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
    if err != nil {
        return nil, nil, errors.Wrap(err, "failed init aws config.")
    }

    client := route53.NewFromConfig(awsConfig)

    ds, err := awsroute53.NewDataSource(client, nil)
    if err != nil {
        return nil, nil, errors.Wrap(err, "failed creating aws route53 datasource.")
    }

    r, err := runner.NewRunner(ds, store, nil)
    if err != nil {
        return nil, nil, errors.Wrap(err, "failed creating route53 runner.")
    }

    return r, ds, nil
}

func initStaticRunner(store store.Store) (*runner.Runner, datasource.Datasource, error) {
    v := viper.New()
    cfgPath := viper.GetString("datasource.static.valueFile")
    v.SetConfigFile(cfgPath)

    err := v.ReadInConfig()
    if err != nil {
        return nil, nil, err
    }

    v.WatchConfig()

    ds := static.NewDataSource(v)

    r, err := runner.NewRunner(ds, store, nil)
    if err != nil {
        return nil, nil, errors.Wrap(err, "failed creating route53 runner.")
    }

    return r, ds, nil
}
