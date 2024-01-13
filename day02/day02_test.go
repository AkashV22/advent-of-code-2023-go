package day02_test

import (
	"testing"

	"github.com/AkashV22/advent-of-code-2023-go/day02"
	"github.com/AkashV22/advent-of-code-2023-go/puzzletest"
)

func TestPuzzleSolver(t *testing.T) {
	test := puzzletest.PuzzleSolverTest{
		T:                       t,
		NewPuzzleSolver:         day02.NewPuzzleSolver,
		ExpectedDay:             2,
		ExpectedPuzzleOneResult: 8,
		ExpectedPuzzleTwoResult: 2286,
	}

	test.Run()
}
