package services

import (
	database "adv/db"
	"fmt"
)

func GetAdvWithoutSorting(page string) {
	if page == "" {
		page = "1"
	}
	d, err := database.InitDB()
	fmt.Println(d, err, page)
}
