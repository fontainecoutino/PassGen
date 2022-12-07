/*
Copyright © 2022 Fontain Coutino
*/
package cmd

import (
	"fmt"

	"github.com/fcoutino/passgen/models"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getCmd)

	// Local flags
	getCmd.Flags().StringP("mode", "m", "strong",
		"Type of password to generate (unsafe, safe, strong, strongest)")

	getCmd.Flags().IntP("length", "l", 0,
		"Password length")

	getCmd.Flags().StringP("type", "t", "",
		`Types of characters {l: lowercase, u: uppercase, n: number, s: symbols}
		ex. passgen get -t lns
		flag ignores any other character SO DON'T EVEN TRY`)
}

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get [-m mode] [-l length] [-t type]",
	Short: "Generate secure password",
	Long: `The 'get' command generates a password. You can choose the mode or specify
	the length or character type. Modes available are unsafe (-l 6 -t lun), safe 
	(-l 8 -t luns), strong (-l 12 -t luns), strongest (-l 20 -t luns). For
	specifications, length has to be grater than 0 and type works for lowercase,
	upercase, number, and special characters.`,
	Run: func(cmd *cobra.Command, args []string) {
		var passwordLenght int
		var characterType string

		// if mode specified
		mode, _ := cmd.Flags().GetString("mode")
		switch {
		case mode == "unsafe":
			passwordLenght = 6
			characterType = "lun"
		case mode == "safe":
			passwordLenght = 8
			characterType = "luns"
		case mode == "strong":
			passwordLenght = 12
			characterType = "luns"
		case mode == "strongest":
			passwordLenght = 20
			characterType = "luns"
		case mode == "":
			passwordLenght = 12
			characterType = "luns"
		default:
			fmt.Println("Mode not valid ... generated strong password anyways")
			passwordLenght = 12
			characterType = "luns"
		}

		// custom length and character type to override
		inputLenght, _ := cmd.Flags().GetInt("length")
		if inputLenght > 0 {
			passwordLenght = inputLenght
		}

		inputType, _ := cmd.Flags().GetString("type")
		if len(inputType) > 0 {
			characterType = inputType
		}

		// generate password
		password := models.GenerateCharacters(passwordLenght, characterType)
		CopyToClipboard(password)
	},
}
