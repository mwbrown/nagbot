package main

import (
	"fmt"
	"os"

	"github.com/mwbrown/nagbot/client"
	"github.com/mwbrown/nagbot/config"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Logs into a Nagbot server.",
	Run:   loginHandler,
}

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logs out of a Nagbot server.",
	Run:   logoutHandler,
}

var checkLoginCmd = &cobra.Command{
	Use:   "checklogin",
	Short: "Checks for a valid Nagbot session.",
	Run:   checkLoginHandler,
}

// Returns a client object. If the object cannot be created for whatever reason,
// the program exits.
func getClient() *client.Client {
	c, err := client.NewClient()
	if err != nil {
		fmt.Println("Error creating client:", err)
		os.Exit(1)
	}

	if err := c.Open(); err != nil {
		fmt.Println("Could not connect client:", err)
		os.Exit(1)
	}

	return c
}

func loginHandler(cmd *cobra.Command, args []string) {
	c := getClient()
	defer c.Close()

	user, pass, err := readUserPass(false)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	token, err := c.Login(user, pass)
	if err != nil {
		fmt.Println("Error logging in:", err)
		os.Exit(1)
	}

	// Sanity check the result.
	if len(token) == 0 {
		fmt.Println("Invalid token was received from the server.")
		os.Exit(1)
	}

	fmt.Println("Login token received successfully.")

	viper.Set(nbconfig.CFG_KEY_RPC_TOKEN, token)
	err = nbconfig.SaveAuthFile()
	if err != nil {
		fmt.Println("Could not save login token:", err)
		os.Exit(1)
	}

	fmt.Println("Successfully logged in as", user)
}

func logoutHandler(cmd *cobra.Command, args []string) {
	c := getClient()
	defer c.Close()

	err := c.Logout()
	if err != nil {
		fmt.Println("Error logging out:", err)
		os.Exit(1)
	}

	fmt.Println("Logged out successfully.")
}

func checkLoginHandler(cmd *cobra.Command, args []string) {
	c := getClient()
	defer c.Close()

	err := c.CheckLogin()
	if err != nil {
		fmt.Println("Check login error:", err)
		os.Exit(1)
	}

	fmt.Println("Current login token is valid.")
}

func init() {
	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(logoutCmd)
	rootCmd.AddCommand(checkLoginCmd)
}
