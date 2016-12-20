package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/tsuyoshiwada/gin-sandbox/models"
)

type TagsController struct {
	Controller
	db *gorm.DB
}

func NewTagsController(db *gorm.DB) *TagsController {
	return &TagsController{db: db}
}

func (ctl TagsController) GetAll(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	tags := []models.Tag{}
	ctl.db.Model(&user).Related(&tags)

	ctl.SuccessResponse(c, gin.H{
		"tags": tags,
	})
}

func (ctl TagsController) Get(c *gin.Context) {
	fmt.Println(ctl, c)
}

type createTagJSON struct {
	Name string `json:"name" binding:"required"`
}

func (ctl TagsController) Create(c *gin.Context) {
	var json createTagJSON
	if c.BindJSON(&json) != nil {
		ctl.ErrorResponse(c, http.StatusBadRequest, "パラメータが不正です")
		return
	}

	user := c.MustGet("user").(models.User)
	tag := models.Tag{
		UserID: user.ID,
		Name:   json.Name,
	}
	ctl.db.Where(&tag).First(&tag)

	if tag.ID > 0 {
		ctl.ErrorResponse(c, http.StatusBadRequest, "既に存在するタグです")
		return
	}

	ctl.db.Create(&tag)
	ctl.SuccessResponse(c, gin.H{
		"tag": tag,
	})
}

func (ctl TagsController) Update(c *gin.Context) {
	fmt.Println(ctl, c)
}

func (ctl TagsController) Delete(c *gin.Context) {
	fmt.Println(ctl, c)
}
