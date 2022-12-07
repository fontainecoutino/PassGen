/*
Copyright © 2022 Fontaine Coutino
*/
package cmd

import (
	"fmt"
	"golang.design/x/clipboard"
)

// SHARED FUNCTIONS

// Copy a string to the clipboard
func CopyToClipboard(seq string) {
	err := clipboard.Init()
	if err != nil {
		fmt.Println("There was an error copying to the clipboard. Here it is printed.")
		fmt.Println(seq)
		return
	}
	clipboard.Write(clipboard.FmtText, []byte(seq))
	fmt.Printf("Copied to clipboard!\n")
}