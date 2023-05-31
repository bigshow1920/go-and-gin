package app

import (
	"context"
	"go-and-gin/config"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewApp(ctx context.Context, cfg config.Config) error {
	app, err := NewClient(ctx, cfg)
	if err != nil {
		return err
	}
	g := gin.New()
	userPath := g.Group("/players")
	{
		userPath.GET("", app.All)
		userPath.GET("/:id", app.Load)
		userPath.POST("", app.Insert)
		userPath.PUT("/:id", app.Update)
		userPath.DELETE("/:id", app.Delete)
	}

	log.Fatal(http.ListenAndServe(":8080", g))
	return nil
}
