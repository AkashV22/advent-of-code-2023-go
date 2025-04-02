package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/AkashV22/advent-of-code-2023-go/day01"
	"github.com/AkashV22/advent-of-code-2023-go/day02"
	"github.com/AkashV22/advent-of-code-2023-go/day03"
	"github.com/AkashV22/advent-of-code-2023-go/puzzle"
)

func slogError(err error) slog.Attr {
	return slog.Any("err", err.Error())
}

func solveAllPuzzles(w http.ResponseWriter, puzzleSolvers []puzzle.Solver) {
	slog.Info("Solving all puzzles.")

	results := make(map[string]int)

	for _, puzzleSolver := range puzzleSolvers {
		for puzzleNumber := puzzle.FirstNumber(); puzzleNumber.Ok(); puzzleNumber++ {
			puzzleDay := puzzleSolver.Day()

			inputPath := fmt.Sprintf("input/%02d.txt", puzzleDay)
			file, err := os.Open(inputPath)
			if err != nil {
				slog.Error("Error opening file.", slogError(err))
				http.Error(w, err.Error(), 500)
				return
			}
			defer file.Close()

			input := bufio.NewScanner(file)
			result, err := puzzleSolver.SolvePuzzle(puzzleNumber, input)

			if err != nil {
				slog.Error("Error solving puzzle.", slogError(err))
				http.Error(w, err.Error(), 500)
				return
			}

			results[fmt.Sprintf("Day %v Puzzle %v", puzzleDay, puzzleNumber)] = result
		}
	}

	slog.Info("Solved all puzzles.", slog.Int("total", len(results)))

	resultsJson, err := json.Marshal(results)

	if err != nil {
		slog.Error("Error mashalling results to JSON.", slogError(err))
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write(resultsJson)
}

func main() {
	puzzleSolvers := []puzzle.Solver{
		day01.NewPuzzleSolver(),
		day02.NewPuzzleSolver(),
		day03.NewPuzzleSolver(),
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		solveAllPuzzles(w, puzzleSolvers)
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		slog.Error("Error starting server, exiting...", slogError(err))
		os.Exit(1)
	}
}
