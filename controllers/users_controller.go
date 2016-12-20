package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/tsuyoshiwada/gin-sandbox/models"
)

type UsersController struct {
	Controller
	db *gorm.DB
}

func NewUsersController(db *gorm.DB) *UsersController {
	return &UsersController{db: db}
}

func (ctl UsersController) GetAll(c *gin.Context) {
	users := []models.User{}
	ctl.db.Preload("Articles").Find(&users)

	ctl.SuccessResponse(c, gin.H{
		"users": users,
	})
}

func (ctl UsersController) Get(c *gin.Context) {
	id := c.Param("id")
	user := models.User{}
	ctl.db.Preload("Articles").First(&user, id)

	ctl.SuccessResponse(c, gin.H{
		"user": user,
	})
}
