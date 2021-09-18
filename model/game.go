package model

import "time"

type Game struct {
	Rows      int       `json:"rows"`
	Columns   int       `json:"columns"`
	Mines     int       `json:"mines"`
	Cells     [][]Cell  `json:"cells"`
	StartTime time.Time `json:"startTime"`
}
