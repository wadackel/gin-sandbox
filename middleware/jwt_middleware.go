package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/tsuyoshiwada/gin-sandbox/models"
	"github.com/tsuyoshiwada/gin-sandbox/shared/jwtauth"
)

func jwtAbort(c *gin.Context, msg string) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"status":  "error",
		"message": msg,
	})
	c.Abort()
}

func JWTMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		fmt.Println(authHeader)
		if authHeader == "" {
			jwtAbort(c, "Authorizationヘッダーが含まれていません")
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			jwtAbort(c, "Authorizationヘッダーが無効です")
			return
		}

		claims, err := jwtauth.ParseToken(parts[1])
		if err != nil {
			jwtAbort(c, "無効なTokenです")
			return
		}

		if time.Now().Unix() > claims.ExpiresAt {
			jwtAbort(c, "有効期限の切れたTokenです")
			return
		}

		user := models.User{}
		db.First(&user, claims.UserID)

		if user.ID != claims.UserID {
			jwtAbort(c, "無効なTokenです")
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
