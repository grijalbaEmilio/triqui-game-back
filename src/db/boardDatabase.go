package db

import (
	"context"
	"fmt"
	"time"

	"github.com/grijalbaEmilio/triqui-game-back/src/dto"
	"github.com/grijalbaEmilio/triqui-game-back/src/model"
	"github.com/redis/go-redis/v9"
)

type BdRedisBoard struct {
	client  *redis.Client
	context context.Context
}

func GetRedisRepo() BdRedisBoard {
	return BdRedisBoard{}
}

func (b BdRedisBoard) getClient() *redis.Client {
	if b.client == nil {
		b.client = redis.NewClient(&redis.Options{
			Addr:     "redisdb:6379",
			Password: "", // si tienes autenticación
			DB:       1,  // número de la base de datos Redis
		})
	}

	return b.client
}

func (b BdRedisBoard) getContext() context.Context {
	if b.context == nil {
		b.context = context.Background()
	}
	return b.context
}

func (b BdRedisBoard) Add(board model.Board) error {
	client := b.getClient()
	ctx := b.getContext()
	id := fmt.Sprint(board.GetId())

	boardJson, err := dto.BoardToJsonToJson(board)
	if err != nil {
		return err
	}

	return client.Set(ctx, id, string(boardJson), time.Minute*10).Err()
}

func (b BdRedisBoard) GetById(id uint) (model.Board, error) {
	client := b.getClient()
	ctx := b.getContext()

	boardStr, err := client.Get(ctx, fmt.Sprint(id)).Result()
	if err != nil {
		return model.Board{}, err
	}

	board, err := dto.JsonToBoard([]byte(boardStr))

	if err != nil {
		return model.Board{}, err
	}

	return board, nil

}

func (b BdRedisBoard) UpdateById(id uint, board model.Board) error {
	client := b.getClient()
	ctx := b.getContext()
	currentId := fmt.Sprint(id)

	boardJson, err := dto.BoardToJsonToJson(board)
	if err != nil {
		return err
	}

	return client.Set(ctx, currentId, string(boardJson), time.Minute).Err()
}
