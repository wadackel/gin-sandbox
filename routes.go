package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/tsuyoshiwada/gin-sandbox/controllers"
	"github.com/tsuyoshiwada/gin-sandbox/middleware"
)

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
		articles := authorized.Group("/articles")
		{
			articles.GET("/", articlesController.GetAll)
			articles.POST("/", articlesController.Create)
			articles.DELETE("/:id", articlesController.Delete)
		}
	}

	return router
}
