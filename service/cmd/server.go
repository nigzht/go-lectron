package main

import (
	"fmt"
	"log"
	"service/cmd/server"

	"github.com/spf13/cobra"
)

var (
	port int

	serverCmd = &cobra.Command{
		Use:     "server",
		Short:   "boot up the rpc server.",
		Run:     serverFunc,
		Example: fmt.Sprintf("%s server --port=6969", appName),
		ValidArgs: []string{
			"port",
		},
	}

	serverFunc = func(cmd *cobra.Command, args []string) {
		if err := server.Start(port); err != nil {
			log.Fatal(err)
		}
	}
)

func init() {
	rootCmd.AddCommand(serverCmd)
}
