package day01

import (
	"bufio"
	"math"
	"strconv"
	"strings"

	"github.com/AkashV22/advent-of-code-2023-go/puzzle"
)

type calculationValueBuilder struct {
	line         string
	first        int
	last         int
	indexOfFirst int
	indexOfLast  int
}

func (b *calculationValueBuilder) findAndUpdate(valueToFind string, valueToUpdateWith int) {
	if index := strings.Index(b.line, valueToFind); index >= 0 && index < b.indexOfFirst {
		b.indexOfFirst = index
		b.first = valueToUpdateWith
	}

	if index := strings.LastIndex(b.line, valueToFind); index >= 0 && index > b.indexOfLast {
		b.indexOfLast = index
		b.last = valueToUpdateWith
	}
}

func (b *calculationValueBuilder) build() int {
	return (b.first * 10) + b.last
}

func newCalculationValueBuilder(line string) calculationValueBuilder {
	return calculationValueBuilder{
		line:         line,
		indexOfFirst: math.MaxInt,
		indexOfLast:  math.MinInt,
	}
}

var digitWordMap = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func getCalculationValue(line string, puzzleNumber puzzle.Number) int {
	builder := newCalculationValueBuilder(line)

	for i := 0; i < 10; i++ {
		builder.findAndUpdate(strconv.Itoa(i), i)
	}

	if puzzleNumber == puzzle.Two {
		for digitWord, digit := range digitWordMap {
			builder.findAndUpdate(digitWord, digit)
		}
	}

	return builder.build()
}

type puzzleSolver struct{}

func (solver *puzzleSolver) Day() int {
	return 1
}

func (solver *puzzleSolver) SolvePuzzle(puzzleNumber puzzle.Number, input *bufio.Scanner) (int, error) {
	sum := 0

	for input.Scan() {
		sum += getCalculationValue(input.Text(), puzzleNumber)
	}

	return sum, nil
}

func NewPuzzleSolver() puzzle.Solver {
	return &puzzleSolver{}
}
