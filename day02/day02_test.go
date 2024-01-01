package day02

import (
	"testing"

	"github.com/AkashV22/advent-of-code-2023-go/puzzletest"
)

func TestPuzzleSolver(t *testing.T) {
	test := puzzletest.PuzzleSolverTest{
		T:                       t,
		NewPuzzleSolver:         NewPuzzleSolver,
		MultipleInputFiles:      false,
		ExpectedDay:             2,
		ExpectedPuzzleOneResult: 8,
		ExpectedPuzzleTwoResult: 2286,
	}

	test.Run()
}
