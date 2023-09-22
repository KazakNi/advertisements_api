package models

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	_ "gopkg.in/go-playground/validator.v9"
)

type User struct {
	Id       string `json:"id" db:"id"`
	Username string `json:"name" db:"name"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password,omitempty" db:"password"`
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}
func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}

type IdPath struct {
	Id int `uri:"id" binding:"required,numeric"`
}

type Advertisement struct {
	Name        string     `json:"name"  db:"name"`
	Description string     `json:"description,omitempty"  db:"description"`
	Price       int        `json:"price"  db:"price"`
	Photos      string     `json:"photos" db:"photos"`
	Created_at  *time.Time `json:"created_at,omitempty"  db:"created_at"`
}

type AdvertisementByID struct {
	Name        string         `json:"name"  db:"name"`
	Description string         `json:"description,omitempty"  db:"description"`
	Price       int            `json:"price"  db:"price"`
	Photos      pq.StringArray `json:"photos" db:"photos"`
}

type ResponseAllAdvertisements struct {
	AllAdvertisements []Advertisement `json:"items"`
	NextPage          int             `json:"next_page"`
}

type PostAdvertisement struct {
	Name        string         `json:"name"  binding:"required,min=3" db:"name"`
	Description string         `json:"description"  binding:"required,min=5" db:"description"`
	Price       int            `json:"price"  binding:"required,gte=1" db:"price"`
	Photos      pq.StringArray `json:"photos" binding:"required,min=2" db:"photos"`
}

type ResponsePost struct {
	Id         int `json:"id,omitempty" binding:"required,numeric" db:"id"`
	Statuscode int `json:"status code"`
}

type CustomClaims struct {
	jwt.RegisteredClaims
}
