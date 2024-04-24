package jwtServer

import (
	"net/http"
	"restApi/pkg/logger"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("your-secret-key") // 请使用安全的密钥

// 创建token
func CreateToken(userID int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 1).Unix(),
	})

	logger.Info(map[string]interface{}{"token": token})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		logger.Info(map[string]interface{}{"err": err})
		return "", err
	}
	logger.Info(map[string]interface{}{"tokenString": token})
	return tokenString, nil
}

// 验证token中间件
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var tokenString string

		// 从请求头中获取token
		authHeader := c.Request.Header.Get("Authorization")
		logger.Info(map[string]interface{}{"authHeader": authHeader})

		if len(authHeader) > 0 {
			splitToken := strings.Split(authHeader, "Bearer ")
			if len(splitToken) == 2 {
				tokenString = splitToken[1]
			}
		}

		// 解析token
		token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			logger.Info(map[string]interface{}{"jwt.ParseWithClaims err, err=": err})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userID, ok := claims["user_id"].(float64)
			if !ok {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
				return
			}
			logger.Info(map[string]interface{}{"userID": userID})
			c.Set("user_id", int64(userID))
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token is not valid"})
			return
		}

		c.Next()
	}
}
