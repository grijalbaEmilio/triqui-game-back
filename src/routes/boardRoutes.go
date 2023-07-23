package routes

import (
	"fmt"
	"net/http"

	"github.com/grijalbaEmilio/triqui-game-back/src/db"
	"github.com/grijalbaEmilio/triqui-game-back/src/dto"
	"github.com/grijalbaEmilio/triqui-game-back/src/model"
	"github.com/labstack/echo/v4"
)

type NewGameRequestBody struct {
	ExPlayer string `json:"exPlayer"`
}

type MachPlayerRequestBody struct {
	BoardId      uint   `json:"boardId"`
	CirclePlayer string `json:"circlePlayer"`
}

type GetBoardRequestBody struct {
	BoardId uint `json:"boardId"`
}

type MarkSquareRequestBody struct {
	BoardId  uint `json:"boardId"`
	Position uint `json:"position"`
}

type GenericResponse struct {
	Message string `json:"message"`
}

type boardRoute struct {
	repo db.Db[model.Board]
}

func GetBoardRouteInstance() boardRoute {
	return boardRoute{
		repo: db.GetRedisRepo(),
	}
}

func (b boardRoute) NewGame(c echo.Context) error {
	var reqBody NewGameRequestBody

	if err := c.Bind(&reqBody); err != nil || len(reqBody.ExPlayer) < 1 {
		messsage := GenericResponse{Message: "Bad request body"}
		return c.JSON(http.StatusBadRequest, messsage)
	}

	board := model.NewBoard(reqBody.ExPlayer, "")

	boardJson, err := dto.BoardToJsonToJson(board)

	if err != nil {
		fmt.Println("Covert board to json error: " + err.Error())
		return c.String(http.StatusInternalServerError, "")
	}

	if err := b.repo.Add(board); err != nil {
		fmt.Println("Repo Add error: " + err.Error())
		return c.String(http.StatusInternalServerError, "")
	}

	return c.JSONBlob(http.StatusCreated, boardJson)
}

func (b boardRoute) MachPlayer(c echo.Context) error {
	var requestBody MachPlayerRequestBody

	if err := c.Bind(&requestBody); err != nil || len(requestBody.CirclePlayer) < 1 {
		messsage := GenericResponse{Message: "Bad request body"}
		return c.JSON(http.StatusBadRequest, messsage)
	}

	board, err := b.repo.GetById(requestBody.BoardId)

	if err != nil {
		fmt.Println("Get board by id error: " + err.Error())
		return c.JSON(http.StatusNotFound, GenericResponse{Message: "Board not found"})
	}

	board.SetCirclePlayer(requestBody.CirclePlayer)

	boardJson, err := dto.BoardToJsonToJson(board)

	if err != nil {
		fmt.Println("Board to JSON error: " + err.Error())
		return c.JSON(http.StatusInternalServerError, "")
	}

	if err := b.repo.UpdateById(board.GetId(), board); err != nil {
		fmt.Println("Update board by id error: " + err.Error())
		return c.JSON(http.StatusInternalServerError, "")
	}

	return c.JSONBlob(http.StatusOK, boardJson)
}

func (b boardRoute) GetBoard(c echo.Context) error {
	requestBody := new(GetBoardRequestBody)

	if err := c.Bind(requestBody); err != nil {
		fmt.Println("Bind error: " + err.Error())
		return c.JSON(http.StatusBadRequest, GenericResponse{Message: "Bad request body"})
	}

	board, err := b.repo.GetById(requestBody.BoardId)

	if err != nil {
		fmt.Println("Get board by id error: " + err.Error())
		return c.JSON(http.StatusNotFound, GenericResponse{Message: "Board not found"})
	}

	boardJson, err := dto.BoardToJsonToJson(board)

	if err != nil {
		fmt.Println("Board to JSON error: " + err.Error())
		return c.JSON(http.StatusInternalServerError, "")
	}

	return c.JSONBlob(http.StatusOK, boardJson)
}

func (b boardRoute) MarkSquare(c echo.Context) error {
	var requestBody MarkSquareRequestBody

	if err := c.Bind(&requestBody); err != nil {
		fmt.Println("Bind error: " + err.Error())
		return c.JSON(http.StatusBadRequest, GenericResponse{Message: "Bad request body"})
	}

	id := requestBody.BoardId
	position := requestBody.Position

	board, err := b.repo.GetById(id)

	if err != nil {
		fmt.Println("Get board by id error: " + err.Error())
		return c.JSON(http.StatusInternalServerError, "")
	}

	if err := board.MarkSquare(position); err != nil {
		return c.JSON(http.StatusBadRequest, GenericResponse{Message: err.Error()})
	}

	if err := b.repo.UpdateById(id, board); err != nil {
		fmt.Println("Update by id error: " + err.Error())
		return c.JSON(http.StatusInternalServerError, "")
	}

	boardJson, err := dto.BoardToJsonToJson(board)

	if err != nil {
		fmt.Println("Board to JSON error: " + err.Error())
		return c.JSON(http.StatusInternalServerError, "")
	}

	return c.JSONBlob(http.StatusOK, boardJson)
}
