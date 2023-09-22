package controllers

import (
	database "adv/db"
	"adv/models"
	"adv/services"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

func getParams(c *gin.Context) (string, string, string, string) {
	page, priceSort, dateSort, fields := c.DefaultQuery("page", "0"), c.Query("price"), c.Query("date"), c.Query("fields")
	return page, priceSort, dateSort, fields
}

func GetAllAdvertisements(c *gin.Context) {
	result := models.ResponseAllAdvertisements{}
	page, priceSort, dateSort, _ := getParams(c)
	if priceSort == "" && dateSort == "" {
		result = services.GetAdvWithoutSorting(page, database.DB)
	} else {
		result = services.GetAdvSorting(page, []string{priceSort, dateSort}, database.DB)
	}
	c.JSON(http.StatusOK, result)
}

func GetAdvertisementByID(c *gin.Context) {
	_, _, _, fields := getParams(c)
	id := c.Param("id")
	result, err := services.GetAdvertisement(fields, id, database.DB)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "искомый объект не найден"})
	} else {
		c.JSON(http.StatusOK, result)
	}

}

func PostAdvertisement(c *gin.Context) {
	var adv models.PostAdvertisement
	if err := c.ShouldBindJSON(&adv); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "неверный формат данных", "error": err.Error()})
	} else {
		result, err := services.PostAdvertisement(adv, database.DB)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "ошибка сервера"})
		} else {
			c.JSON(http.StatusCreated, result)
		}
	}
}

func SignIn(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "неверный формат данных", "error": err.Error()})
	} else {
		result, err := services.GetUserByEmail(user.Email, database.DB)
		if result.Email == "" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "Почта либо пароль не совпадают!"})
		} else if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "ошибка сервера", "error": err.Error()})
		} else if result.CheckPassword(user.Password) != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "Почта либо пароль не совпадают!"})
		} else {
			claims := models.CustomClaims{jwt.RegisteredClaims{ID: result.Id, ExpiresAt: jwt.NewNumericDate(time.Now().Add(72 * time.Hour))}}
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			err := godotenv.Load()
			if err != nil {
				log.Fatal("Error loading .env file")
			}
			token_string, err := token.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
			if err != nil {
				log.Fatalln("Token key was not discovered!", err)
				c.AbortWithError(500, err)
			}
			c.Header("Authorization", "Bearer "+token_string)
		}
	}
}

func SignUp(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "неверный формат данных", "error": err.Error()})
	} else {
		userExists, err := services.IsUserExists(user.Email, database.DB)
		if userExists {
			c.JSON(http.StatusBadRequest, gin.H{"status": "пользователь с такой почтой уже существует!"})
		} else if err != nil {
			log.Fatalln(err)
			c.AbortWithStatus(http.StatusInternalServerError)
		} else {
			result, err := services.CreateUser(user, database.DB)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"message": "ошибка сервера"})
			} else {
				c.JSON(http.StatusCreated, result)
			}
		}
	}
}
