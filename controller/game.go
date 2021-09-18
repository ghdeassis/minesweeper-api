package controller

import (
	"github.com/gin-gonic/gin"
	"minesweeper-api/service"
	"net/http"
)

func NewGame(c *gin.Context) {
	request := NewGameRequest{}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	service.NewGame(request.Rows, request.Columns, request.Mines)

	c.JSON(http.StatusOK, "")
}

type NewGameRequest struct {
	Rows    int `json:"rows" binding:"required"`
	Columns int `json:"columns" binding:"required"`
	Mines   int `json:"mines" binding:"required"`
}
