package main

import "fmt"

func SplitNumberByDigit(number int) []int {
	var digits []int
	for number > 0 {
		digits = append(digits, number%10)
		number /= 10
	}
	return digits
}

func IsValidNumber(number int) bool {
	digits := SplitNumberByDigit(number)
	for i := 0; i < len(digits)-1; i++ {
		if digits[i] < digits[i+1] {
			return false
		}
	}

	numOfFives := 0
	for _, digit := range digits {
		if digit == 5 {
			numOfFives++
		}
	}

	return numOfFives >= 2
}

func main() {
	var validNumber []int
	const start = 11098
	const end = 98123

	for i := start; i <= end; i++ {
		if IsValidNumber(i) {
			validNumber = append(validNumber, i)
		}
	}

	fmt.Printf("submit %d-%d\n", len(validNumber), validNumber[55])
}
