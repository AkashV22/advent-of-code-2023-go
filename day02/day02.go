package day02

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/AkashV22/advent-of-code-2023-go/puzzle"
)

func processCubeSets(cubeSets []string, processCubeTotalAndColour func(int, string) bool) error {
	for _, cubeSet := range cubeSets {
		cubes := strings.Split(cubeSet, ", ")

		for _, cube := range cubes {
			cubeData := strings.Split(cube, " ")
			total, err := strconv.Atoi(cubeData[0])
			if err != nil {
				return err
			}

			colour := cubeData[1]

			if returnImmediately := processCubeTotalAndColour(total, colour); returnImmediately {
				return nil
			}
		}
	}

	return nil
}

func passThroughValueIfGameIsPossible(cubeSets []string, value int) (int, error) {
	isGamePossible := true

	processCubeSets(cubeSets, func(total int, colour string) bool {
		if (colour == "red" && total > 12) || (colour == "green" && total > 13) || (colour == "blue" && total > 14) {
			isGamePossible = false
			return true
		}

		return false
	})

	if isGamePossible {
		return value, nil
	}
	return 0, nil
}

func calculatePowerOfCubes(cubeSets []string) (int, error) {
	var maxRedTotal, maxGreenTotal, maxBlueTotal int

	processCubeSets(cubeSets, func(total int, colour string) bool {
		switch {
		case colour == "red" && total > maxRedTotal:
			maxRedTotal = total
		case colour == "green" && total > maxGreenTotal:
			maxGreenTotal = total
		case colour == "blue" && total > maxBlueTotal:
			maxBlueTotal = total
		}

		return false
	})

	return maxRedTotal * maxGreenTotal * maxBlueTotal, nil
}

type puzzleSolver struct{}

func (solver *puzzleSolver) Day() int {
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
		case puzzle.NumberOne:
			handleCubeSets = func(cubeSets []string) (int, error) {
				return passThroughValueIfGameIsPossible(cubeSets, id)
			}
		case puzzle.NumberTwo:
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
