package main

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   appName,
	Short: appDescription,
	Run:   serverFunc,
}

func init() {
	rootCmd.PersistentFlags().IntVarP(&port, "port", "p", 6969, "port to run the rpc server on")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
