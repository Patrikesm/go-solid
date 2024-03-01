package http

import (
	"fmt"
	"net/http"

	"github.com/Patrikesm/basic-solid-api/internal/models"
	"github.com/Patrikesm/basic-solid-api/internal/user"
	"github.com/gin-gonic/gin"
)

type UserHandlers struct {
	UserUC user.UseCase
}

func NewUserHandlers(userUC user.UseCase) user.Handlers {
	return &UserHandlers{UserUC: userUC}
}

func (h *UserHandlers) Create(c *gin.Context) {
	var user *models.User

	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error in field binding",
		})
		fmt.Println(err)
		return
	}

	err = h.UserUC.Create(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error creating User",
		})
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (h *UserHandlers) ListAll(c *gin.Context) {
	output, err := h.UserUC.ListAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error getting Users",
		})
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, output)
}
