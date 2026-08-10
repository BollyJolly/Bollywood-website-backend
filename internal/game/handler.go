package game

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler exposes HTTP endpoints for gameplay.
type Handler struct {
	service Service
}

// NewHandler creates a game HTTP handler with an injected service.
func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

// GetCard handles GET /game/:roomId/card.
func (h *Handler) GetCard(c *gin.Context) {
	// TODO: read roomId and playerID, call h.service.GetCard
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// Mark handles POST /game/:roomId/mark.
func (h *Handler) Mark(c *gin.Context) {
	// TODO: bind song ID, call h.service.Mark
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// Bingo handles POST /game/:roomId/bingo.
func (h *Handler) Bingo(c *gin.Context) {
	// TODO: call h.service.ClaimBingo and return result
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}
