package models

import (
	"time"

	"github.com/lib/pq"
	_ "gopkg.in/go-playground/validator.v9"
)

type User struct {
	Username string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

/* func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.password = string(bytes)
	return nil
}
func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
*/

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

type ResponsePostAdv struct {
	Id         int `json:"id" binding:"required,numeric" db:"id"`
	Statuscode int `json:"status code"`
}

// func (a *Advertisement) ShowOptionalParams(visible bool) *Advertisement
/*func (ms MyStruct) MarshalJSON() ([]byte, error) {
    m := map[string]interface{}{} // ideally use make with the right capacity
    m["nickname"] = ms.Nickname
    m["email_address"] = ms.EmailAddress
    if ms.all {
        m["phone_number"] = ms.PhoneNumber
        m["mailing_address"] = ms.MailingAddress
    }
    return json.Marshal(m)
} */
// Optional struct for parameter "fields"
