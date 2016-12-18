package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/tsuyoshiwada/gin-sandbox/models"
)

type ArticlesController struct {
	db *gorm.DB
}

func NewArticlesController(db *gorm.DB) *ArticlesController {
	return &ArticlesController{db}
}

func (controller ArticlesController) GetAll(c *gin.Context) {
	articles := []models.Article{}
	controller.db.Find(&articles)

	c.JSON(http.StatusOK, gin.H{"status": "ok", "articles": articles})
}
