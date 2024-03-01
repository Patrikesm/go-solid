package user

import "github.com/gin-gonic/gin"

type Handlers interface {
	Create(c *gin.Context)
	ListAll(c *gin.Context)
}
