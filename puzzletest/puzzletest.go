package puzzletest

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/AkashV22/advent-of-code-2023-go/puzzle"
)

func TestDay(t *testing.T, puzzleSolver puzzle.Solver, expected int) {
	result := puzzleSolver.Day()
	if expected != result {
		t.Errorf("Expected %v, received %v.", expected, result)
	}
}

type solvePuzzleTestCase struct {
	puzzleNumber puzzle.Number
	expected     int
}

func (tc solvePuzzleTestCase) run(t *testing.T, puzzleSolver puzzle.Solver, multipleInputFiles bool) {
	puzzleNumber := tc.puzzleNumber

	var inputPath string
	if multipleInputFiles {
		inputPath = fmt.Sprintf("example%v.txt", puzzleNumber)
	} else {
		inputPath = "example.txt"
	}

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

	if expected := tc.expected; expected != result {
		t.Errorf("Expected %v, received %v.", expected, result)
	}
}

func TestSolvePuzzle(t *testing.T, puzzleSolver puzzle.Solver, multipleInputFiles bool, puzzleOneExpected int, puzzleTwoExpected int) {
	testCases := [2]solvePuzzleTestCase{
		{puzzleNumber: puzzle.One, expected: puzzleOneExpected},
		{puzzleNumber: puzzle.Two, expected: puzzleTwoExpected},
	}

	for _, tc := range testCases {
		t.Run(strconv.Itoa(int(tc.puzzleNumber)), func(t *testing.T) {
			tc.run(t, puzzleSolver, multipleInputFiles)
		})
	}
}
