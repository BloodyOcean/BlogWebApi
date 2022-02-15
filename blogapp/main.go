package main

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
)

// @title        Blog Application
// @description  This is a blog management application
// @version      1.0
// @host         localhost:8081
// @BasePath     /api/v1
// @schemes      http
func main() {
	// THROWS!
	// Dont use := if global variable already exist.
	db, err = gorm.Open(sqlite.Open("postsdb.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schemas
	db.AutoMigrate(&Post{}, &Comment{})

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "GOLANG WEB API MAIN PAGE!")
	})
	v1 := e.Group("/api/v1")
	{
		v1.GET("/post", GetAllPosts)
		v1.GET("/post/:id", GetPostById)
		v1.GET("/comment", GetAllComments)
		v1.GET("/comment/:id", GetCommentById)
		v1.POST("/post", CreatePost)
		v1.POST("/comment", CreateComment)
		v1.DELETE("/post", DeletePost)
		v1.DELETE("/post/:id", DeletePostById)
		v1.PUT("/post", UpdatePost)
		v1.PUT("/comment", UpdateComment)
		v1.DELETE("/comment/:id", DeleteCommentById)
		v1.DELETE("/comment", DeleteComment)
		v1.GET("/comment/post/:id", GetAllCommentByPostId)
	}
	e.GET("/swagger/*", echoSwagger.WrapHandler) // SWAGGER.

	e.Logger.Fatal(e.Start(":8081"))
}
