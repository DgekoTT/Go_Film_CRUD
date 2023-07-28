package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go_crud/initializers"
	"go_crud/models"
	"net/http"
	"os"
	"time"
)

// Функция проверки наличия токена в куки
func getTokenFromCookie(c *gin.Context) (string, error) {
	tokenString, err := c.Cookie("Auth")
	return tokenString, err
}

// Функция проверки авторизации
func checkAuthorization(c *gin.Context, tokenString string) bool {
	if tokenString == "" {
		return false
	}
	// Здесь можно добавить другие проверки токена, если необходимо.
	return true
}

// Функция, возвращающая ошибку "Вы не авторизован"
func returnUnauthorizedError(c *gin.Context, message string) {
	c.AbortWithStatus(http.StatusUnauthorized)
	c.JSON(http.StatusUnauthorized, gin.H{
		"message": message,
	})
}

func ParseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func RequireAuth(c *gin.Context) {
	tokenString, err := getTokenFromCookie(c)
	if err != nil || !checkAuthorization(c, tokenString) {
		returnUnauthorizedError(c, "у вас нет токена")
		return
	}

	token, err := ParseToken(tokenString)

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			returnUnauthorizedError(c, "ваш токен просрочен")
		}
		var user models.User
		initializers.DB.First(&user, claims["userId"])

		if user.ID == 0 {
			returnUnauthorizedError(c, "такого пользователя не существет")
		}

		c.Set("user", user)
		c.Next()
	}
}
