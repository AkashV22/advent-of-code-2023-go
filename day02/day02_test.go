package day02

import (
	"bufio"
	"fmt"
	"os"
	"testing"

	"github.com/AkashV22/advent-of-code-2023-go/puzzle"
)

func TestGetDay(t *testing.T) {
	expected := 2
	result := NewPuzzleSolver().GetDay()
	if expected != result {
		t.Errorf("Expected %v, received %v.", expected, result)
	}
}

func doTestSolvePuzzle(t *testing.T, puzzleNumber puzzle.Number, expected int) {
	inputPath := fmt.Sprintf("example%v.txt", puzzleNumber)
	file, err := os.Open(inputPath)
	defer file.Close()

	if err != nil {
		t.Fatal("Error opening file.", err)
	}

	input := bufio.NewScanner(file)

	result := NewPuzzleSolver().SolvePuzzle(puzzleNumber, input)

	if expected != result {
		t.Errorf("Expected %v, received %v.", expected, result)
	}
}

func TestSolvePuzzle(t *testing.T) {
	testCases := [2]struct {
		puzzleNumber puzzle.Number
		expected     int
	}{
		{puzzleNumber: puzzle.One, expected: 8},
		{puzzleNumber: puzzle.Two, expected: 0},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Puzzle %v", tc.puzzleNumber), func(t *testing.T) {
			doTestSolvePuzzle(t, tc.puzzleNumber, tc.expected)
		})
	}
}