package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
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
			jwtAbort(c, "Tokenが無効です")
			return
		}

		if time.Now().Unix() > claims.ExpiresAt {
			jwtAbort(c, "有効期限の切れたTokenです")
			return
		}

		c.Next()
	}
}
