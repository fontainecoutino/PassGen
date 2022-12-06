/*
Copyright © 2022 Fontain Coutino
*/
package cmd

import (
	"fmt"

	"github.com/fcoutino/passgen/models"
	"github.com/spf13/cobra"
	"golang.design/x/clipboard"
)

// Flags
var mode string

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Generate secure password",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// mode
		if mode != "" {
			getModePassword()
		}

		// custom length and character type

	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Cobra supports local flags which will only run when this command
	getCmd.Flags().StringVarP(&mode, "mode", "m", "strong",
		"Type of password to generate. (unsafe, safe, strong, strongest)")
}

func getModePassword() {
	var password string
	switch {
	case mode == "unsafe":
		password = models.GenerateCharacters(6, "lun")
	case mode == "safe":
		password = models.GenerateCharacters(8, "luns")
	case mode == "strong":
		password = models.GenerateCharacters(12, "luns")
	case mode == "strongest":
		password = models.GenerateCharacters(20, "luns")
	default:
		fmt.Println("Mode not valid ... generated strong password anyways.")
		mode = "strong"
		password = models.GenerateCharacters(12, "luns")
	}

	copyToClipboard(password)
}

func copyToClipboard(password string) {
	err := clipboard.Init()
	if err != nil {
		fmt.Println("There was an error copying to the clipboard. Here it is printed.")
		fmt.Println(password)
		return
	}
	clipboard.Write(clipboard.FmtText, []byte(password))
	fmt.Printf("Copied %s password to clipboard!\n", mode)
}
