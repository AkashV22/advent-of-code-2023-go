package day01

import (
	"bufio"
	"os"
	"testing"
)

func doTestSolvePuzzle(t *testing.T, puzzle int, expected string) {
	var inputPath string
	var solvePuzzle func(*bufio.Scanner) (string, error)
	switch puzzle {
	case 1:
		inputPath = "example1.txt"
		solvePuzzle = SolvePuzzle1
	case 2:
		inputPath = "example2.txt"
		solvePuzzle = SolvePuzzle2
	default:
		t.Fatalf("Unknown puzzle: %v.", puzzle)
	}

	file, err := os.Open(inputPath)
	defer file.Close()

	if err != nil {
		t.Fatal("Error opening file.", err)
	}

	lines := bufio.NewScanner(file)

	var result string
	result, err = solvePuzzle(lines)

	if err != nil {
		t.Fatal("Error solving puzzle.", err)
	}

	if expected != result {
		t.Fatalf("Expected %v, received %v.", expected, result)
	}
}

func TestSolvePuzzle1(t *testing.T) {
	doTestSolvePuzzle(t, 1, "142")
}

func TestSolvePuzzle2(t *testing.T) {
	doTestSolvePuzzle(t, 2, "281")
}
