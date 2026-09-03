package feedback

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

type CreateRequest struct {
	Rating  int    `json:"rating"`
	Message string `json:"message"`
}

func (h *Handler) Create(
	c *gin.Context,
) {

	userID := c.GetString(
		"userId",
	)

	var request CreateRequest

	err := c.ShouldBindJSON(
		&request,
	)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "invalid request body",
			},
		)

		return
	}

	if request.Rating < 1 || request.Rating > 5 {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "rating must be between 1 and 5",
			},
		)

		return
	}

	err = h.service.Create(
		userID,
		request.Rating,
		request.Message,
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
		gin.H{
			"message": "feedback submitted successfully",
		},
	)
}