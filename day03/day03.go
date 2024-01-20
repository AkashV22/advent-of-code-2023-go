package day03

import (
	"bufio"
	"strconv"
	"unicode"

	"github.com/AkashV22/advent-of-code-2023-go/puzzle"
)

type puzzleSolver struct{}

func (*puzzleSolver) Day() int {
	return 3
}

func isSymbol(schematic []string, x, y int) bool {
	if y < 0 || y >= len(schematic) {
		return false
	}

	row := []rune(schematic[y])

	if x < 0 || x >= len(row) {
		return false
	}

	ch := row[x]

	return !unicode.IsDigit(ch) && ch != '.'
}

func isAdjacentSymbolPresent(schematic []string, x, y int) bool {
	return isSymbol(schematic, x-1, y) || // left
		isSymbol(schematic, x-1, y-1) || // top-left
		isSymbol(schematic, x, y-1) || // top
		isSymbol(schematic, x+1, y-1) || // top-right
		isSymbol(schematic, x+1, y) || // right
		isSymbol(schematic, x+1, y+1) || // bottom-right
		isSymbol(schematic, x, y+1) || // bottom
		isSymbol(schematic, x-1, y+1) // bottom-left
}

func (*puzzleSolver) SolvePuzzle(puzzleNumber puzzle.Number, input *bufio.Scanner) (int, error) {
	if puzzleNumber == puzzle.NumberTwo {
		return 0, nil
	}

	var schematic []string
	for input.Scan() {
		schematic = append(schematic, input.Text())
	}

	sum := 0
	for y, row := range schematic {
		adjacentSymbolPresent := false
		partNumberStr := ""
		for x, ch := range row {
			if unicode.IsDigit(ch) {
				partNumberStr += string(ch)
				if isAdjacentSymbolPresent(schematic, x, y) {
					adjacentSymbolPresent = true
				}
			} else {
				if adjacentSymbolPresent {
					if partNumber, err := strconv.Atoi(partNumberStr); err != nil {
						return 0, err
					} else {
						sum += partNumber
					}
				}
				adjacentSymbolPresent = false
				partNumberStr = ""
			}

			if x == len(row)-1 {
				if adjacentSymbolPresent {
					if partNumber, err := strconv.Atoi(partNumberStr); err != nil {
						return 0, err
					} else {
						sum += partNumber
					}
				}
			}
		}
	}

	return sum, nil
}

func NewPuzzleSolver() puzzle.Solver {
	return &puzzleSolver{}
}
