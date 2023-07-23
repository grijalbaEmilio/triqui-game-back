package dto

import (
	"encoding/json"

	"github.com/grijalbaEmilio/triqui-game-back/src/model"
)

type boardDtoJson struct {
	Id           uint      `json:"id"`
	ExPlayer     string    `json:"exPlayer"`
	CirclePlayer string    `json:"circlePlayer"`
	Turn         string    `json:"turn"`
	Board        [9]string `json:"board"`
}

func BoardToJsonToJson(b model.Board) ([]byte, error) {
	dtoJson := boardDtoJson{
		Id:           b.GetId(),
		ExPlayer:     b.GetExPlayer(),
		CirclePlayer: b.GetCirclePlayer(),
		Turn:         b.GetTurn(),
		Board:        b.GetBoard(),
	}

	return json.Marshal(dtoJson)
}

func JsonToBoard(jsonBoard []byte) (model.Board, error) {
	var dtoJson boardDtoJson
	err := json.Unmarshal(jsonBoard, &dtoJson)
	if err != nil {
		return model.Board{}, err
	}

	board := model.CreateBoard(
		dtoJson.Id,
		dtoJson.ExPlayer,
		dtoJson.CirclePlayer,
		dtoJson.Turn,
		dtoJson.Board,
	)

	return board, nil
}
