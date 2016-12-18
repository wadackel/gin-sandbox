package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/tsuyoshiwada/gin-sandbox/models"
)

type UserController struct {
	db *gorm.DB
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{db}
}

func (controller UserController) GetAll(c *gin.Context) {
	users := []models.User{}
	controller.db.Debug().Find(&users)
	c.JSON(http.StatusOK, gin.H{"status": "ok", "users": users})
}
