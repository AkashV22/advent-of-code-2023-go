package day01_test

import (
	"testing"

	"github.com/AkashV22/advent-of-code-2023-go/day01"
	"github.com/AkashV22/advent-of-code-2023-go/puzzletest"
)

func TestPuzzleSolver(t *testing.T) {
	test := puzzletest.PuzzleSolverTest{
		T:                       t,
		NewPuzzleSolver:         day01.NewPuzzleSolver,
		MultipleInputFiles:      true,
		ExpectedDay:             1,
		ExpectedPuzzleOneResult: 142,
		ExpectedPuzzleTwoResult: 281,
	}

	test.Run()
}
