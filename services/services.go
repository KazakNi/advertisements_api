package services

import (
	database "adv/db"
	"adv/models"
	"database/sql"
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

func GetAdvertisement(fields string, id string, db *sqlx.DB) (interface{}, error) {
	var advertisement models.AdvertisementByID
	var adv_without_fields models.Advertisement
	if fields == "yes" {
		err := db.Get(&advertisement, "SELECT name, price, photos[1] FROM advertisements WHERE id = $1", id)
		if err != nil {
			log.Printf("error while querying: %s", err)
			return nil, err
		}
		return advertisement, nil
	} else {
		err := db.Get(&adv_without_fields, "SELECT name, price, description, photos FROM advertisements WHERE id = $1", id)
		if err != nil {
			log.Printf("error while querying: %s", err)
			return nil, err
		}
		return adv_without_fields, nil
	}

}

var dummy_sql = "INSERT INTO advertisements (name, description, price, created_at, photos) VALUES ($1, $2, $3, $4, $5)"

func PostAdvertisement(adv models.PostAdvertisement, db *sqlx.DB) (adv_id models.ResponsePost, err error) {
	var lastInsertId int

	row := db.QueryRow("INSERT INTO advertisements (name, description, price, created_at, photos) VALUES ($1, $2, $3, $4, $5) RETURNING id", adv.Name, adv.Description, adv.Price, database.Creation_time, adv.Photos)
	err = row.Scan(&lastInsertId)
	if err != nil {
		log.Printf("error while insert values: %s", err)
		return models.ResponsePost{Statuscode: 400}, err
	}
	return models.ResponsePost{Id: lastInsertId, Statuscode: 201}, nil
}

func IsUserExists(email string, db *sqlx.DB) (exists bool, err error) {
	var id int = -1
	err = db.Get(&id, "SELECT id FROM users WHERE email = $1", email)
	if err == sql.ErrNoRows {
		return false, nil
	} else if err != nil {
		log.Printf("error while querying: %s", err)
		return false, err
	} else {
		return true, nil
	}
}

func CreateUser(user models.User, db *sqlx.DB) (user_id models.ResponsePost, err error) {
	var id int
	user.HashPassword(user.Password)
	row := db.QueryRow("INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id", user.Username, user.Email, user.Password)
	err = row.Scan(&id)
	if err != nil {
		log.Printf("error while insert values: %s", err)
		return models.ResponsePost{Statuscode: 400}, err
	}
	return models.ResponsePost{Id: id, Statuscode: 201}, nil

}

func GetUserByEmail(email string, db *sqlx.DB) (user models.User, err error) {
	var current_user models.User
	err = db.Get(&current_user, "SELECT id, name, email, password FROM users WHERE email = $1", email)
	if err == sql.ErrNoRows {
		return models.User{}, nil
	} else if err != nil {
		log.Printf("error while querying user: %s", err)
		return current_user, err
	} else {
		return current_user, nil
	}
}
