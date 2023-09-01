package database

import (
	"time"
)

var Creation_time = time.Now()
var schema = `
CREATE TABLE IF NOT EXISTS advertisements(
	id SERIAL PRIMARY KEY,
	name TEXT NOT NULL,
	description TEXT,
	price INT NOT NULL,
	created_at TIMESTAMP NOT NULL,
	photos TEXT[] NOT NULL
);`

var insert_query = "INSERT INTO advertisements (name, description, price, created_at, photos) VALUES ($1, $2, $3, $4, $5)"
