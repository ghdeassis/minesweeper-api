package service

import (
	"math/rand"
	"minesweeper-api/model"
	"time"
)

var games []model.Game

func NewGame(rows, columns, mines int) {
	cells := make([][]model.Cell, rows)
	for i := 0; i < rows; i++ {
		cells[i] = make([]model.Cell, columns)
	}

	game := model.Game{
		Rows:      rows,
		Columns:   columns,
		Mines:     mines,
		Cells:     cells,
		StartTime: time.Now(),
	}

	generateMines(game)

	games = append(games, game)
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
