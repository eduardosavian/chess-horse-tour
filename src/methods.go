package main

import (
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

func backtrack(board [][]int, moveNum, x, y, boardSize int) bool {
	if moveNum == boardSize*boardSize {
		board[x][y] = moveNum
		return true
	}

	nextMoves := findNextMoves(x, y, boardSize)

	for _, move := range nextMoves {
		nextX, nextY := move[0], move[1]
		if board[nextX][nextY] == 0 {
			board[nextX][nextY] = moveNum

			if backtrack(board, moveNum+1, nextX, nextY, boardSize) {
				return true
			}

			board[nextX][nextY] = 0
		}
	}

	return false
}
