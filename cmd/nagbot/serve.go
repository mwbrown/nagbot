package main

import (
	"github.com/mwbrown/nagbot/server"
	"github.com/spf13/cobra"
	"log"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts a Nagbot server.",
	Run:   serveCmdHandler,
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func serveCmdHandler(cmd *cobra.Command, args []string) {
	nb, err := server.NewServer()

	if err != nil {
		log.Fatalf("Could not create server: %v\n", err)
	}

	if err := nb.ServeGrpc(); err != nil {
		log.Fatalf("Nagbot server terminated: %v\n", err)
	}
}
