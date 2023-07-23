package main

import (
	"github.com/grijalbaEmilio/triqui-game-back/src/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	boardGroup := e.Group("/board")

	boardRoutes := routes.GetBoardRouteInstance()

	boardGroup.POST("/newGame", boardRoutes.NewGame)
	boardGroup.POST("/machPlayer", boardRoutes.MachPlayer)
	boardGroup.POST("/markSquare", boardRoutes.MarkSquare)
	boardGroup.POST("/get", boardRoutes.GetBoard)

	e.Start(":8080")
}
