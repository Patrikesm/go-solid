package server

import (
	"github.com/Patrikesm/basic-solid-api/internal/user/delivery/http"
	"github.com/Patrikesm/basic-solid-api/internal/user/repository"
	"github.com/Patrikesm/basic-solid-api/internal/user/usecase"
	"github.com/gin-gonic/gin"
)

func (s *Server) MapHandlers(r *gin.Engine) error {
	userRepo := repository.NewUserRepository(s.db)

	userUC := usecase.NewUserUseCase(*userRepo)

	userHandlers := http.NewUserHandlers(userUC)

	http.MapUserRoutes(r, userHandlers)

	return nil
}
