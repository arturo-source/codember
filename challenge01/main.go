package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Data to be present in each line:
usr: nombre de usuario
eme: email
psw: contraseña
age: edad
loc: ubicación
fll: número de seguidores
*/
var DATA_TO_BE_PRESENT = []string{"usr", "eme", "psw", "age", "loc", "fll"}

func ReadAllLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func LinesToData(lines []string) []string {
	var data []string
	var lineBuffer string
	for _, line := range lines {
		if line == "" {
			data = append(data, lineBuffer)
			lineBuffer = ""
			continue
		}
		lineBuffer += line + " "
	}
	if lineBuffer != "" {
		data = append(data, lineBuffer)
	}

	return data
}

func Contains(slice []string, element string) bool {
	for _, a := range slice {
		if a == element {
			return true
		}
	}
	return false
}

func DataIsValid(data string) bool {
	keys := []string{}
	keyValuePairs := strings.Split(data, " ")

	for _, pair := range keyValuePairs {
		keyValue := strings.Split(pair, ":")
		keys = append(keys, keyValue[0])
	}

	for _, key := range DATA_TO_BE_PRESENT {
		if !Contains(keys, key) {
			return false
		}
	}

	return true
}

func GetUsername(data string) string {
	keyValuePairs := strings.Split(data, " ")
	for _, pair := range keyValuePairs {
		keyValue := strings.Split(pair, ":")
		if keyValue[0] == "usr" {
			return keyValue[1]
		}
	}
	return ""
}

func main() {
	lines, err := ReadAllLines("data")
	if err != nil {
		panic(err)
	}
	data := LinesToData(lines)

	invalidUsers := 0
	lastValidUsername := ""
	for _, dataLine := range data {
		if DataIsValid(dataLine) {
			invalidUsers++
			lastValidUsername = GetUsername(dataLine)
		}
	}

	fmt.Print("submit ", invalidUsers, lastValidUsername)
}
