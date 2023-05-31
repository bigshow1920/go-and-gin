package app

import (
	"context"
	"database/sql"
	"go-and-gin/config"
	"go-and-gin/handler"
	"go-and-gin/service"
)

func NewClient(ctx context.Context, config config.Config) (*handler.PlayerHandler, error) {
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		return nil, err
	}
	playerService := service.NewPLayerService(conn)
	playerHandler := handler.NewPlayerHandler(playerService)
	return playerHandler, nil
}
