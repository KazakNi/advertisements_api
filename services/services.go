package services

import (
	database "adv/db"
	"adv/models"
	"log"
	"strconv"

	"github.com/jmoiron/sqlx"
)

func GetAdvWithoutSorting(page string, db *sqlx.DB) models.ResponseAllAdvertisements {
	page_param := 0
	if page != "" {
		page_param, err := strconv.Atoi(page)
		if err != nil {
			log.Fatalf("invalid http parameter: 'page'=%v\n, error:%s", page_param, err)
		}
		if page_param < 2 || database.CountRows(database.DB) <= 10 {
			page_param = 0
		} else {
			page_param = (page_param - 1) * 10
		}
	}
	var advertisements = []models.AllAdvertisements{}

	err := db.Select(&advertisements, "SELECT name, price, photos[1] FROM advertisements LIMIT 10 OFFSET $1", page_param)
	if err != nil {
		log.Fatalf("error while querying: %s", err)
	}
	result := models.ResponseAllAdvertisements{AllAdvertisements: advertisements, NextPage: page_param + 1}
	return result
}
