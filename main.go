package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"minesweeper-api/controller"
	"minesweeper-api/util"
)

func main() {
	gin.DisableConsoleColor()
	gin.SetMode(gin.ReleaseMode)
	ginrouter := gin.Default()
	ginrouter.Use(util.CORSMiddleware())
	ginrouter.POST("/newGame", controller.NewGame)
	ginrouter.POST("/putFlag", controller.PutFlag)
	ginrouter.POST("/removeFlag", controller.RemoveFlag)
	ginrouter.POST("/revealCell", controller.RevealCell)
	err := ginrouter.Run()
	if err != nil {
		log.Fatal(err)
	}
}
