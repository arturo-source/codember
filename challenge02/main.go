package main

import (
	"fmt"
	"os"
	"strconv"
)

func ReadFile(path string) (string, error) {
	dat, err := os.ReadFile(path)
	return string(dat), err
}

func IsValidChar(char rune) bool {
	return char >= 'a' && char <= 'z'
}

func main() {
	content, err := ReadFile("data")
	if err != nil {
		panic(err)
	}

	var asciiCode string
	var message string
	for _, char := range content {
		if char == ' ' {
			asciiInt, _ := strconv.Atoi(asciiCode)
			asciiCode = ""

			message += string(rune(asciiInt))
			message += " "

			continue
		}

		asciiCode += string(char)
		asciiInt, _ := strconv.Atoi(asciiCode)
		if IsValidChar(rune(asciiInt)) {
			message += string(rune(asciiInt))
			asciiCode = ""
		}
	}

	fmt.Println("submit", message)
}
