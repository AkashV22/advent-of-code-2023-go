package day01

import (
	"testing"

	"github.com/AkashV22/advent-of-code-2023-go/puzzletest"
)

func TestDay(t *testing.T) {
	puzzletest.TestDay(t, NewPuzzleSolver(), 1)
}

func TestSolvePuzzle(t *testing.T) {
	puzzletest.TestSolvePuzzle(t, NewPuzzleSolver(), true, 142, 281)
}
