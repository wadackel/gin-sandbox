package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/tsuyoshiwada/gin-sandbox/models"
)

type ArticlesController struct {
	Controller
	db *gorm.DB
}

func NewArticlesController(db *gorm.DB) *ArticlesController {
	return &ArticlesController{db: db}
}

func (ctl ArticlesController) GetAll(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	articles := []models.Article{}
	ctl.db.Model(&user).Related(&articles)

	ctl.SuccessResponse(c, gin.H{
		"articles": articles,
	})
}

func (ctl ArticlesController) Get(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	id := c.Param("id")
	article := models.Article{}
	ctl.db.Model(&user).Related(&[]models.Article{}).First(&article, id)

	if article.ID < 1 {
		ctl.ErrorResponse(c, http.StatusBadRequest, "パラメータが不正です")
		return
	}

	ctl.SuccessResponse(c, gin.H{
		"article": article,
	})
}

type CreateArticleJSON struct {
	Title string `json:"title" binding:"required"`
	Body  string `json:"body" binding:"required"`
}

func (ctl ArticlesController) Create(c *gin.Context) {
	var json CreateArticleJSON
	if c.BindJSON(&json) != nil {
		ctl.ErrorResponse(c, http.StatusBadRequest, "パラメータが不正です")
		return
	}

	user := c.MustGet("user").(models.User)
	article := models.Article{
		UserID: user.ID,
		Title:  json.Title,
		Body:   json.Body,
	}
	ctl.db.Create(&article)

	ctl.SuccessResponse(c, gin.H{
		"article": article,
	})
}

type updateArticleJSON struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func (ctl ArticlesController) Update(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	id := c.Param("id")
	article := models.Article{}
	ctl.db.Model(&user).Related(&[]models.Article{}).First(&article, id)
	if article.ID < 1 {
		ctl.ErrorResponse(c, http.StatusBadRequest, "パラメータが不正です")
		return
	}

	var json updateArticleJSON
	if c.BindJSON(&json) != nil {
		ctl.ErrorResponse(c, http.StatusBadRequest, "パラメータが不正です")
		return
	}

	ctl.db.Model(&article).Updates(&models.Article{
		Title: json.Title,
		Body:  json.Body,
	})

	ctl.SuccessResponse(c, gin.H{
		"article": article,
	})
}

func (ctl ArticlesController) Delete(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	id := c.Param("id")
	article := models.Article{}
	ctl.db.Model(&user).Related(&[]models.Article{}).First(&article, id)
	if article.ID < 1 {
		ctl.ErrorResponse(c, http.StatusBadRequest, "パラメータが不正です")
		return
	}

	ctl.db.Delete(&article)
	ctl.SuccessResponse(c, gin.H{
		"article": article,
	})
}
