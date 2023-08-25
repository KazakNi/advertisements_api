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

}

func PostAdvertisement(c *gin.Context) {

}
