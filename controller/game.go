package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewGame(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}
