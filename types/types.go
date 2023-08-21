package types

import "time"

var User struct {
	name  string
	email string
}

var Advertisement struct {
	name        string
	description string
	price       int
	created_at  time.Time
}
