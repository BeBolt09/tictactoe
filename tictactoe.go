package main

import (
	"fmt"
)

const (
	EMPTY = 0
	X     = 1
	O     = 2
)

type Board struct {
	board [3][3]int
	turn  int
}

func NewBoard() *Board {
	board := Board{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board.board[i][j] = EMPTY
		}
	}
	board.turn = X
	return &board
}

func (b *Board) Print() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			switch b.board[i][j] {
			case EMPTY:
				fmt.Print("-")
			case X:
				fmt.Print("X")
			case O:
				fmt.Print("O")
			}
		}
		fmt.Println()
	}
}

func (b *Board) GetInput() int {
	fmt.Print("Enter your move (1-9): ")
	var move int
	_, err := fmt.Scanf("%d", &move)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	if move < 1 || move > 9 {
		fmt.Println("Invalid move.")
		return -1
	}
	return move - 1
}

func (b *Board) MakeMove(move int) bool {
	if b.board[move/3][move%3] != EMPTY {
		return false
	}
	b.board[move/3][move%3] = b.turn
	b.turn = 3 - b.turn
	return true
}

func (b *Board) IsWinner() bool {
	for i := 0; i < 3; i++ {
		if b.board[i][0] == b.board[i][1] && b.board[i][1] == b.board[i][2] && b.board[i][0] != EMPTY {
			return true
		}
		if b.board[0][i] == b.board[1][i] && b.board[1][i] == b.board[2][i] && b.board[0][i] != EMPTY {
			return true
		}
	}
	if b.board[0][0] == b.board[1][1] && b.board[1][1] == b.board[2][2] && b.board[0][0] != EMPTY {
		return true
	}
	if b.board[0][2] == b.board[1][1] && b.board[1][1] == b.board[2][0] && b.board[0][2] != EMPTY {
		return true
	}
	return false
}

func main() {
	board := NewBoard()
	for {
		move := board.GetInput()
		if move == -1 {
			continue
		}
		if !board.MakeMove(move) {
			fmt.Println("Invalid move.")
			continue
		}
		board.Print()
		if board.IsWinner() {
			fmt.Println("Player ", board.turn, " won!")
			break
		}
	}
}
