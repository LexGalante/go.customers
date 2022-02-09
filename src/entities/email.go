package entities

import (
	"github.com/lexgalante/go.customers/src/utils"
)

//Email -> Represent a entity customer_email
type Email struct {
	ID         uint64 `json:"id" gorm:"primaryKey;autoIncrement"`
	CustomerID uint64 `json:"customer_id"`
	Principal  bool   `json:"principal" gorm:"default:false"`
	Email      string `json:"email" gorm:"type:varchar(250)"`
}

//Name -> gorm config
func (Email) Name() string {
	return "email"
}

//TableName -> gorm config
func (Email) TableName() string {
	return "customers_emails"
}

//Validate -> validate instance
func (e *Email) Validate() (map[string]string, error) {
	errors := make(map[string]string)

	if e.Email == "" || len(e.Email) > 250 || !utils.IsEmailValid(e.Email) {
		errors["INVALID_EMAIL"] = "is mandatory and must contain a maximum of 250 characters and valid e-mail"
	}

	return errors, nil
}
