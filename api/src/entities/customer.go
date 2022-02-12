package entities

import (
	"time"
)

//Customer -> Represent a entity customer
type Customer struct {
	ID        uint64    `json:"id" gorm:"primaryKey;autoIncrement"`
	Login     string    `json:"login" gorm:"unique;type:varchar(15)"`
	FirstName string    `json:"firstname" gorm:"type:varchar(150)"`
	LastName  string    `json:"lastname" gorm:"type:varchar(150)"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	Active    bool      `json:"active" gorm:"default:true"`
	Addresses []Address `json:"addresses" gorm:"foreignKey:CustomerID"`
	Emails    []Email   `json:"emails" gorm:"foreignKey:CustomerID"`
	Phones    []Phone   `json:"phones" gorm:"foreignKey:CustomerID"`
}

//TableName -> gorm config
func (Customer) TableName() string {
	return "customers"
}

//Validate -> validate instance
func (c *Customer) Validate() (map[string]string, error) {
	errors := make(map[string]string)

	if c.Login == "" || len(c.Login) > 15 {
		errors["INVALID_LOGIN"] = "is mandatory and must contain a maximum of 15 characters"
	}

	if len(c.FirstName) > 150 {
		errors["INVALID_FIRSTNAME"] = "must contain a maximum of 150 characters"
	}

	if len(c.LastName) > 150 {
		errors["INVALID_LASTNAME"] = "must contain a maximum of 150 characters"
	}

	return errors, nil
}
