package controllers

import (
	database "adv/db"
	"adv/models"
	"adv/services"
	"net/http"

	"github.com/gin-gonic/gin"
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
