package models

//Address -> represent https://viacep.com.br/ address
type Address struct {
	ID         string `json:"ibge"`
	Cep        string `json:"cep"`
	StreetName string `json:"logradouro"`
	Details    string `json:"complemento"`
	District   string `json:"bairro"`
	City       string `json:"localidade"`
	Uf         string `json:"uf"`
}
