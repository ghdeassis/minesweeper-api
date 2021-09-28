package service

import (
	"errors"
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
		Flags:     0,
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

func PutFlag(gameID, row, column int) (model.Game, error) {
	game := games[gameID]
	if game.Flags == game.Mines {
		return game, errors.New("there are not remaining flags")
	}
	game.Cells[row][column].IsFlagged = true
	game.Flags++
	return game, nil
}

func RemoveFlag(gameID, row, column int) (model.Game, error) {
	game := games[gameID]
	if game.Cells[row][column].IsFlagged {
		return game, errors.New("this cell is not flagged")
	}
	game.Cells[row][column].IsFlagged = false
	game.Flags++
	return game, nil
}

func RevealCell(gameID, row, column int) model.Game {
	game := games[gameID]
	if game.Cells[row][column].HasMine {
		game.FinishTime = time.Now()
		game.IsLoser = true
		return game
	}

	game.Cells[row][column].IsRevealed = true

	minesNear := checkNeighborhood(game, row, column)
	if minesNear > 0 {
		game.Cells[row][column].MinesNear = minesNear
	} else {
		if column > 0 {
			RevealCell(gameID, row, column-1)
		}
		if column < game.Columns-1 {
			RevealCell(gameID, row, column+1)
		}
		if row > 0 {
			RevealCell(gameID, row-1, column)
		}
		if row < game.Rows-1 {
			RevealCell(gameID, row+1, column)
		}
		if row > 0 && column > 0 {
			RevealCell(gameID, row-1, column-1)
		}
		if row > 0 && column < game.Columns-1 {
			RevealCell(gameID, row-1, column+1)
		}
		if row < game.Rows-1 && column > 0 {
			RevealCell(gameID, row+1, column-1)
		}
		if row < game.Rows-1 && column < game.Columns-1 {
			RevealCell(gameID, row+1, column+1)
		}
	}

	if checkWinner(game) {
		game.FinishTime = time.Now()
		game.IsWinner = true
	}

	return game
}

func checkNeighborhood(game model.Game, row, column int) int {
	nearMinesCount := 0
	if checkLeft(game, row, column) {
		nearMinesCount++
	}
	if checkRight(game, row, column) {
		nearMinesCount++
	}
	if checkTop(game, row, column) {
		nearMinesCount++
	}
	if checkBottom(game, row, column) {
		nearMinesCount++
	}
	if checkTopLeft(game, row, column) {
		nearMinesCount++
	}
	if checkTopRight(game, row, column) {
		nearMinesCount++
	}
	if checkBottomLeft(game, row, column) {
		nearMinesCount++
	}
	if checkBottomRight(game, row, column) {
		nearMinesCount++
	}
	return nearMinesCount
}

func checkLeft(game model.Game, row, column int) bool {
	if column > 0 {
		if game.Cells[row][column-1].HasMine {
			return true
		}
	}
	return false
}

func checkRight(game model.Game, row, column int) bool {
	if column < game.Columns-1 {
		if game.Cells[row][column+1].HasMine {
			return true
		}
	}
	return false
}

func checkTop(game model.Game, row, column int) bool {
	if row > 0 {
		if game.Cells[row-1][column].HasMine {
			return true
		}
	}
	return false
}

func checkBottom(game model.Game, row, column int) bool {
	if row < game.Rows-1 {
		if game.Cells[row+1][column].HasMine {
			return true
		}
	}
	return false
}

func checkTopLeft(game model.Game, row, column int) bool {
	if row > 0 && column > 0 {
		if game.Cells[row-1][column-1].HasMine {
			return true
		}
	}
	return false
}

func checkTopRight(game model.Game, row, column int) bool {
	if row > 0 && column < game.Columns-1 {
		if game.Cells[row-1][column+1].HasMine {
			return true
		}
	}
	return false
}

func checkBottomLeft(game model.Game, row, column int) bool {
	if row < game.Rows-1 && column > 0 {
		if game.Cells[row+1][column-1].HasMine {
			return true
		}
	}
	return false
}

func checkBottomRight(game model.Game, row, column int) bool {
	if row < game.Rows-1 && column < game.Columns-1 {
		if game.Cells[row+1][column+1].HasMine {
			return true
		}
	}
	return false
}

func checkWinner(game model.Game) bool {
	if game.Mines != game.Flags {
		return false
	}

	for i := 0; i < game.Rows; i++ {
		for j := 0; j < game.Columns; j++ {
			if game.Cells[i][j].HasMine && !game.Cells[i][j].IsFlagged {
				return false
			}
		}
	}

	return true
}
