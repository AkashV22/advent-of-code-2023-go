package day02

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/AkashV22/advent-of-code-2023-go/puzzle"
)

func isGamePossible(cubeSets []string) (bool, error) {
	for _, cubeSet := range cubeSets {
		cubes := strings.Split(cubeSet, ", ")

		for _, cube := range cubes {
			cubeData := strings.Split(cube, " ")
			total, err := strconv.Atoi(cubeData[0])
			if err != nil {
				return false, err
			}

			colour := cubeData[1]

			if (colour == "red" && total > 12) || (colour == "green" && total > 13) || (colour == "blue" && total > 14) {
				return false, nil
			}
		}
	}

	return true, nil
}

type puzzleSolver struct{}

func (solver *puzzleSolver) GetDay() int {
	return 2
}

func (solver *puzzleSolver) SolvePuzzle(puzzleNumber puzzle.Number, input *bufio.Scanner) (int, error) {
	if puzzleNumber == 2 {
		return 0, nil
	}

	sum := 0

	for input.Scan() {
		game := strings.Replace(input.Text(), "Game ", "", 1)

		gameData := strings.Split(game, ": ")

		id, err := strconv.Atoi(gameData[0])
		if err != nil {
			return 0, err
		}

		cubeSets := strings.Split(gameData[1], "; ")

		isGamePossible, err := isGamePossible(cubeSets)
		if err != nil {
			return 0, err
		}

		if isGamePossible {
			sum += id
		}
	}

	return sum, nil
}

func NewPuzzleSolver() puzzle.Solver {
	return &puzzleSolver{}
}
