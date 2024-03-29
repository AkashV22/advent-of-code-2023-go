package puzzletest

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/AkashV22/advent-of-code-2023-go/puzzle"
)

func testDay(t *testing.T, puzzleSolver puzzle.Solver, expected int) {
	result := puzzleSolver.Day()
	if expected != result {
		t.Errorf("Expected %v, received %v.", expected, result)
	}
}

type solvePuzzleTestCase struct {
	puzzleNumber puzzle.Number
	expected     int
}

func (tc solvePuzzleTestCase) run(t *testing.T, puzzleSolver puzzle.Solver) {
	puzzleNumber := tc.puzzleNumber

	file, err := os.Open(fmt.Sprintf("example%v.txt", puzzleNumber))
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

type solvePuzzleSubTest struct {
	puzzleSolver            puzzle.Solver
	expectedPuzzleOneResult int
	expectedPuzzleTwoResult int
}

func (test solvePuzzleSubTest) run(t *testing.T) {
	testCases := [2]solvePuzzleTestCase{
		{puzzleNumber: puzzle.NumberOne, expected: test.expectedPuzzleOneResult},
		{puzzleNumber: puzzle.NumberTwo, expected: test.expectedPuzzleTwoResult},
	}

	for _, tc := range testCases {
		t.Run(strconv.Itoa(int(tc.puzzleNumber)), func(t *testing.T) {
			tc.run(t, test.puzzleSolver)
		})
	}
}

type PuzzleSolverTest struct {
	T                       *testing.T
	NewPuzzleSolver         func() puzzle.Solver
	ExpectedDay             int
	ExpectedPuzzleOneResult int
	ExpectedPuzzleTwoResult int
}

func (test PuzzleSolverTest) Run() {
	puzzleSolver := test.NewPuzzleSolver()

	test.T.Run("Day", func(t *testing.T) {
		testDay(t, puzzleSolver, test.ExpectedDay)
	})

	test.T.Run("SolvePuzzle", func(t *testing.T) {
		solvePuzzleSubTest := solvePuzzleSubTest{
			puzzleSolver:            puzzleSolver,
			expectedPuzzleOneResult: test.ExpectedPuzzleOneResult,
			expectedPuzzleTwoResult: test.ExpectedPuzzleTwoResult,
		}

		solvePuzzleSubTest.run(t)
	})
}
