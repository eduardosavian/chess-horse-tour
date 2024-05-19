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

type Move struct {
	X, Y     int
	Priority int
}

var possibleMoves = [][]int{
	{-2, -1}, {-1, -2}, {1, -2}, {2, -1},
	{2, 1}, {1, 2}, {-1, 2}, {-2, 1},
}

func isMoveValid(x, y, boardSize int) bool {
	return x >= 0 && x < boardSize && y >= 0 && y < boardSize
}

func findNextMoves(x, y, boardSize int, board [][]int, searchType string) []Move {
	validMoves := []Move{}

	for _, move := range possibleMoves {
		nextX := x + move[0]
		nextY := y + move[1]

		if isMoveValid(nextX, nextY, boardSize) && board[nextX][nextY] == 0 {
			validMoves = append(validMoves, Move{nextX, nextY, 0})
		}
	}

	switch searchType {
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

func greedySearch(board [][]int, x, y, boardSize int, searchType string) bool {
	board[x][y] = 1
	for moveNum := 2; moveNum <= boardSize*boardSize; moveNum++ {
		nextMoves := findNextMoves(x, y, boardSize, board, searchType)
		if len(nextMoves) == 0 {
			return false
		}
		move := nextMoves[0]
		x, y = move.X, move.Y
		board[x][y] = moveNum
	}
	return true
}

func backtrackSearch(board [][]int, moveNum, x, y, boardSize int, backtrackType string) bool {
	board[x][y] = moveNum

	if moveNum == boardSize*boardSize {
		return true
	}

	nextMoves := findNextMoves(x, y, boardSize, board, backtrackType)

	for _, move := range nextMoves {
		if backtrackSearch(board, moveNum+1, move.X, move.Y, boardSize, backtrackType) {
			return true
		}
	}

	board[x][y] = 0
	return false
}
