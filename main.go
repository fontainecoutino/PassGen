/*
Copyright © 2022 Fontaine Coutino
*/
package main

import (
	"fmt"

	"github.com/fcoutino/passgen/models"
	"github.com/fcoutino/passgen/cmd"
)

// GENERATE PASSWORD AND LET USER PICK STRENGTH TYPE (UNSAFE/SAFE/STRONG/SRONGEST)
// GENERTAE PASSWORD BASED ON ATRTIBUTES (LENGTH/TYPE OF CHARS)
// GENERATE PASSWORD SUBSTITUTING CHARACTERS IN WORD TO ASCII VALS

func main() {
	fmt.Println(models.GenerateCharacters(10, ""))
	cmd.Execute()
}