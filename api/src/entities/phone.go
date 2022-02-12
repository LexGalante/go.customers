package entities

//Phone -> Represent a entity customer_phone
type Phone struct {
	ID         uint64 `json:"id" gorm:"primaryKey;autoIncrement"`
	CustomerID uint64 `json:"customer_id"`
	Principal  bool   `json:"principal" gorm:"default:false"`
	Ddi        int    `json:"ddi" gorm:"default:55"`
	Phone      string `json:"phone" gorm:"type:varchar(20)"`
}

//Name -> gorm config
func (Phone) Name() string {
	return "phone"
}

//TableName -> gorm config
func (Phone) TableName() string {
	return "customers_phones"
}

//Validate -> validate instance
func (p *Phone) Validate() (map[string]string, error) {
	errors := make(map[string]string)

	if p.Ddi < 10 || p.Ddi > 999 {
		errors["INVALID_DDI"] = "is mandatory and must greater than equal 10 and less than equal 999"
	}

	if p.Phone == "" || len(p.Phone) < 7 || len(p.Phone) > 20 {
		errors["INVALID_PHONE"] = "is mandatory and must contain a minimum characters 7 and maximum of 20 characters"
	}

	return errors, nil
}
