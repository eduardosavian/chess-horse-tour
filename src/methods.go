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

func findNextMoves(x, y, boardSize int, method string) [][]int {
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

	if method == "warnsdorff" {
		moves := make([]Move, 0)

		for _, move := range validMoves {
			nextX, nextY := move[0], move[1]
			validMovesFromNext := len(findNextMoves(nextX, nextY, boardSize, "default"))
			moves = append(moves, Move{nextX, nextY, validMovesFromNext})
		}

		sort.Slice(moves, func(i, j int) bool {
			return moves[i].Priority < moves[j].Priority
		})

		sortedMoves := [][]int{}
		for _, move := range moves {
			sortedMoves = append(sortedMoves, []int{move.X, move.Y})
		}

		return sortedMoves
	}

	return validMoves
}

func backtrack(board [][]int, moveNum, x, y, boardSize int, method string) bool {
	if moveNum == boardSize*boardSize {
		board[x][y] = moveNum
		return true
	}

	nextMoves := findNextMoves(x, y, boardSize, method)

	for _, move := range nextMoves {
		nextX, nextY := move[0], move[1]
		if board[nextX][nextY] == 0 {
			board[nextX][nextY] = moveNum

			if backtrack(board, moveNum+1, nextX, nextY, boardSize, method) {
				return true
			}

			board[nextX][nextY] = 0
		}
	}

	return false
}
