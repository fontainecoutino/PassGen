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
	safe, strong, stringest, unsafe(not recommended). It can also turn a word or phrase into
	a password. Lastly, it also generates passwords based on specific user inputs for length
	and character types.`,
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.passgen.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

