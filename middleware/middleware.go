package middleware

import (
	"adv/models"
	"errors"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

func load_secret() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("TOKEN_SECRET")
}

func ValidateIdParam(c *gin.Context) {
	var currentId models.IdPath
	if err := c.ShouldBindUri(&currentId); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"Ошибка валидации url-пути": "id должен быть целым числом"})
		return
	}
	return
}

func AuthRequiredCheck(c *gin.Context) {
	AuthHeader := c.Request.Header.Get("Authorization")
	token := strings.Fields(AuthHeader)[1]
	err := ParseToken(token)
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{"error": "Invalid or missed token"})
	}
	return
}

func SetToken(c *gin.Context, claims models.CustomClaims) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	token_string, err := token.SignedString([]byte(load_secret()))
	if err != nil {
		log.Fatalln("Token key was not discovered!", err)
		c.AbortWithError(500, err)
	}
	c.Header("Authorization", "Bearer "+token_string)
}

func ParseToken(mytoken string) error {

	token, err := jwt.ParseWithClaims(mytoken, &models.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(load_secret()), nil
	})
	if err != nil {
		log.Print(err)
		return errors.New("Токен невалиден!")
	}
	if claims, ok := token.Claims.(*models.CustomClaims); ok && token.Valid {
		if claims.RegisteredClaims.ExpiresAt.Unix() < time.Now().Unix() {
			return errors.New("Токен истёк!")
		}
	} else {
		log.Println(err)
		return errors.New("Требуется авторизация!")
	}
	return nil
}
