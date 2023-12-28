package day01

import (
	"testing"

	"github.com/AkashV22/advent-of-code-2023-go/testutils"
)

func TestGetDay(t *testing.T) {
	testutils.TestGetDay(t, NewPuzzleSolver(), 1)
}

func TestSolvePuzzle(t *testing.T) {
	testutils.TestSolvePuzzle(t, NewPuzzleSolver(), true, 142, 281)
}
