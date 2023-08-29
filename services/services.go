package services

import (
	database "adv/db"
	"adv/models"
	"log"
	"strconv"

	"github.com/jmoiron/sqlx"
)

var advertisements = []models.Advertisement{}

func getCurrentPage(page string) (int, int) {
	page_response, err := strconv.Atoi(page)
	page_param := page_response

	if err != nil {
		log.Fatalf("invalid http parameter: 'page'=%v\n, error:%s", page_param, err)
	}
	if page_param < 2 || database.CountRows(database.DB) <= 10 {
		page_param = 0
	} else {
		page_param = (page_param - 1) * 10
	}
	return page_response, page_param
}

func GetAdvWithoutSorting(page string, db *sqlx.DB) models.ResponseAllAdvertisements {
	page_response, page_param := getCurrentPage(page)
	err := db.Select(&advertisements, "SELECT name, price, photos[1] FROM advertisements LIMIT 10 OFFSET $1", page_param)
	if err != nil {
		log.Fatalf("error while querying: %s", err)
	}
	result := models.ResponseAllAdvertisements{AllAdvertisements: advertisements, NextPage: page_response + 1}
	return result
}

func GetAdvSorting(page string, criteria []string, db *sqlx.DB) models.ResponseAllAdvertisements {
	var result models.ResponseAllAdvertisements
	page_response, page_param := getCurrentPage(page)
	sortByPrice, sortByDate := criteria[0], criteria[1]

	if sortByPrice == "ASC" {
		err := db.Select(&advertisements, "SELECT name, price, photos[1] FROM advertisements ORDER BY price LIMIT 10 OFFSET $1", page_param)
		if err != nil {
			log.Fatalf("error while querying: %s", err)
		}
		result = models.ResponseAllAdvertisements{AllAdvertisements: advertisements, NextPage: page_response + 1}
		return result
	} else if sortByPrice == "DESC" {
		err := db.Select(&advertisements, "SELECT name, price, photos[1] FROM advertisements ORDER BY price DESC LIMIT 10 OFFSET $1", page_param)
		if err != nil {
			log.Fatalf("error while querying: %s", err)
		}
		result = models.ResponseAllAdvertisements{AllAdvertisements: advertisements, NextPage: page_response + 1}
		return result
	}

	if sortByDate == "ASC" {
		err := db.Select(&advertisements, "SELECT name, price, photos[1] FROM advertisements ORDER BY created_at LIMIT 10 OFFSET $1", page_param)
		if err != nil {
			log.Fatalf("error while querying: %s", err)
		}
		result = models.ResponseAllAdvertisements{AllAdvertisements: advertisements, NextPage: page_response + 1}
		return result
	} else if sortByDate == "DESC" {
		err := db.Select(&advertisements, "SELECT name, price, photos[1] FROM advertisements ORDER BY created_at DESC LIMIT 10 OFFSET $1", page_param)
		if err != nil {
			log.Fatalf("error while querying: %s", err)
		}
		result = models.ResponseAllAdvertisements{AllAdvertisements: advertisements, NextPage: page_response + 1}
		return result
	}
	return result
}
