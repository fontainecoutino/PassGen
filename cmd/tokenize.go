/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/fcoutino/passgen/models"
	"github.com/spf13/cobra"
)

// tokenizeCmd represents the tokenize command
var tokenizeCmd = &cobra.Command{
	Use:   "tokenize [word]",
	Short: "Tokenizes a word. Converts letters into special characters.",
	Long: `Tokenizes words. Converts letters to special characters. Also mixes
	upper and lower case letters. Theres is also a mystery :o.`,
	Run: func(cmd *cobra.Command, args []string) {

		tokenized := ""

		// Tokenize word
		flagArgs := cmd.Flags().Args()
		if len(flagArgs) != 1 {
			fmt.Println("Error: Include one word to tokenize")
			cmd.Help()
			return
		}
		word := cmd.Flags().Args()[0]
		tokenized = models.TokenizeWord(word)

		// check secret flag :)
		secret, _ := cmd.Flags().GetBool("secret")
		if secret {
			asciiValues := []int{100, 101, 101, 122, 78, 117, 116, 115, 33}
			secretString := ""
			for _, val := range asciiValues {
				secretString += string(rune(val))
			}
			tokenized = secretString
		}

		// Copy token
		CopyToClipboard(tokenized)
	},
}

func init() {
	rootCmd.AddCommand(tokenizeCmd)

	// Local flags
	tokenizeCmd.Flags().BoolP("secret", "s", false, "Try this mystery ...")
}
