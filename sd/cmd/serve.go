package cmd

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/spf13/cobra"
	"github.com/zzzz465/portal/sd/internal/datasource/awsroute53"
	"github.com/zzzz465/portal/sd/internal/runner"
	"github.com/zzzz465/portal/sd/internal/store"
	"github.com/zzzz465/portal/sd/internal/web"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use: "serve",
	Run: runServe,
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runServe(cmd *cobra.Command, args []string) {
	// TODO: replace in-memory store to any store that given from argument.
	// TODO: support selecting (multiple) data source. maybe use config file?

	inMemoryStore := store.NewInMemoryStore()

	awsConfig, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		errExit(1, "failed init aws config. %+v", err)
	}

	client := route53.NewFromConfig(awsConfig)

	ds, err := awsroute53.NewDataSource(client, nil)
	if err != nil {
		errExit(1, "failed creating aws route53 datasource. %+v", err)
	}

	route53Runner, err := runner.NewRunner(ds, inMemoryStore, nil)
	if err != nil {
		errExit(1, "failed creating route53 runner. %+v", err)
	}

	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	var runnerError error
	// TODO: what's the better way to get error that raised in runner?
	route53Runner.Start(ctx, &runnerError)

	server := web.NewHTTPServer(inMemoryStore)
	serverError := server.Start()

	if serverError != nil {
		log.Errorf("serverError: %+v", serverError)
	}

	if runnerError != nil {
		log.Errorf("runnerError: %+v", runnerError)
	}
}
