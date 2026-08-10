package category

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler exposes HTTP endpoints for categories.
type Handler struct {
	service Service
}

// NewHandler creates a category HTTP handler with an injected service.
func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

// List handles GET /categories.
func (h *Handler) List(c *gin.Context) {
	// TODO: call h.service.List and return categories
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not Rupesh pandey"})
}
