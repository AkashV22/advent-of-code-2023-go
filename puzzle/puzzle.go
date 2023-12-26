package puzzle

import "bufio"

type Number int

func (number Number) Ok() bool {
	switch number {
	case One:
	case Two:
		return true
	}
	return false
}

const (
	One Number = iota + 1
	Two
)

type Solver interface {
	GetDay() int

	SolvePuzzle(puzzleNumber Number, input *bufio.Scanner) (int, error)
}
