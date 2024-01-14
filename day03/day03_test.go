package day03_test

import (
	"testing"

	"github.com/AkashV22/advent-of-code-2023-go/day03"
	"github.com/AkashV22/advent-of-code-2023-go/puzzletest"
)

func TestPuzzleSolver(t *testing.T) {
	test := puzzletest.PuzzleSolverTest{
		T:                       t,
		NewPuzzleSolver:         day03.NewPuzzleSolver,
		ExpectedDay:             3,
		ExpectedPuzzleOneResult: 4361,
		ExpectedPuzzleTwoResult: 0,
	}

	test.Run()
}
