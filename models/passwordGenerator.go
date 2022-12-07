package models

import (
	"math/rand"
	"strings"
	"time"
)

// Character constants
var Lowercase = "abcdefghijklmnopqrstuvwxyz"
var Uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
var Numbers = "0123456789"
var Symbols = "!@#$%^&*()-_=+[];:.,?"

func GenerateCharacters(length int, types string) string {
	// to set random seeding
	source := rand.New(rand.NewSource(time.Now().UnixNano()))

	// get character array from given type list
	characters := ""
	if strings.Contains(types, "l") {
		characters += Lowercase
	}
	if strings.Contains(types, "u") {
		characters += Uppercase
	}
	if strings.Contains(types, "n") {
		characters += Numbers
	}
	if strings.Contains(types, "s") {
		characters += Symbols
	}
	if len(characters) <= 0 {
		characters += Lowercase + Uppercase + Numbers + Symbols
	}

	// shuffle characters to be used
	inRune := []rune(characters)
	source.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})
	characters = string(inRune)

	// generate n-sequence of characters from chaacters array
	var sequence = ""
	l := length
	for l > 0 {
		sequence += selectCharFromArray(characters, source)
		l -= 1
	}

	// validate
	if generatedSequenceIsValid(sequence) {
		return sequence
	} else {
		return GenerateCharacters(length, types)
	}
}

// ADD RULES BELOW AS IF STATEMENTS
func generatedSequenceIsValid(sequence string) bool {
	if strings.Contains(sequence, ".") ||
		strings.HasPrefix(sequence, "@") ||
		strings.HasPrefix(sequence, "-") {
		return false
	}
	return true
}

func selectCharFromArray(characters string, source *rand.Rand) string {
	return string(characters[source.Intn(len(characters))])
}
