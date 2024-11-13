package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"test-gin/controllers/helpers"
	"test-gin/models"
)

func LoginPage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"home/login.html",
		gin.H{},
	)
}

func SignupPage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"home/signup.html",
		gin.H{},
	)
}

func Signup(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	confirmPassword := c.PostForm("confirm_password")

	// проверяем, существует ли email
	available := models.UserCheckAvailability(email)
	fmt.Println(available)
	if !available {
		c.HTML(
			http.StatusNotAcceptable,
			"home/signup.html",
			gin.H{"alert": "Email already exist!"},
		)
		return
	}
	// проверяем, совпадают ли пароль и подтверждение
	if password != confirmPassword {
		c.HTML(
			http.StatusNotAcceptable,
			"home/signup.html",
			gin.H{"alert": "Password mismatch!"},
		)
		return
	}

	// создаем пользователя
	user := models.UserCreate(email, password)
	if user.ID == 0 {
		c.HTML(
			http.StatusNotAcceptable,
			"home/signup.html",
			gin.H{"alert": "Unable to create user!"},
		)
	} else {
		// создание пользователя прошло успешно, настраиваем сессию
		helpers.SessionSet(c, user.ID)
		//редирект в корневую страницу
		c.Redirect(http.StatusMovedPermanently, "/")
	}

}

func Login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	user := models.UserCheck(email, password)
	if user != nil {
		// Set session
		helpers.SessionSet(c, user.ID)
		c.Redirect(http.StatusMovedPermanently, "/")
	} else {
		c.HTML(
			http.StatusNotAcceptable,
			"home/login.html",
			gin.H{"alert": "Email and/or login mismatch!"},
		)
	}

}

func Logout(c *gin.Context) {
	helpers.SessionClear(c)
	c.HTML(
		http.StatusOK,
		"home/login.html",
		gin.H{"alert": "Logged out"},
	)
}
