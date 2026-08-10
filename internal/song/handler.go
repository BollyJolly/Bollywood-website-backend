package song

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler exposes HTTP endpoints for songs.
type Handler struct {
	service Service
}

// NewHandler creates a song HTTP handler with an injected service.
func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

// ListByCategory handles GET /categories/:id/songs.
func (h *Handler) ListByCategory(c *gin.Context) {
	// TODO: read :id, call h.service.ListByCategory, return songs
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}
