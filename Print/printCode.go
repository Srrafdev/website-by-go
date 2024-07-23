package box

import (
	"strings"
)

func PrintWord(text string, contentAry []string) string {
	var printChar string

	if text == "" {
		return ""
	}
	// Create a slice of slices to hold the characters of the word
	k := make([][]string, len(text))
	for i, val := range text {
		if val-32 < 0{
			return "error"
		}
		k[i] = strings.Split(contentAry[val-32], "\n")
	}

	for q := 0; q < 8; q++ {
		for l := 0; l < len(text); l++ {
			if l >= len(k) || q >= len(k[l]) {
				return "try Again"
			}
			printChar += (k[l][q])
		}
		printChar += "\r\n"
	}
	return printChar
}
//4565465 45645647564564
