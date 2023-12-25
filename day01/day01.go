package day01

import (
	"bufio"
	"strconv"
	"unicode"
)

func getCalculationValue(line string) (int, error) {
	firstAssigned := false
	var first string
	var last string
	for _, ch := range line {
		if !unicode.IsDigit(ch) {
			continue
		}

		if !firstAssigned {
			first = string(ch)
			firstAssigned = true
		}

		last = string(ch)
	}
	return strconv.Atoi(first + last)
}

func SolvePuzzle1(lines *bufio.Scanner) (string, error) {
	sum := 0

	for lines.Scan() {
		calculationValue, err := getCalculationValue(lines.Text())

		if err != nil {
			return "", err
		}

		sum += calculationValue
	}

	return strconv.Itoa(sum), nil
}

func SolvePuzzle2(lines *bufio.Scanner) (string, error) {
	return "", nil
}
