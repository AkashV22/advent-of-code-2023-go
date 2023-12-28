package day02

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/AkashV22/advent-of-code-2023-go/puzzle"
)

func passThroughValueIfGameIsPossible(cubeSets []string, value int) (int, error) {
	for _, cubeSet := range cubeSets {
		cubes := strings.Split(cubeSet, ", ")

		for _, cube := range cubes {
			cubeData := strings.Split(cube, " ")
			total, err := strconv.Atoi(cubeData[0])
			if err != nil {
				return 0, err
			}

			colour := cubeData[1]

			if (colour == "red" && total > 12) || (colour == "green" && total > 13) || (colour == "blue" && total > 14) {
				return 0, nil
			}
		}
	}

	return value, nil
}

func calculatePowerOfCubes(cubeSets []string) (int, error) {
	var maxRedTotal, maxGreenTotal, maxBlueTotal int

	for _, cubeSet := range cubeSets {
		cubes := strings.Split(cubeSet, ", ")

		for _, cube := range cubes {
			cubeData := strings.Split(cube, " ")
			total, err := strconv.Atoi(cubeData[0])
			if err != nil {
				return 0, err
			}

			colour := cubeData[1]

			switch {
			case colour == "red" && total > maxRedTotal:
				maxRedTotal = total
			case colour == "green" && total > maxGreenTotal:
				maxGreenTotal = total
			case colour == "blue" && total > maxBlueTotal:
				maxBlueTotal = total
			}
		}
	}

	return maxRedTotal * maxGreenTotal * maxBlueTotal, nil
}

type puzzleSolver struct{}

func (solver *puzzleSolver) GetDay() int {
	return 2
}

func (solver *puzzleSolver) SolvePuzzle(puzzleNumber puzzle.Number, input *bufio.Scanner) (int, error) {
	sum := 0

	for input.Scan() {
		game := strings.Replace(input.Text(), "Game ", "", 1)

		gameData := strings.Split(game, ": ")

		id, err := strconv.Atoi(gameData[0])
		if err != nil {
			return 0, err
		}

		cubeSets := strings.Split(gameData[1], "; ")

		var handleCubeSets func([]string) (int, error)
		switch puzzleNumber {
		case puzzle.One:
			handleCubeSets = func(cubeSets []string) (int, error) {
				return passThroughValueIfGameIsPossible(cubeSets, id)
			}
		case puzzle.Two:
			handleCubeSets = calculatePowerOfCubes
		default:
			return 0, errors.New(fmt.Sprintf("Invalid puzzle number: %v.", puzzleNumber))
		}

		if result, err := handleCubeSets(cubeSets); err != nil {
			return 0, err
		} else {
			sum += result
		}
	}

	return sum, nil
}

func NewPuzzleSolver() puzzle.Solver {
	return &puzzleSolver{}
}
