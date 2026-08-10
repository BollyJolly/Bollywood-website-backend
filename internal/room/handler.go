package room

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(
	service *Service,
) *Handler {

	return &Handler{
		service: service,
	}
}

func (h *Handler) Create(
	c *gin.Context,
) {

	var request CreateRoomRequest

	if err := c.ShouldBindJSON(
		&request,
	); err != nil {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	userID := c.GetString(
		"userId",
	)

	room, err := h.service.Create(
		userID,
		request.PlaylistID,
	)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	c.JSON(
		http.StatusCreated,
		room,
	)
}

func (h *Handler) Join(
	c *gin.Context,
) {

	roomCode := c.Param(
		"code",
	)

	userID := c.GetString(
		"userId",
	)

	err := h.service.Join(
		roomCode,
		userID,
	)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "joined room successfully",
		},
	)
}

func (h *Handler) GetAll(
	c *gin.Context,
) {

	rooms, err := h.service.GetAll()

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	c.JSON(
		http.StatusOK,
		rooms,
	)
}

func (h *Handler) GetByCode(
	c *gin.Context,
) {

	roomCode := c.Param(
		"code",
	)

	room, err := h.service.GetByCode(
		roomCode,
	)

	if err != nil {
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"error": "room not found",
			},
		)

		return
	}

	c.JSON(
		http.StatusOK,
		room,
	)
}

func (h *Handler) Start(
	c *gin.Context,
) {

	roomCode := c.Param(
		"code",
	)

	userID := c.GetString(
		"userId",
	)

	err := h.service.Start(
		roomCode,
		userID,
	)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "game started",
		},
	)
}

func (h *Handler) Leave(
	c *gin.Context,
) {

	roomCode := c.Param(
		"code",
	)

	userID := c.GetString(
		"userId",
	)

	err := h.service.Leave(
		roomCode,
		userID,
	)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "left room successfully",
		},
	)
}

func (h *Handler) Next(
	c *gin.Context,
) {

	roomCode := c.Param(
		"code",
	)

	userID := c.GetString(
		"userId",
	)

	number, url, err := h.service.Next(
		roomCode,
		userID,
	)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"number": number,
			"url":    url,
		},
	)
}