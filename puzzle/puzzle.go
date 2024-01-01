package puzzle

import "bufio"

type Number int

func (number Number) Ok() bool {
	switch number {
	case One, Two:
		return true
	}
	return false
}

const (
	One Number = iota + 1
	Two
)

func FirstNumber() Number {
	return One
}

type Solver interface {
	Day() int

	SolvePuzzle(puzzleNumber Number, input *bufio.Scanner) (int, error)
}
