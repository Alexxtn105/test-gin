package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"test-gin/controllers"
)

func main() {
	fmt.Println("Hello, World!")

	r := gin.Default()
	r.Use(gin.Logger())

	r.Static("/vendor", "./static/vendor")

	//r.LoadHTMLGlob("templates/**/*")

	//models.ConnectDatabase()
	//models.DBMigrate()

	notes := r.Group("/notes")
	{
		notes.GET("/", controllers.NotImplemented)
		//notes.GET("/new", controllers.NotImplemented)
		//notes.POST("/", controllers.NotImplemented)
		//notes.GET("/:id", controllers.NotImplemented)
		//notes.GET("edit/:id", controllers.NotImplemented)
		//notes.POST("/:id", controllers.NotImplemented)
		//notes.DELETE("/:id", controllers.NotImplemented)
	}

	r.GET("/login", controllers.NotImplemented)
	r.GET("/signup", controllers.NotImplemented)

	r.POST("/signup", controllers.NotImplemented)
	r.POST("/login", controllers.NotImplemented)
	r.POST("/logout", controllers.NotImplemented)

	addr := ":8089"
	log.Println("Server started on " + addr)
	r.Run(addr)
}
