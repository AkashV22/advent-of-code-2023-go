package day01

import (
	"bufio"
	"os"
	"testing"
)

func TestSolvePuzzle1(t *testing.T) {
	file, err := os.Open("example1.txt")
	defer file.Close()

	if err != nil {
		t.Fatal("Error opening file.", err)
	}

	lines := bufio.NewScanner(file)

	result, err := SolvePuzzle1(lines)

	if err != nil {
		t.Fatal("Error solving puzzle.", err)
	}

	if "142" != result {
		t.Fatalf("Expected 142, received %v.", result)
	}
}
