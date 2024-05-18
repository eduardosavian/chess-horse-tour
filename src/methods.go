package main

import (
	"sort"
	"encoding/json"
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

func findNextMoves(x, y, boardSize int, method string) []Move {
	validMoves := []Move{}

	for _, move := range possibleMoves {
		nextX := x + move[0]
		nextY := y + move[1]

		if isMoveValid(nextX, nextY, boardSize) {
			validMoves = append(validMoves, Move{nextX, nextY, 0})
		}
	}

	if method == "Warnsdorff" {
		moveCounts := make([][]int, boardSize)
		for i := range moveCounts {
			moveCounts[i] = make([]int, boardSize)
			for j := range moveCounts[i] {
				moveCounts[i][j] = len(findNextMoves(i, j, boardSize, "default"))
			}
		}

		for i := range validMoves {
			move := &validMoves[i]
			move.Priority = moveCounts[move.X][move.Y]
		}

		sort.Slice(validMoves, func(i, j int) bool {
			return validMoves[i].Priority < validMoves[j].Priority
		})
	}

	return validMoves
}

func prioritizeMoves(x, y, boardSize int) []Move {
	validMoves := []Move{}

	for _, move := range possibleMoves {
		nextX := x + move[0]
		nextY := y + move[1]

		if isMoveValid(nextX, nextY, boardSize) {
			validMoves = append(validMoves, Move{nextX, nextY, 0})
		}
	}

	sort.Slice(validMoves, func(i, j int) bool {
		countI := len(findNextMoves(validMoves[i].X, validMoves[i].Y, boardSize, "default"))
		countJ := len(findNextMoves(validMoves[j].X, validMoves[j].Y, boardSize, "default"))
		return countI < countJ
	})

	return validMoves
}

func backtrackWithWarnsdorff(board [][]int, moveNum, x, y, boardSize int) bool {
	board[x][y] = moveNum

	if moveNum == boardSize*boardSize {
		return true
	}

	nextMoves := prioritizeMoves(x, y, boardSize)

	for _, move := range nextMoves {
		nextX, nextY := move.X, move.Y
		if board[nextX][nextY] == 0 {
			if backtrackWithWarnsdorff(board, moveNum+1, nextX, nextY, boardSize) {
				return true
			}
		}
	}

	board[x][y] = 0
	return false
}