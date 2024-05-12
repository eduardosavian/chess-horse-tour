package main

import (
	"fmt"
	"os"
  "strconv"
)

var(
  boardSize = 8
  boardArea   = boardSize * boardSize
  concurrency = 10
)

func main() {
  fmt.Println("Hello, world")

  if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <startX> <startY>")
		return
	}

	startX, err := strconv.Atoi(os.Args[1])
	if err != nil || startX < 0 || startX >= boardSize {
		fmt.Println("Invalid startX. Must be an integer between 0 and", boardSize-1)
		return
	}

	startY, err := strconv.Atoi(os.Args[2])
	if err != nil || startY < 0 || startY >= boardSize {
		fmt.Println("Invalid startY. Must be an integer between 0 and", boardSize-1)
		return
	}
}