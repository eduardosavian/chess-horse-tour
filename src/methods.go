package main

import (
	"sort"
)

const(
	possibleMoves := [][]int{
		{-2, -1}, {-1, -2}, {1, -2}, {2, -1},
		{2, 1}, {1, 2}, {-1, 2}, {-2, 1},
	}
)


func isMoveValid(x, y, boardSize int) bool {
	return x >= 0 && x < boardSize && y >= 0 && y < boardSize
}

type Move struct {
    X, Y     int
    Priority int
}

func findNextMovesWarnsdorff(x, y, boardSize int) [][]int {
    moves := make([]Move, 0)

    for _, move := range possibleMoves {
        nextX := x + move[0]
        nextY := y + move[1]

        if isMoveValid(nextX, nextY, boardSize) {
            validMovesFromNext := len(findNextMoves(nextX, nextY, boardSize))
            moves = append(moves, Move{nextX, nextY, validMovesFromNext})
        }
    }

    sort.Slice(moves, func(i, j int) bool {
        return moves[i].Priority < moves[j].Priority
    })

    validMoves := [][]int{}
    for _, move := range moves {
        validMoves = append(validMoves, []int{move.X, move.Y})
    }

    return validMoves
}

func backtrackWarnsdorff(board [][]int, moveNum, x, y, boardSize int, ) bool {
	nextMoves := findNextMovesWarnsdorff(x, y, boardSize)

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

