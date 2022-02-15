package main

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm/clause"
	"net/http"
)

// CreateComment godoc
// @Summary  Create new comment
// @Router   /comment [post]
// @Success  201  {object}  Comment
// @Produce  json
func CreateComment(c echo.Context) error {
	comment := new(Comment)
	if err := c.Bind(comment); err != nil {
		return err
	}
	db.Create(&comment)
	return c.JSON(http.StatusCreated, comment)
}

// GetAllComments godoc
// @Summary  Show all comments
// @Produce  json
// @Success  200  {object}  Comment
// @Router   /comment [get]
func GetAllComments(c echo.Context) error {
	comments := []Comment{}
	db.Find(&comments)
	return c.JSON(http.StatusOK, comments)
}

// GetAllCommentByPostId godoc
// @Summary      Show all comments by corresponding post id
// @Produce      json
// @Description  get string by ID
// @Param        id   path      int  true  "Id"
// @Success      200  {object}  Comment
// @Router       /comment/post/{id} [get]
func GetAllCommentByPostId(c echo.Context) error {
	id := c.Param("id")
	comments := []Comment{}
	db.Where("post_id = ?", id).Find(&comments)
	return c.JSON(http.StatusOK, comments)
}

// GetCommentById godoc
// @Summary  Show comment by corresponding id
// @Router   /comment/{id} [get]
// @Param    id  path  int  true  "Id"
// @Produce  json
// @Success  200  {object}  Comment
func GetCommentById(c echo.Context) error {
	id := c.Param("id")
	comment := Comment{}
	db.First(&comment, id)
	return c.JSON(http.StatusOK, comment)
}

// UpdateComment godoc
// @Summary  update comment from request body by its ID
// @Router   /comment [put]
// @Accept   json
// @Produce  json
// @Success  200  {object}  Comment
func UpdateComment(c echo.Context) error {
	comment := new(Comment)
	if err := c.Bind(comment); err != nil {
		return err
	}
	db.Save(&comment)
	return c.JSON(http.StatusOK, comment)
}

// DeleteCommentById godoc
// @Summary      delete comment from db by its ID
// @Description  get string by ID
// @Param        id  path  int  true  "Id"
// @Router       /comment/{id} [delete]
// @Produce      json
// @Success      200  {object}  Comment
func DeleteCommentById(c echo.Context) error {
	id := c.Param("id")
	var comment Comment
	db.First(&comment, id)
	db.Unscoped().Select(clause.Associations).Delete(&Comment{}, id)
	return c.JSON(http.StatusOK, comment)
}

// DeleteComment godoc
// @Summary  delete comment from request body by its ID
// @Router   /comment [delete]
// @Accept   json
// @Produce  json
// @Success  200  {object}  Comment
func DeleteComment(c echo.Context) error {
	comment := new(Comment)
	if err := c.Bind(comment); err != nil {
		return err
	}
	db.Unscoped().Select(clause.Associations).Delete(&comment)
	return c.JSON(http.StatusOK, comment)
}
