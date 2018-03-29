package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/vito/go-interact/interact"

	// Importing config at the root also initializes Viper defaults.
	"github.com/mwbrown/nagbot/config"
)

var rootCmd = &cobra.Command{
	Use:   "nagbot",
	Short: "Nagbot CLI application.",
}

func init() {

	// Load JSON files only for now.
	viper.SupportedExts = []string{"json"}

	viper.SetConfigName("nagbot")
	viper.AddConfigPath("/etc/nagbot")
	viper.AddConfigPath("$HOME/.config/nagbot")

	// Bind environment variables.
	viper.SetEnvPrefix("nagbot")
	viper.AutomaticEnv()

	// Ignore any errors if a file cannot be found.
	viper.ReadInConfig()

	// Attempt to read in the auth file if it exists.
	nbconfig.LoadAuthFile()
}

func main() {

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// readUserPass returns both a username and string from a CLI interaction,
// with optional password confirmation (when creating a new user, for example).
func readUserPass(confirmPass bool) (user string, pass string, err error) {
	var pw_interact interact.Password
	var pw_confirm interact.Password

	err = interact.NewInteraction("Username").Resolve(interact.Required(&user))
	if err != nil {
		return
	}

	for {
		if err = interact.NewInteraction("Password").Resolve(interact.Required(&pw_interact)); err != nil {
			return
		}

		if !confirmPass {
			break
		}

		if err = interact.NewInteraction("Password (confirm)").Resolve(interact.Required(&pw_confirm)); err != nil {
			return
		}

		if string(pw_interact) == string(pw_confirm) {
			break
		}

		fmt.Println("Passwords do not match.")
	}

	pass = string(pw_interact)
	return
}
