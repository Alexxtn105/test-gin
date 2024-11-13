package helpers

import (
	"github.com/gin-gonic/gin"
	"test-gin/models"
)

func GetUserFromRequest(c *gin.Context) *models.User {
	//TODO
	return &models.User{
		ID:   0,
		Name: "",
	}
}

func IsUserLoggedIn(c *gin.Context) bool {
	//TODO
	return true
}

func SessionClear(c *gin.Context) {
	//TODO
}

func SessionSet(c *gin.Context, userId int) {
	//TODO
}

func SetPayload(c *gin.Context, h map[string]any) any {
	//TODO
	return nil
}
