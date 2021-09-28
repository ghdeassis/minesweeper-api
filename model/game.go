package model

import "time"

type Game struct {
	ID         int       `json:"id"`
	Rows       int       `json:"rows"`
	Columns    int       `json:"columns"`
	Mines      int       `json:"mines"`
	StartTime  time.Time `json:"startTime"`
	FinishTime time.Time `json:"finishTime"`
	Cells      [][]Cell  `json:"cells"`
}
