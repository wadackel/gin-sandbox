package controllers

import "github.com/gin-gonic/gin"

type ResourceController interface {
	GetAll(c *gin.Context)
	Get(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type Controller struct{}

func (ctl Controller) buildSuccessData(data interface{}) map[string]interface{} {
	return gin.H{
		"status": "success",
		"error":  false,
		"data":   data,
	}
}

func (ctl Controller) buildErrorData(message string) map[string]interface{} {
	return gin.H{
		"status":  "error",
		"error":   true,
		"message": message,
	}
}

func (ctl Controller) SuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(200, ctl.buildSuccessData(data))
}

func (ctl Controller) ErrorResponse(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, ctl.buildErrorData(message))
}
