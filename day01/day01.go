package day01

import (
	"bufio"
	"strconv"
	"strings"
)

func getCalculationValue(line string) (int, error) {
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
