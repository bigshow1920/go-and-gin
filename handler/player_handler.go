package handler

import (
	"go-and-gin/models"
	"go-and-gin/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PlayerHandler struct {
	service service.PlayerService
}

func NewPlayerHandler(service service.PlayerService) *PlayerHandler {
	return &PlayerHandler{service: service}
}

func (h *PlayerHandler) All(g *gin.Context) {
	res, err := h.service.GetAllPlayers()
	if err != nil {
		g.String(http.StatusInternalServerError, err.Error())
		return
	}
	g.JSON(http.StatusOK, res)
}

func (h *PlayerHandler) Load(g *gin.Context) {
	id := g.Param("id")
	if len(id) == 0 {
		g.String(http.StatusBadRequest, "Id cant not empty")
		return
	}
	id64, err := Convert(id)
	if err != nil {
		g.String(http.StatusBadRequest, "Id cant not convert")
		return
	}
	res, err := h.service.GetPlayer(id64)
	if err != nil {
		g.String(http.StatusBadRequest, err.Error())
		return
	}
	g.JSON(http.StatusOK, res)
}

func (h *PlayerHandler) Insert(g *gin.Context) {
	var Player models.Player

	er1 := g.ShouldBindJSON(&Player)
	if er1 != nil {
		g.String(http.StatusBadRequest, er1.Error())
		return
	}

	res, er2 := h.service.InsertPlayer(Player)
	if er2 != nil {
		g.String(http.StatusBadRequest, er2.Error())
		return
	}
	g.JSON(http.StatusOK, res)
}

func (h *PlayerHandler) Update(g *gin.Context) {
	var Player models.Player
	er1 := g.ShouldBindJSON(&Player)

	if er1 != nil {
		g.String(http.StatusBadRequest, er1.Error())
		return
	}
	id := g.Param("id")
	if len(id) == 0 {
		g.String(http.StatusBadRequest, "Id cant not empty")
		return
	}
	id64, err := Convert(id)
	if err != nil {
		g.String(http.StatusBadRequest, "Id cant not convert")
		return
	}
	if Player.ID == 0 {
		Player.ID = id64
	} else if id64 != Player.ID {
		g.String(http.StatusBadRequest, "Id not match")
		return
	}

	res, er2 := h.service.UpdatePlayer(Player.ID, Player)
	if er2 != nil {
		g.String(http.StatusInternalServerError, er2.Error())
		return
	}
	g.JSON(http.StatusOK, res)
}

func Convert(id string) (int64, error) {
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return -1, err
	}
	return i, nil
}
func (h *PlayerHandler) Delete(g *gin.Context) {
	id := g.Param("id")
	if len(id) == 0 {
		g.String(http.StatusBadRequest, "Id cannot be empty")
		return
	}
	id64, err := Convert(id)
	if err != nil {
		g.String(http.StatusBadRequest, "Id cannot convert")
		return
	}
	res, err := h.service.DeletePlayer(id64)
	if err != nil {
		g.String(http.StatusInternalServerError, err.Error())
		return
	}
	g.JSON(http.StatusOK, res)
}
