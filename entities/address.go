package entities

//Address -> Represent a entity customer_address
type Address struct {
	ID           uint64 `json:"id" gorm:"primaryKey;autoIncrement"`
	CustomerID   uint64 `json:"customer_id" gorm:"type:int8"`
	Principal    bool   `json:"principal" gorm:"type:bool;default:false"`
	PostalCode   string `json:"postal_code" gorm:"type:varchar(20)"`
	StreetName   string `json:"street_name" gorm:"type:varchar(250)"`
	StreetNumber string `json:"street_number" gorm:"type:varchar(5)"`
	District     string `json:"district" gorm:"type:varchar(150)"`
	City         string `json:"city" gorm:"type:varchar(150)"`
	Country      string `json:"country" gorm:"type:varchar(2);default:BR"`
}

//Name -> gorm config
func (Address) Name() string {
	return "address"
}

//TableName -> gorm config
func (Address) TableName() string {
	return "customers_addresses"
}

//Validate -> validate instance
func (a *Address) Validate() (map[string]string, error) {
	errors := make(map[string]string)

	if a.PostalCode == "" || len(a.PostalCode) > 20 {
		errors["INVALID_POSTAL_CODE"] = "is mandatory and must contain a maximum of 20 characters"
	}

	if a.StreetName == "" || len(a.StreetName) > 250 {
		errors["INVALID_STREET_NAME"] = "is mandatory and must contain a maximum of 250 characters"
	}

	if len(a.StreetNumber) > 5 {
		errors["INVALID_STREET_NUMBER"] = "must contain a maximum of 5 characters"
	}

	if len(a.District) > 150 {
		errors["INVALID_DISTRICT"] = "must contain a maximum of 150 characters"
	}

	if len(a.City) > 150 {
		errors["INVALID_CITY"] = "must contain a maximum of 150 characters"
	}

	if a.Country == "" || len(a.Country) > 2 {
		errors["INVALID_COUNTRY"] = "is mandatory and must contain a maximum of 2 characters"
	}

	return errors, nil
}
