package day01

import (
	"bufio"
	"fmt"
	"os"
	"testing"

	"github.com/AkashV22/advent-of-code-2023-go/puzzle"
)

func TestGetDay(t *testing.T) {
	expected := 1
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

	var result int
	result, err = NewPuzzleSolver().SolvePuzzle(puzzleNumber, input)

	if err != nil {
		t.Fatal("Error solving puzzle.", err)
	}

	if expected != result {
		t.Errorf("Expected %v, received %v.", expected, result)
	}
}

func TestSolvePuzzle(t *testing.T) {
	testCases := [2]struct {
		puzzleNumber puzzle.Number
		expected     int
	}{
		{puzzleNumber: puzzle.One, expected: 142},
		{puzzleNumber: puzzle.Two, expected: 281},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Puzzle %v", tc.puzzleNumber), func(t *testing.T) {
			doTestSolvePuzzle(t, tc.puzzleNumber, tc.expected)
		})
	}
}
