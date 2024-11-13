package main

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"test-gin/controllers"
	"test-gin/controllers/helpers"
	"test-gin/middlewares"
	"test-gin/models"
)

func main() {
	fmt.Println("Hello, World!")

	r := gin.Default()
	r.Use(gin.Logger())

	r.Static("/vendor", "./static/vendor")

	r.LoadHTMLGlob("templates/**/*")

	// инициализируем БД
	models.ConnectDatabase()
	models.DBMigrate()

	//инициализируем сессии
	store := memstore.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("notes", store))

	//прикручиваем проверку аутентифицированного пользователя
	r.Use(middlewares.AuthenticateUser())

	// прописываем маршруты
	notes := r.Group("/notes")
	{
		notes.GET("/", controllers.NotesIndex)
		notes.GET("/new", controllers.NotesNew)
		notes.POST("/", controllers.NotesCreate)
		notes.GET("/:id", controllers.NotesShow)
		notes.GET("edit/:id", controllers.NotesEditPage)
		notes.POST("/:id", controllers.NotesUpdate)
		notes.DELETE("/:id", controllers.NotesDelete)
	}

	r.GET("/login", controllers.LoginPage)
	r.GET("/signup", controllers.SignupPage)

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.POST("/logout", controllers.Logout)

	r.GET("/", func(c *gin.Context) {
		c.HTML(
			http.StatusOK,
			"home/index.html",
			//TODO - add helper!
			helpers.SetPayload(c, gin.H{
				"title":     "Notes Application",
				"logged in": helpers.IsUserLoggedIn(c),
			}),
		)
	})

	addr := ":8089"
	log.Println("Server started on " + addr)
	r.Run(addr)
}
