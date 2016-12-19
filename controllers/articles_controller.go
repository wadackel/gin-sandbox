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
	user := c.MustGet("user").(models.User)
	articles := []models.Article{}
	controller.db.Where(&models.Article{UserID: user.ID}).Find(&articles)

	c.JSON(http.StatusOK, gin.H{"status": "ok", "articles": articles})
}

func (controller ArticlesController) Get(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	id := c.Param("id")
	article := models.Article{}
	controller.db.Where(&models.Article{UserID: user.ID}).First(&article, id)
	if article.ID < 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "パラメータが不正です",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"article": article,
	})
}

type CreateJSON struct {
	Title string `json:"title" binding:"required"`
	Body  string `json:"body" binding:"required"`
}

func (controller ArticlesController) Create(c *gin.Context) {
	var json CreateJSON
	if c.BindJSON(&json) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "パラメータが無効です",
		})
		return
	}

	user := c.MustGet("user").(models.User)
	article := models.Article{
		UserID: user.ID,
		Title:  json.Title,
		Body:   json.Body,
	}
	controller.db.Create(&article)

	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"article": article,
	})
}

func (controller ArticlesController) Delete(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	id := c.Param("id")
	article := models.Article{}
	controller.db.Where(&models.Article{UserID: user.ID}).First(&article, id)
	if article.ID < 1 {
		c.JSON(http.StatusOK, gin.H{
			"status":  "error",
			"message": "パラメータが不正です",
		})
		return
	}

	controller.db.Delete(&article)
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"article": article,
	})
}
