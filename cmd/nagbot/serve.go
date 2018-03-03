package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts a Nagbot server.",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("TODO: serve cmd")
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
