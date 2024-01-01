package day02

import (
	"testing"

	"github.com/AkashV22/advent-of-code-2023-go/puzzletest"
)

func TestGetDay(t *testing.T) {
	puzzletest.TestGetDay(t, NewPuzzleSolver(), 2)
}

func TestSolvePuzzle(t *testing.T) {
	puzzletest.TestSolvePuzzle(t, NewPuzzleSolver(), false, 8, 2286)
}
