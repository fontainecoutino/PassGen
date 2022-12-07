package models

import (
	"math/rand"
	"strings"
	"time"
)

func TokenizeWord(word string) string {
	tokenized := ""
	for _, ch := range word {
		tokenized += pickRandomTokenFromCharacter(string(ch))
	}
	if word == tokenized {
		return TokenizeWord(word)
	}
	return tokenized
}

func pickRandomTokenFromCharacter(char string) string {
	source := rand.New(rand.NewSource(time.Now().UnixNano())) // to set random seeding
	letters := getLettersMap()
	lowerChar := strings.ToLower(char)
	if tokens, ok := letters[lowerChar]; ok {
		return tokens[source.Intn(len(tokens))]
	} else {
		return char
	}
}

// Hardcoded letter map
func getLettersMap() map[string][]string {
	var letters = make(map[string][]string)

	letters["a"] = []string{"a", "A", "@", "4"}
	letters["b"] = []string{"b", "B", "6", "8"}
	letters["c"] = []string{"c", "C", "("}
	letters["d"] = []string{"d", "D"}
	letters["e"] = []string{"e", "E", "3"}
	letters["f"] = []string{"f", "F"}
	letters["g"] = []string{"g", "G", "9"}
	letters["h"] = []string{"h", "H", "#"}
	letters["i"] = []string{"i", "I", "1", "!"}
	letters["j"] = []string{"j", "J"}
	letters["k"] = []string{"k", "K"}
	letters["l"] = []string{"l", "L"}
	letters["m"] = []string{"m", "M"}
	letters["n"] = []string{"n", "N", "^"}
	letters["o"] = []string{"o", "O", "0"}
	letters["p"] = []string{"p", "P"}
	letters["q"] = []string{"q", "Q"}
	letters["r"] = []string{"r", "R"}
	letters["s"] = []string{"s", "S", "5", "$"}
	letters["t"] = []string{"t", "T", "+"}
	letters["u"] = []string{"u", "U"}
	letters["v"] = []string{"v", "V"}
	letters["w"] = []string{"w", "W"}
	letters["x"] = []string{"x", "X"}
	letters["y"] = []string{"y", "Y"}
	letters["z"] = []string{"z", "Z"}

	return letters
}
