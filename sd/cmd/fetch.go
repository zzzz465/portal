/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/spf13/cobra"
	"github.com/zzzz465/portal/sd/internal/datasource/awsroute53"
)

// fetchCmd represents the fetch command
var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "fetch current records from datasource.",
}

var fetchRoute53RecordsCmd = &cobra.Command{
	Use:   "aws-route-53",
	Short: "use aws Route53 as data source.",
	Run:   fetchRoute53Records,
}

func init() {
	fetchCmd.AddCommand(fetchRoute53RecordsCmd)
	rootCmd.AddCommand(fetchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fetchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fetchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func fetchRoute53Records(cmd *cobra.Command, args []string) {
	awsConfig, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		errExit(1, "failed init aws config. %+v", err)
	}

	client := route53.NewFromConfig(awsConfig)

	ds, err := awsroute53.NewDataSource(client, nil)
	if err != nil {
		errExit(1, "failed creating aws route53 datasource. %+v", err)
	}

	records, err := ds.FetchRecords()
	if err != nil {
		errExit(1, "failed fetching records. %+v", err)
	}

	log.Infof("found %d records.", len(records))

	for _, record := range records {
		log.Info(record.Host)
	}
}
