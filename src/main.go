package main

import (
	"fmt"
	"os"
	"strconv"
)

func validateInput(args []string) (int, int, int, int, error) {
	usageMsg    := "Usage: go run main.go <startX> <startY> [<numberThreads> <boardSize>]"
	rangeErrMsg := "Invalid argument. Must be an integer between 0 and %d"

	defaultConcurrency := 4
	defaultBoardSize   := 8

	if len(args) < 3 || len(args) > 5 {
		return 0, 0, 0, 0, fmt.Errorf("Invalid number of arguments. %s", usageMsg)
	}

	startX, err := strconv.Atoi(args[1])
	if err != nil || startX < 0 || startX >= defaultBoardSize {
		return 0, 0, 0, 0, fmt.Errorf(rangeErrMsg, defaultBoardSize-1)
	}

	startY, err := strconv.Atoi(args[2])
	if err != nil || startY < 0 || startY >= defaultBoardSize {
		return 0, 0, 0, 0, fmt.Errorf(rangeErrMsg, defaultBoardSize-1)
	}

	concurrency := defaultConcurrency
	if len(args) > 3 {
		concurrency, err = strconv.Atoi(args[3])
		if err != nil || concurrency <= 0 {
			return 0, 0, 0, 0, fmt.Errorf("Invalid number of threads. Must be a positive integer.")
		}
	}

	boardSize := defaultBoardSize
	if len(args) > 4 {
		boardSize, err = strconv.Atoi(args[4])
		if err != nil || boardSize <= 0 {
			return 0, 0, 0, 0, fmt.Errorf("Invalid board size. Must be a positive integer.")
		}
	}

	return startX, startY, concurrency, boardSize, nil
}

func main() {
	fmt.Println("Hello, world")

	startX, startY, concurrency, boardSize, err := validateInput(os.Args)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Valid input: startX=%d, startY=%d, concurrency=%d, boardSize=%d\n", startX, startY, concurrency, boardSize)

	boardArea := boardSize * boardSize

}
