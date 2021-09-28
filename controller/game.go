package controller

import (
	"github.com/gin-gonic/gin"
	"minesweeper-api/request"
	"minesweeper-api/service"
	"net/http"
)

func NewGame(c *gin.Context) {
	req := request.NewGameRequest{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, service.NewGame(*req.Rows, *req.Columns, *req.Mines))
}

func PutFlag(c *gin.Context) {
	req := request.PlayRequest{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	game, err := service.PutFlag(*req.GameID, *req.Row, *req.Column)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, game)
}

func RemoveFlag(c *gin.Context) {
	req := request.PlayRequest{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	game, err := service.RemoveFlag(*req.GameID, *req.Row, *req.Column)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, game)
}

func RevealCell(c *gin.Context) {
	req := request.PlayRequest{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, service.RevealCell(*req.GameID, *req.Row, *req.Column))
}
