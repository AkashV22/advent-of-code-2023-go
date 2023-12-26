package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/AkashV22/advent-of-code-2023-go/day01"
)

type solverInfo struct {
	day       int
	puzzle    int
	inputPath string
	solver    func(*bufio.Scanner) (string, error)
}

func solvePuzzles(solvers []solverInfo) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		results := make(map[string]string)

		for _, solverInfo := range solvers {
			file, err := os.Open(solverInfo.inputPath)
			defer file.Close()

			lines := bufio.NewScanner(file)

			result, err := solverInfo.solver(lines)

			if err != nil {
				slog.Error("Error solving puzzle.", err)
				http.Error(w, err.Error(), 500)
				return
			}

			results[fmt.Sprintf("Day %v Puzzle %v", solverInfo.day, solverInfo.puzzle)] = result
		}

		resultsJson, err := json.Marshal(results)

		if err != nil {
			slog.Error("Error mashalling results to JSON.", err)
			http.Error(w, err.Error(), 500)
			return
		}

		w.Write(resultsJson)

	}
}

func main() {
	solvers := []solverInfo{
		{day: 1, puzzle: 1, inputPath: "day01/input.txt", solver: day01.SolvePuzzle1},
		{day: 1, puzzle: 2, inputPath: "day01/input.txt", solver: day01.SolvePuzzle2},
	}

	http.HandleFunc("/", solvePuzzles(solvers))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		slog.Error("Error starting server, exiting...", err)
		os.Exit(1)
	}
}
