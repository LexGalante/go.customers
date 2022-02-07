package entities

import (
	"fmt"
	"time"

	"gorm.io/gorm"
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

//Address -> Represent a entity customer_address
type Address struct {
	CustomerID   uint64 `json:"customer_id"`
	Principal    bool   `json:"principal" gorm:"default:false"`
	PostalCode   string `json:"postal_code" gorm:"type:varchar(20)"`
	StreetName   string `json:"street_name" gorm:"type:varchar(250)"`
	StreetNumber string `json:"street_number" gorm:"type:varchar(5)"`
	District     string `json:"district" gorm:"type:varchar(150)"`
	City         string `json:"city" gorm:"type:varchar(150)"`
	Country      string `json:"country" gorm:"type:varchar(2);default:BR"`
}

//Email -> Represent a entity customer_email
type Email struct {
	CustomerID uint64 `json:"customer_id"`
	Principal  bool   `json:"principal" gorm:"default:false"`
	City       string `json:"email" gorm:"type:varchar(250)"`
}

//Phone -> Represent a entity customer_phone
type Phone struct {
	CustomerID uint64 `json:"customer_id"`
	Principal  bool   `json:"principal" gorm:"default:false"`
	Ddi        int    `json:"ddi" gorm:"default:55"`
	Phone      string `json:"phone" gorm:"type:varchar(20)"`
}

//TableName -> gorm config
func (Customer) TableName() string {
	return "customers"
}

//TableName -> gorm config
func (Address) TableName() string {
	return "customers_addresses"
}

//TableName -> gorm config
func (Email) TableName() string {
	return "customers_emails"
}

//TableName -> gorm config
func (Phone) TableName() string {
	return "customers_phones"
}

//Validate -> validate instance
func (c *Customer) Validate(db *gorm.DB) (map[string]string, error) {
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
	// when validate create mode
	if c.ID == 0 {
		var customer Customer
		db.Select("login").Where("login = ?", c.Login).Limit(1).Find(&customer)
		if customer.Login != "" {
			errors["LOGIN_ALREADY_EXIST"] = fmt.Sprintf("login %s already exists", customer.Login)
		}
	}

	return errors, nil
}
