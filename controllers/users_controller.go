package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/tsuyoshiwada/gin-sandbox/models"
)

type UsersController struct {
	db *gorm.DB
}

func NewUsersController(db *gorm.DB) *UsersController {
	return &UsersController{db}
}

func (controller UsersController) GetAll(c *gin.Context) {
	users := []models.User{}
	controller.db.Find(&users)

	c.JSON(http.StatusOK, gin.H{"status": "ok", "users": users})
}

func (controller UsersController) Get(c *gin.Context) {
	id := c.Param("id")
	user := models.User{}
	articles := []models.Article{}
	controller.db.First(&user, id)
	controller.db.Model(&user).Related(&articles)
	user.Articles = articles
	c.JSON(http.StatusOK, gin.H{"status": "ok", "user": user})
}
