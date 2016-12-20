package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/tsuyoshiwada/gin-sandbox/controllers"
	"github.com/tsuyoshiwada/gin-sandbox/middleware"
)

func makeResource(r gin.IRouter, ctl controllers.ResourceController) {
	r.GET("/", ctl.GetAll)
	r.GET("/:id", ctl.Get)
	r.POST("/", ctl.Create)
	r.PATCH("/", ctl.Update)
	r.DELETE("/:id", ctl.Delete)
}

func buildRoutes(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	// Auth
	authController := controllers.NewAuthController(db)
	auth := router.Group("/auth")
	{
		auth.POST("/", authController.Auth)
	}

	// Authentication required
	authorized := router.Group("/")
	authorized.Use(middleware.JWTMiddleware(db))
	{
		// Users
		userController := controllers.NewUsersController(db)
		users := authorized.Group("/users")
		{
			users.GET("/", userController.GetAll)
			users.GET("/:id", userController.Get)
		}

		// Articles
		articlesController := controllers.NewArticlesController(db)
		makeResource(authorized.Group("/articles"), articlesController)

		// Tags
		tagsController := controllers.NewTagsController(db)
		makeResource(authorized.Group("/tags"), tagsController)
	}

	return router
}
