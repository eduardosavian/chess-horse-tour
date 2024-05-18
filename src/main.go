package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args

	if len(args) < 5 {
		fmt.Println("Usage: <startX> <startY> <boardSize> <algorithm>")
		return
	}

	startX, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Invalid startX:", err)
		return
	}

	startY, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("Invalid startY:", err)
		return
	}

	boardSize, err := strconv.Atoi(args[3])
	if err != nil || boardSize <= 0 {
		fmt.Println("Invalid board size. Must be a positive integer.")
		return
	}

	algorithm := args[4]
	if algorithm != "warnsdorff" && algorithm != "backtrack" && algorithm != "highDegree" && algorithm != "shuffle" {
		fmt.Println("Invalid algorithm. Must be either 'warnsdorff', 'backtrack', 'highDegree', or 'shuffle'.")
		return
	}


    board := make([][]int, boardSize)
	for i := range board {
		board[i] = make([]int, boardSize)
	}

	if !backtrackWithMethod(board, 1, startX, startY, boardSize, algorithm)  {
		fmt.Println("No valid Knight's tour found.")
	}
    boardJson, err := convertBoardToJSON(board)
    if err != nil {
        fmt.Println("Error converting board to JSON:", err)
        return
    }

    fmt.Println(string(boardJson))
}