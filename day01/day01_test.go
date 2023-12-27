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
		t.Fatalf("Expected %v, received %v.", expected, result)
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
		t.Fatalf("Expected %v, received %v.", expected, result)
	}
}

func TestSolvePuzzle1(t *testing.T) {
	doTestSolvePuzzle(t, 1, 142)
}

func TestSolvePuzzle2(t *testing.T) {
	doTestSolvePuzzle(t, 2, 281)
}
