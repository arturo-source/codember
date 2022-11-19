package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func ReadJson(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	var colors []string
	err = json.NewDecoder(f).Decode(&colors)

	return colors, err
}

func main() {
	colors, err := ReadJson("data")
	if err != nil {
		panic(err)
	}

	var maxZebraPoints = 0
	var maxZebraLastColor = ""

	var lastColor = ""
	var nextColor = colors[0]
	var currMaxPoints = 1

	for _, currColor := range colors {
		if currColor != nextColor || lastColor == currColor {
			currMaxPoints = 1
		}

		currMaxPoints++
		lastColor, nextColor = currColor, lastColor

		if currMaxPoints > maxZebraPoints {
			maxZebraPoints = currMaxPoints
			maxZebraLastColor = lastColor
		}
	}

	fmt.Printf("submit %d@%s\n", maxZebraPoints, maxZebraLastColor)
}
