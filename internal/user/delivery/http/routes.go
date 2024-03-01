package http

import (
	"github.com/Patrikesm/basic-solid-api/internal/user"
	"github.com/gin-gonic/gin"
)

func MapUserRoutes(router *gin.Engine, h user.Handlers) {
	userEndpoints := router.Group("/user")

	{
		userEndpoints.POST("", h.Create)
		userEndpoints.GET("", h.ListAll)
	}
}
