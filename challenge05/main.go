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

	var patrons []string
	err = json.NewDecoder(f).Decode(&patrons)

	return patrons, err
}

func ArrayFilter(arr []int, fn func(i int, el int) bool) []int {
	var n = 0
	for i, el := range arr {
		if fn(i, el) {
			arr[n] = el
			n++
		}
	}

	return arr[:n]
}

func ArrayIndexes(arr []string) []int {
	arrIndexes := make([]int, len(arr))

	for i := range arr {
		arrIndexes[i] = i
	}

	return arrIndexes
}

func main() {
	patrons, err := ReadJson("data")
	if err != nil {
		panic(err)
	}
	patronsIndexes := ArrayIndexes(patrons)

	for len(patronsIndexes) > 1 {
		isEven := len(patronsIndexes)%2 == 0

		patronsIndexes = ArrayFilter(patronsIndexes, func(i int, el int) bool {
			return i%2 == 0
		})

		if !isEven {
			patronsIndexes = patronsIndexes[1:]
		}
	}

	fmt.Printf("submit %s-%d\n", patrons[patronsIndexes[0]], patronsIndexes[0])
}
