package testutils

import (
	"bufio"
	"fmt"
	"os"
	"testing"

	"github.com/AkashV22/advent-of-code-2023-go/puzzle"
)

func TestGetDay(t *testing.T, puzzleSolver puzzle.Solver, expected int) {
	result := puzzleSolver.GetDay()
	if expected != result {
		t.Errorf("Expected %v, received %v.", expected, result)
	}
}

func doTestSolvePuzzle(t *testing.T, puzzleSolver puzzle.Solver, puzzleNumber puzzle.Number, expected int) {
	inputPath := fmt.Sprintf("example%v.txt", puzzleNumber)
	file, err := os.Open(inputPath)
	defer file.Close()

	if err != nil {
		t.Fatal("Error opening file.", err)
	}

	input := bufio.NewScanner(file)

	var result int
	result, err = puzzleSolver.SolvePuzzle(puzzleNumber, input)

	if err != nil {
		t.Fatal("Error solving puzzle.", err)
	}

	if expected != result {
		t.Errorf("Expected %v, received %v.", expected, result)
	}
}

func TestSolvePuzzle(t *testing.T, puzzleSolver puzzle.Solver, puzzleOneExpected int, puzzleTwoExpected int) {
	testCases := [2]struct {
		puzzleNumber puzzle.Number
		expected     int
	}{
		{puzzleNumber: puzzle.One, expected: puzzleOneExpected},
		{puzzleNumber: puzzle.Two, expected: puzzleTwoExpected},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Puzzle %v", tc.puzzleNumber), func(t *testing.T) {
			doTestSolvePuzzle(t, puzzleSolver, tc.puzzleNumber, tc.expected)
		})
	}
}
