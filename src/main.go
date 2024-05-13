package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Board struct {
	Rows [][]int `json:"board"`
}

func convertBoardToJSON(board [][]int) ([]byte, error) {
	data := Board{
		Rows: board,
	}
	return json.MarshalIndent(data, "", "  ")
}

func validateInput(args []string) (int, int, int, int, error) {
	if len(args) < 3 || len(args) > 6 {
		return 0, 0, 0, 0, fmt.Errorf("Usage: go run main.go <startX> <startY> [<boardSize> [<timeoutMinutes>]]")
	}

	startX, err := strconv.Atoi(args[1])
	if err != nil {
		return 0, 0, 0, 0, fmt.Errorf("Invalid startX argument: %v", err)
	}

	startY, err := strconv.Atoi(args[2])
	if err != nil {
		return 0, 0, 0, 0, fmt.Errorf("Invalid startY argument: %v", err)
	}


	boardSize := 8 // default board size
	if len(args) > 3 {
		boardSize, err = strconv.Atoi(args[3])
		if err != nil || boardSize <= 0 {
			return 0, 0, 0, 0, fmt.Errorf("Invalid board size, must be a positive integer")
		}
	}

	if startX < 0 || startX >= boardSize {
		return 0, 0, 0, 0, fmt.Errorf("Invalid startX value, must be between 0 and %d (exclusive)", boardSize)
	}

	if startY < 0 || startY >= boardSize {
		return 0, 0, 0, 0, fmt.Errorf("Invalid startY value, must be between 0 and %d (exclusive)", boardSize)
	}

	timeout := 5
	if len(args) > 4 {
		timeout, err = strconv.Atoi(args[4])
		if err != nil || timeout <= 0 {
			return 0, 0, 0, 0, fmt.Errorf("Invalid timeout value, must be a positive integer")
		}
	}

	return startX, startY, boardSize, timeout, nil
}

func main() {
	startX, startY, boardSize, timeout, err := validateInput(os.Args)
	if err != nil {
		fmt.Println(err)
		return
	}

	board := make([][]int, boardSize)
	for i := range board {
		board[i] = make([]int, boardSize)
	}

	timer := time.After(time.Duration(timeout) * time.Minute)

	go func() {
		<-timer
		fmt.Printf("\nTimeout reached after %d minutes. Exiting...\n", timeout)
		os.Exit(0)
	}()


	method := "warnsdorff"

	if !backtrack(board, 1, startX, startY, boardSize, method) {
		fmt.Println("No valid Knight's Tour found after exhausting all attempts.")
		return
	}

	if jsonBytes, err := convertBoardToJSON(board); err == nil {
		fmt.Println(string(jsonBytes))
	} else {
		fmt.Println("Error converting board to JSON:", err)
	}
}

