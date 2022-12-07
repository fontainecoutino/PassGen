/*
Copyright © 2022 Fontaine Coutino
*/
package cmd

import (
	"fmt"
	"golang.design/x/clipboard"
)

// Shared functions
func CopyToClipboard(password string) {
	err := clipboard.Init()
	if err != nil {
		fmt.Println("There was an error copying to the clipboard. Here it is printed.")
		fmt.Println(password)
		return
	}
	clipboard.Write(clipboard.FmtText, []byte(password))
	fmt.Printf("Copied password to clipboard!\n")
}