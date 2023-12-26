package day01

import (
	"bufio"
	"strconv"
	"strings"
)

var digitWordMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func getCalculationValue(line string, includeDigitWords bool) (int, error) {
	var first, last string
	indexOfFirst, indexOfLast := len(line), -1

	for i := 0; i < 10; i++ {
		iAsString := strconv.Itoa(i)

		if index := strings.Index(line, iAsString); index >= 0 && index < indexOfFirst {
			indexOfFirst = index
			first = iAsString
		}

		if index := strings.LastIndex(line, iAsString); index >= 0 && index > indexOfLast {
			indexOfLast = index
			last = iAsString
		}
	}

	if includeDigitWords {
		for digitWord, digit := range digitWordMap {
			if index := strings.Index(line, digitWord); index >= 0 && index < indexOfFirst {
				indexOfFirst = index
				first = digit
			}

			if index := strings.LastIndex(line, digitWord); index >= 0 && index > indexOfLast {
				indexOfLast = index
				last = digit
			}
		}
	}

	return strconv.Atoi(first + last)
}

func solvePuzzle(lines *bufio.Scanner, includeDigitWords bool) (string, error) {
	sum := 0

	for lines.Scan() {
		calculationValue, err := getCalculationValue(lines.Text(), includeDigitWords)

		if err != nil {
			return "", err
		}

		sum += calculationValue
	}

	return strconv.Itoa(sum), nil
}

func SolvePuzzle1(lines *bufio.Scanner) (string, error) {
	return solvePuzzle(lines, false)
}

func SolvePuzzle2(lines *bufio.Scanner) (string, error) {
	return solvePuzzle(lines, true)
}
