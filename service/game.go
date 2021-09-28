package service

import (
	"math/rand"
	"minesweeper-api/model"
	"time"
)

var nextId = 0
var games []model.Game

func NewGame(rows, columns, mines int) model.Game {
	cells := make([][]model.Cell, rows)
	for i := 0; i < rows; i++ {
		cells[i] = make([]model.Cell, columns)
	}

	game := model.Game{
		ID:        nextId,
		Rows:      rows,
		Columns:   columns,
		Mines:     mines,
		Cells:     cells,
		StartTime: time.Now(),
	}
	nextId++

	generateMines(game)

	games = append(games, game)

	return game
}

func generateMines(game model.Game) {
	minesPosition := rand.Perm(game.Rows * game.Columns)[0:game.Mines]
	row := 0
	column := 0
	minesCount := 0
	for i := 0; i < game.Rows*game.Columns; i++ {
		for _, p := range minesPosition {
			if p == i {
				game.Cells[row][column].HasMine = true
				minesCount++
				if minesCount == game.Mines {
					return
				}
			}
		}

		column++
		if column == game.Columns {
			column = 0
			row++
		}
	}
}

func PutFlag(gameID, row, column int) model.Game {
	game := games[gameID]
	game.Cells[row][column].IsFlagged = true
	return game
}

func RemoveFlag(gameID, row, column int) model.Game {
	game := games[gameID]
	game.Cells[row][column].IsFlagged = false
	return game
}
