package cmd

import (
	"github.com/spf13/cobra"
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

	s := store.NewInMemoryStore()
	server := web.NewHTTPServer(s)
	if err := server.Start(); err != nil {
		log.Error(err)
	}
}
