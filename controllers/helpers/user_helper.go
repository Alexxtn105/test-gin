package helpers

import (
	"github.com/gin-gonic/gin"
	"test-gin/models"
)

func GetUserFromRequest(c *gin.Context) *models.User {

	return &models.User{
		ID:   0,
		Name: "",
	}
}

func IsUserLoggedIn(c *gin.Context) bool {

	return true
}
