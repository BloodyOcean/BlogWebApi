package main

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm/clause"
	"net/http"
)

// CreatePost godoc
// @Router   /post [post]
// @Summary  Create new post
// @Produce  json
// @Success  200  {object}  Post
func CreatePost(c echo.Context) error {
	post := new(Post)
	if err := c.Bind(post); err != nil {
		return err
	}
	db.Create(post)
	return c.JSON(http.StatusCreated, post)
}

// GetAllPosts godoc
// @Router   /post [get]
// @Summary  Show all posts
// @Produce  json
// @Success  200  {object}  Post
func GetAllPosts(c echo.Context) error {
	posts := []Post{}
	db.Preload(clause.Associations).Find(&posts)
	return c.JSON(http.StatusOK, &posts)
}

// GetPostById godoc
// @Router       /post/{id} [get]
// @Summary      Show all posts by corresponding id
// @Produce      json
// @Param        id  path  int  true  "Id"
// @Description  get string by ID
// @Success      200  {object}  Post
func GetPostById(c echo.Context) error {
	id := c.Param("id")
	post := Post{}
	db.Preload(clause.Associations).First(&post, id)
	return c.JSON(http.StatusOK, &post)
}

// UpdatePost godoc
// @Router   /post [put]
// @Summary  update post from request body by its ID
// @Accept   json
// @Produce  json
// @Success  200  {object}  Post
func UpdatePost(c echo.Context) error {
	post := new(Post)
	if err := c.Bind(post); err != nil {
		return err
	}
	db.Save(&post)
	return c.JSON(http.StatusOK, post)
}

// DeletePostById godoc
// @Router       /post/{id} [delete]
// @Summary      delete post by its ID
// @Description  get string by ID
// @Param        id  path  int  true  "Id"
// @Produce      json
// @Success      200  {object}  Post
func DeletePostById(c echo.Context) error {
	id := c.Param("id")
	var item Post
	db.First(&item, id)
	db.Unscoped().Select(clause.Associations).Delete(&Post{}, id)
	return c.JSON(http.StatusOK, item)
}

// DeletePost godoc
// @Router   /post [delete]
// @Summary  delete post from request body by its ID
// @Accept   json
// @Produce  json
// @Success  200  {object}  Post
func DeletePost(c echo.Context) error {
	post := new(Post)
	if err := c.Bind(post); err != nil {
		return err
	}
	db.Unscoped().Select(clause.Associations).Delete(&post)
	return c.JSON(http.StatusOK, post)
}
