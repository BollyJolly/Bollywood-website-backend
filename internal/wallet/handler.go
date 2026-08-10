package wallet

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler exposes HTTP endpoints for wallet operations.
type Handler struct {
	service Service
}

// NewHandler creates a wallet HTTP handler with an injected service.
func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

// GetBalance handles GET /wallet.
func (h *Handler) GetBalance(c *gin.Context) {
	// TODO: read userID, call h.service.GetBalance
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// GetHistory handles GET /wallet/history.
func (h *Handler) GetHistory(c *gin.Context) {
	// TODO: read userID, call h.service.GetHistory
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}
