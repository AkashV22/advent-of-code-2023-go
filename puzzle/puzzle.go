package puzzle

import "bufio"

type Number int

func (number Number) Ok() bool {
	switch number {
	case NumberOne, NumberTwo:
		return true
	}
	return false
}

const (
	NumberOne Number = iota + 1
	NumberTwo
)

func FirstNumber() Number {
	return NumberOne
}

type Solver interface {
	Day() int

	SolvePuzzle(puzzleNumber Number, input *bufio.Scanner) (int, error)
}
