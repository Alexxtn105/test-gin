package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"test-gin/controllers/helpers"
	"test-gin/models"
)

type FormData struct {
	Name    string `form:"name"`
	Content string `form:"content"`
}

func NotImplemented(c *gin.Context) {
	//c.HTML(http.StatusOK, "home/index.html", )
}

func NotesIndex(c *gin.Context) {
	currentUser := helpers.GetUserFromRequest(c)

	if currentUser == nil || currentUser.ID == 0 {
		c.HTML(
			http.StatusUnauthorized,
			"notes/index.html",
			helpers.SetPayload(c, gin.H{
				"alert": "Unauthorized Access!",
			}),
		)
		return
	}

	notes := models.NotesAll(currentUser)
	c.HTML(
		http.StatusOK,
		"notes/index.html",
		helpers.SetPayload(c, gin.H{
			"notes": notes,
		}),
	)
}

func NotesNew(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"notes/index.html",
		helpers.SetPayload(c, gin.H{}),
	)
}

func NotesCreate(c *gin.Context) {
	currentUser := helpers.GetUserFromRequest(c)
	if currentUser == nil || currentUser.ID == 0 {
		c.HTML(
			http.StatusUnauthorized,
			"notes/index.html",
			helpers.SetPayload(c, gin.H{
				"alert": "Unauthorized Access!",
			}),
		)
		return
	}
}

func NotesShow(c *gin.Context) {
	//TODO NotesShow
}

func NotesEditPage(c *gin.Context) {
	//TODO NotesEditPage
}

func NotesUpdate(c *gin.Context) {
	//TODO
	currentUser := helpers.GetUserFromRequest(c)
	if currentUser == nil || currentUser.ID == 0 {
		c.HTML(
			http.StatusUnauthorized,
			"notes/index.html",
			helpers.SetPayload(c, gin.H{
				"alert": "Unauthorized Access!",
			}),
		)
		return
	}

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	// TODO - проверить models.NotesUpdate
	models.NotesUpdate(currentUser, id)
	c.Redirect(http.StatusSeeOther, "/notes")

}

func NotesDelete(c *gin.Context) {
	currentUser := helpers.GetUserFromRequest(c)
	if currentUser == nil || currentUser.ID == 0 {
		c.HTML(
			http.StatusUnauthorized,
			"notes/index.html",
			helpers.SetPayload(c, gin.H{
				"alert": "Unauthorized Access!",
			}),
		)
		return
	}

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	models.NotesMarkDelete(currentUser, id)
	c.Redirect(http.StatusSeeOther, "/notes")

}
