package main

import (
	"encoding/json"
	"math/rand"
	"sort"
)

type Board struct {
	Rows [][]int `json:"board"`
}

func convertBoardToJSON(board [][]int) ([]byte, error) {
	data := Board{
		Rows: board,
	}
	return json.Marshal(data)
}

func isMoveValid(x, y, boardSize int) bool {
	return x >= 0 && x < boardSize && y >= 0 && y < boardSize
}

type Move struct {
	X, Y     int
	Priority int
}

var possibleMoves = [][]int{
	{-2, -1}, {-1, -2}, {1, -2}, {2, -1},
	{2, 1}, {1, 2}, {-1, 2}, {-2, 1},
}

func findNextMoves(x, y, boardSize int, board [][]int, method string) []Move {
	validMoves := []Move{}

	for _, move := range possibleMoves {
		nextX := x + move[0]
		nextY := y + move[1]

		if isMoveValid(nextX, nextY, boardSize) && board[nextX][nextY] == 0 {
			validMoves = append(validMoves, Move{nextX, nextY, 0})
		}
	}

	switch method {
	case "warnsdorff":
		for i := range validMoves {
			move := &validMoves[i]
			move.Priority = len(findNextMoves(move.X, move.Y, boardSize, board, "default"))
		}
		sort.Slice(validMoves, func(i, j int) bool {
			return validMoves[i].Priority < validMoves[j].Priority
		})
	case "highDegree":
		for i := range validMoves {
			move := &validMoves[i]
			move.Priority = len(findNextMoves(move.X, move.Y, boardSize, board, "default"))
		}
		sort.Slice(validMoves, func(i, j int) bool {
			return validMoves[i].Priority > validMoves[j].Priority
		})
	case "shuffle":
		rand.Shuffle(len(validMoves), func(i, j int) {
			validMoves[i], validMoves[j] = validMoves[j], validMoves[i]
		})
	}

	return validMoves
}

func backtrackWithMethod(board [][]int, moveNum, x, y, boardSize int, method string) bool {
	board[x][y] = moveNum

	if moveNum == boardSize*boardSize {
		return true
	}

	nextMoves := findNextMoves(x, y, boardSize, board, method)

	for _, move := range nextMoves {
		if backtrackWithMethod(board, moveNum+1, move.X, move.Y, boardSize, method) {
			return true
		}
	}

	board[x][y] = 0
	return false
}
