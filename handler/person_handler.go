package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Ping(c *gin.Context) {
	c.JSON(http.StatusCreated,Response{
		Message: "Assalomu Alaykum giybatchilar tolmanglar",
	})
}