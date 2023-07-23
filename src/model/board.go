package model

import (
	"fmt"
	"math/rand"
)

var WINNER_COMBOS = [][3]int{
	{0, 1, 2},
	{3, 4, 5},
	{6, 7, 8},
	{0, 3, 6},
	{1, 4, 7},
	{2, 5, 8},
	{0, 4, 8},
	{2, 4, 6},
}

type Board struct {
	id           uint
	exPlayer     string
	circlePlayer string
	turn         Turn
	board        [9]string
}

func NewBoard(exPlayer string, circlePlayer string) Board {

	b := Board{}
	b.id = uint(rand.Uint32())
	b.exPlayer = exPlayer
	b.circlePlayer = circlePlayer
	b.turn = X
	b.board = [9]string{}

	return b
}

func CreateBoard(id uint, exPlayer string, circlePlayer string, turn string, board [9]string) Board {
	b := Board{}
	var newTurn Turn
	if turn == "X" {
		newTurn = X
	} else {
		newTurn = O
	}

	b.id = id
	b.exPlayer = exPlayer
	b.circlePlayer = circlePlayer
	b.turn = newTurn
	b.board = board

	return b
}

func (b *Board) SetCirclePlayer(circlePlayer string) {
	b.circlePlayer = circlePlayer
}

func (b Board) GetId() uint {
	return b.id
}

func (b Board) GetExPlayer() string {
	return b.exPlayer
}

func (b Board) GetCirclePlayer() string {
	return b.circlePlayer
}

func (b Board) GetTurn() string {
	return b.turn.String()
}

func (b Board) GetBoard() [9]string {
	return b.board
}

func (b *Board) ChangeTurn() {
	if b.turn == X {
		b.turn = O
		return
	}

	b.turn = X
}

func (b *Board) MarkSquare(position uint) error {
	if position > 8 {
		return fmt.Errorf("Invalid position")
	}

	isEmpty, err := b.SquareIsEmpty(position)
	if err != nil {
		return err
	}

	if !isEmpty || b.HasWinner() || b.HasTie() {
		return nil
	}

	b.board[position] = b.turn.String()
	b.ChangeTurn()
	return nil
}

func (b Board) SquareIsEmpty(position uint) (bool, error) {
	if position > 8 {
		return false, fmt.Errorf("Invalid position")
	}
	return b.board[position] == "", nil
}

func (b Board) HasWinner() bool {
	for _, combo := range WINNER_COMBOS {
		if b.board[combo[0]] != "" && b.board[combo[0]] == b.board[combo[1]] && b.board[combo[1]] == b.board[combo[2]] {
			return true
		}
	}
	return false
}

func (b Board) IsFull() bool {
	for _, square := range b.board {
		if square == "" {
			return false
		}
	}
	return true
}

func (b Board) HasTie() bool {
	return !b.HasWinner() && b.IsFull()
}
