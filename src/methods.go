package main

import (
	"math/rand"
	"sync"
)

var (
	foundTour  = make(chan [][]int)
	stopSearch = make(chan struct{})
	waitGroup  sync.WaitGroup
	boardMutex sync.Mutex
)


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

func tourWorker(startX, startY, boardSize int) {
	defer waitGroup.Done()

	board := make([][]int, boardSize)
	for i := range board {
		board[i] = make([]int, boardSize)
	}

	boardMutex.Lock()
	board[startX][startY] = 1
	boardMutex.Unlock()

	if backtrack(board, 1, startX, startY, boardSize) {
		select {
		case foundTour <- board:
			close(stopSearch)
		default:
		}
	}
}

func backtrack(board [][]int, moveNum, x, y, boardSize int) bool {
	if moveNum == boardSize*boardSize {
		return true
	}

	select {
	case <-stopSearch:
		return false
	default:
	}

	nextMoves := findNextMoves(x, y, boardSize)
	rand.Shuffle(len(nextMoves), func(i, j int) {
		nextMoves[i], nextMoves[j] = nextMoves[j], nextMoves[i]
	})

	for _, move := range nextMoves {
		nextX, nextY := move[0], move[1]
		if board[nextX][nextY] == 0 {
			boardMutex.Lock()
			board[nextX][nextY] = moveNum + 1
			boardMutex.Unlock()

			if backtrack(board, moveNum+1, nextX, nextY, boardSize) {
				return true
			}

			boardMutex.Lock()
			board[nextX][nextY] = 0
			boardMutex.Unlock()
		}
	}

	return false
}