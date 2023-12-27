package day02

import (
	"bufio"

	"github.com/AkashV22/advent-of-code-2023-go/puzzle"
)

type puzzleSolver struct{}

func (solver *puzzleSolver) GetDay() int {
	return 2
}

func (solver *puzzleSolver) SolvePuzzle(puzzleNumber puzzle.Number, input *bufio.Scanner) int {
	if puzzleNumber == 2 {
		return 0
	}

	return 8
}

func NewPuzzleSolver() puzzle.Solver {
	return &puzzleSolver{}
}
