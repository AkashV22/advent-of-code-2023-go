package day01

import (
	"bufio"
	"strconv"
	"strings"
)

type calculationValueBuilder struct {
	line         string
	first        string
	last         string
	indexOfFirst int
	indexOfLast  int
}

func (b calculationValueBuilder) init(line string) calculationValueBuilder {
	b.line = line
	b.indexOfFirst = len(line)
	b.indexOfLast = 0
	return b
}

func (b *calculationValueBuilder) findAndUpdate(valueToFind, valueToUpdateWith string) {
	if index := strings.Index(b.line, valueToFind); index >= 0 && index < b.indexOfFirst {
		b.indexOfFirst = index
		b.first = valueToUpdateWith
	}

	if index := strings.LastIndex(b.line, valueToFind); index >= 0 && index > b.indexOfLast {
		b.indexOfLast = index
		b.last = valueToUpdateWith
	}
}

func (b *calculationValueBuilder) build() (int, error) {
	return strconv.Atoi(b.first + b.last)
}

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
	builder := calculationValueBuilder{}.init(line)

	for i := 0; i < 10; i++ {
		iAsString := strconv.Itoa(i)
		builder.findAndUpdate(iAsString, iAsString)
	}

	if includeDigitWords {
		for digitWord, digit := range digitWordMap {
			builder.findAndUpdate(digitWord, digit)
		}
	}

	return builder.build()
}

func solvePuzzle(lines *bufio.Scanner, includeDigitWords bool) (int, error) {
	sum := 0

	for lines.Scan() {
		calculationValue, err := getCalculationValue(lines.Text(), includeDigitWords)

		if err != nil {
			return 0, err
		}

		sum += calculationValue
	}

	return sum, nil
}

func SolvePuzzle1(lines *bufio.Scanner) (int, error) {
	return solvePuzzle(lines, false)
}

func SolvePuzzle2(lines *bufio.Scanner) (int, error) {
	return solvePuzzle(lines, true)
}
