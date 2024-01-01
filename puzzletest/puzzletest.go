package puzzletest

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/AkashV22/advent-of-code-2023-go/puzzle"
)

type testCase interface {
	name() string

	run(t *testing.T, puzzleSolver puzzle.Solver)
}

type dayTestCase int

func (tc dayTestCase) name() string {
	return "Day"
}

func (tc dayTestCase) run(t *testing.T, puzzleSolver puzzle.Solver) {
	expected := int(tc)
	result := puzzleSolver.Day()
	if expected != result {
		t.Errorf("Expected %v, received %v.", expected, result)
	}
}

type solvePuzzleSubTestCase struct {
	puzzleNumber puzzle.Number
	expected     int
}

func (tc solvePuzzleSubTestCase) run(t *testing.T, puzzleSolver puzzle.Solver, multipleInputFiles bool) {
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

type solvePuzzleTestCase struct {
	multipleInputFiles      bool
	expectedPuzzleOneResult int
	expectedPuzzleTwoResult int
}

func (tc solvePuzzleTestCase) name() string {
	return "SolvePuzzle"
}

func (tc solvePuzzleTestCase) run(t *testing.T, puzzleSolver puzzle.Solver) {
	subTestCases := [2]solvePuzzleSubTestCase{
		{puzzleNumber: puzzle.One, expected: tc.expectedPuzzleOneResult},
		{puzzleNumber: puzzle.Two, expected: tc.expectedPuzzleTwoResult},
	}

	for _, subTc := range subTestCases {
		t.Run(strconv.Itoa(int(subTc.puzzleNumber)), func(t *testing.T) {
			subTc.run(t, puzzleSolver, tc.multipleInputFiles)
		})
	}
}

type PuzzleSolverTest struct {
	T                       *testing.T
	NewPuzzleSolver         func() puzzle.Solver
	MultipleInputFiles      bool
	ExpectedDay             int
	ExpectedPuzzleOneResult int
	ExpectedPuzzleTwoResult int
}

func (test PuzzleSolverTest) Run() {
	testCases := [2]testCase{
		dayTestCase(test.ExpectedDay),
		solvePuzzleTestCase{
			multipleInputFiles:      test.MultipleInputFiles,
			expectedPuzzleOneResult: test.ExpectedPuzzleOneResult,
			expectedPuzzleTwoResult: test.ExpectedPuzzleTwoResult,
		},
	}

	puzzleSolver := test.NewPuzzleSolver()

	for _, tc := range testCases {
		test.T.Run(tc.name(), func(t *testing.T) {
			tc.run(t, puzzleSolver)
		})
	}
}
