package box

import (
	"errors"
	"os"
	"strings"
)

var Result string

func ChangeToAsciiArt(text, font string) (string, error) {
	Result = ""
	FontPath := CheckArgs(font)

	// read file and check error
	contentByte, err := os.ReadFile(FontPath)
	if err != nil {
		return "", errors.New("error reading file: " + err.Error())
	}

	content := string(contentByte[1:])
	if FontPath == "fs/thinkertoy.txt" {
		content = strings.ReplaceAll(content, "\r", "")
	}

	// check character is printable
	var hundleText string
	for _, char := range text {
		if char <= 126 && char >= 32 || char == 10 {
			hundleText += string(char)
		}
	}

	aryName := strings.Split(hundleText, "\n")
	contentAry := strings.Split(content, "\n\n")

	// Print out the output
	for i := range aryName {
		if aryName[i] == "" {
			Result += "\n"
		}
		Result += PrintWord(aryName[i], contentAry)
	}
	
	return Result, nil
}
