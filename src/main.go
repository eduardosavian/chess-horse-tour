package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

var (
	foundTour   = make(chan [][]int)
	waitGroup   sync.WaitGroup
)

func validateInput(args []string) (int, int, int, int, error) {
	usageMsg    := "Usage: go run main.go <startX> <startY> [<numberThreads> <boardSize>]"
	rangeErrMsg := "Invalid argument. Must be an integer between 0 and %d"

	defaultConcurrency := 4
	defaultBoardSize   := 8

	if len(args) < 3 || len(args) > 5 {
		return 0, 0, 0, 0, fmt.Errorf("invalid number of arguments. %s", usageMsg)
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
			return 0, 0, 0, 0, fmt.Errorf("invalid number of threads\n must be a positive integer")
		}
	}

	boardSize := defaultBoardSize
	if len(args) > 4 {
		boardSize, err = strconv.Atoi(args[4])
		if err != nil || boardSize <= 0 {
			return 0, 0, 0, 0, fmt.Errorf("invalid board size\n must be a positive integer")
		}
	}

	return startX, startY, concurrency, boardSize, nil
}

func isMoveValid(x, y, boardSize int) bool {
	return x >= 0 && x < boardSize && y >= 0 && y < boardSize
}

func findNextMoves(x, y, boardSize int) [][]int {
	possibleMoves := [][]int{
		{-2, -1}, {-1, -2}, {1, -2}, {2, -1},
		{2, 1}, {1, 2}, {-1, 2}, {-2, 1},
	}

	validMoves := [][]int{}

	for _, move := range possibleMoves {
		nextX := x + move[0]
		nextY := y + move[1]

		if isMoveValid(nextX, nextY, boardSize) {
			validMoves = append(validMoves, []int{nextX, nextY})
		}
	}

	return validMoves
}


func main() {
	fmt.Println("Hello, world")

	startX, startY, concurrency, boardSize, err := validateInput(os.Args)
	if err != nil {
		fmt.Println(err)
		return
	}

	waitGroup.Add(concurrency)
	for i := 0; i < concurrency; i++ {
		go tourWorker(startX, startY, boardSize)
	}

	go func() {
		waitGroup.Wait()
		close(foundTour)
	}()

	for tour := range foundTour {
		fmt.Println("Found Knight's Tour:")
		for _, row := range tour {
			for _, value := range row {
				fmt.Printf("%3d ", value)
			}
			fmt.Println()
		}
		return
	}

	fmt.Println("No valid Knight's Tour found after exhausting all attempts.")
}
