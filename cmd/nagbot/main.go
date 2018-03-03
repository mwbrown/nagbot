package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nagbot",
	Short: "Nagbot CLI application.",
}

/*
func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {

}
*/

func main() {

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
