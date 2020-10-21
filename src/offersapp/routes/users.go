package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UsersRegister(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"user_id": "some_id",
	})
}
