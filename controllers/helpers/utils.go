package helpers

import (
	"github.com/gin-gonic/gin"
	"test-gin/models"
)

func GetUserFromRequest(c *gin.Context) *models.User {
	userID := c.GetUint64("user_id")
	var currentUser *models.User

	if userID > 0 {
		currentUser = models.UserFind(userID)
	} else {
		currentUser = nil
	}

	return currentUser
}

func IsUserLoggedIn(c *gin.Context) bool {
	return c.GetUint64("user_id") > 0
}

func SetPayload(c *gin.Context, h gin.H) gin.H {
	email := c.GetString("email")
	if len(email) > 0 {
		h["email"] = email
	}
	return h
}

func SessionClear(c *gin.Context) {
	//TODO
}

func SessionSet(c *gin.Context, userId int) {
	//TODO
}
