package controllers

import (
	"adv/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllAdvertisements(c *gin.Context) {
	priceSort, dateSort := c.Param("price"), c.Param("date")
	page := c.Param("page")
	if priceSort == "" && dateSort == "" {
		services.GetAdvWithoutSorting(page)
	}
	c.String(http.StatusOK, "Hello")

}

func GetAdvertisementByID(c *gin.Context) {

}

func PostAdvertisement(c *gin.Context) {

}
