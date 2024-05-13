package main

import (
	"sort"
)

func isMoveValid(x, y, boardSize int) bool {
	return x >= 0 && x < boardSize && y >= 0 && y < boardSize
}

type Move struct {
	X, Y     int
	Priority int
}

func findNextMoves(x, y, boardSize int, method string) []Move {
	possibleMoves := [][]int{
		{-2, -1}, {-1, -2}, {1, -2}, {2, -1},
		{2, 1}, {1, 2}, {-1, 2}, {-2, 1},
	}

	validMoves := []Move{}

	for _, move := range possibleMoves {
		nextX := x + move[0]
		nextY := y + move[1]

		if isMoveValid(nextX, nextY, boardSize) {
			validMoves = append(validMoves, Move{nextX, nextY, 0})
		}
	}

	if method == "warnsdorff" {
		for i := range validMoves {
			move := &validMoves[i]
			move.Priority = len(findNextMoves(move.X, move.Y, boardSize, "default"))
		}

		sort.Slice(validMoves, func(i, j int) bool {
			return validMoves[i].Priority < validMoves[j].Priority
		})
	}

	return validMoves
}


func backtrack(board [][]int, moveNum, x, y, boardSize int, method string) bool {
	board[x][y] = moveNum

	if moveNum == boardSize*boardSize {
		return true
	}

	nextMoves := findNextMoves(x, y, boardSize, method)

	for _, move := range nextMoves {
		nextX, nextY := move.X, move.Y
		if board[nextX][nextY] == 0 {
			if backtrack(board, moveNum+1, nextX, nextY, boardSize, method) {
				return true
			}
		}
	}

	board[x][y] = 0 
	return false
}
