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
	err := ginrouter.Run()
	if err != nil {
		log.Fatal(err)
	}
}
