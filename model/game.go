package model

import "time"

type Game struct {
	ID         int       `json:"id"`
	Rows       int       `json:"rows"`
	Columns    int       `json:"columns"`
	Mines      int       `json:"mines"`
	Flags      int       `json:"flags"`
	StartTime  time.Time `json:"startTime"`
	FinishTime time.Time `json:"finishTime"`
	IsWinner   bool      `json:"isWinner"`
	IsLoser    bool      `json:"isLoser"`
	Cells      [][]Cell  `json:"cells"`
}
