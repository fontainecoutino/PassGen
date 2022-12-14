/*
Copyright © 2022 Fontaine Coutino
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "passgen",
	Short: "Password Generator CLI is an application to generate secure passwords.",
	Long: `Password Generator CLI is an application written in Go. Built to generate
	secure passwords based on users desires. It can generate multiple types of passwords;
	safe, strong, strongest, unsafe. Custom length and character type can be specified
	as well. PassGen can also turn a word or phrase into a password.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Global flags
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.passgen.yaml)")
}