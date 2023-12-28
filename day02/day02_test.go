package day02

import (
	"testing"

	"github.com/AkashV22/advent-of-code-2023-go/testutils"
)

func TestGetDay(t *testing.T) {
	testutils.TestGetDay(t, NewPuzzleSolver(), 2)
}

func TestSolvePuzzle(t *testing.T) {
	testutils.TestSolvePuzzle(t, NewPuzzleSolver(), false, 8, 2286)
}
