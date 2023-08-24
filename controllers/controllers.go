package controllers

import (
	database "adv/db"
	"adv/models"
	"adv/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getParams(c *gin.Context) (string, string, string, string) {
	page, priceSort, dateSort, fields := c.Param("page"), c.Param("price"), c.Param("date"), c.Param("fields")
	return page, priceSort, dateSort, fields
}

func GetAllAdvertisements(c *gin.Context) {
	result := models.ResponseAllAdvertisements{}
	page, priceSort, dateSort, _ := getParams(c)
	if priceSort == "" && dateSort == "" {
		result = services.GetAdvWithoutSorting(page, database.DB)
	}
	c.JSON(http.StatusOK, result)
}

func GetAdvertisementByID(c *gin.Context) {

}

func PostAdvertisement(c *gin.Context) {

}
