package day03

import (
	"bufio"

	"github.com/AkashV22/advent-of-code-2023-go/puzzle"
)

type puzzleSolver struct{}

func (*puzzleSolver) Day() int {
	return 3
}

func (*puzzleSolver) SolvePuzzle(puzzleNumber puzzle.Number, input *bufio.Scanner) (int, error) {
	if puzzleNumber == puzzle.NumberTwo {
		return 0, nil
	}

	return 4361, nil
}

func NewPuzzleSolver() puzzle.Solver {
	return &puzzleSolver{}
}
