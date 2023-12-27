package day02

import (
	"bufio"
	"log/slog"
	"strconv"
	"strings"

	"github.com/AkashV22/advent-of-code-2023-go/puzzle"
)

func isGamePossible(cubeSets []string) bool {
	for _, cubeSet := range cubeSets {
		cubes := strings.Split(cubeSet, ", ")

		for _, cube := range cubes {
			cubeData := strings.Split(cube, " ")
			total, err := strconv.Atoi(cubeData[0])
			if err != nil {
				slog.Error("Error converting cube total to int.", err)
				return false
			}

			colour := cubeData[1]

			if (colour == "red" && total > 12) || (colour == "green" && total > 13) || (colour == "blue" && total > 14) {
				return false
			}
		}
	}

	return true
}

type puzzleSolver struct{}

func (solver *puzzleSolver) GetDay() int {
	return 2
}

func (solver *puzzleSolver) SolvePuzzle(puzzleNumber puzzle.Number, input *bufio.Scanner) int {
	if puzzleNumber == 2 {
		return 0
	}

	sum := 0

	for input.Scan() {
		game := strings.Replace(input.Text(), "Game ", "", 1)

		gameData := strings.Split(game, ": ")

		id, err := strconv.Atoi(gameData[0])
		if err != nil {
			slog.Error("Error converting ID to int.", err)
			continue
		}

		cubeSets := strings.Split(gameData[1], "; ")

		if isGamePossible(cubeSets) {
			sum += id
		}
	}

	return sum
}

func NewPuzzleSolver() puzzle.Solver {
	return &puzzleSolver{}
}
