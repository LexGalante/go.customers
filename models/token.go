package models

//AuthToken -> represent jwk
type AuthToken struct {
	AccessToken string `json:"access_token"`
	ExpiresAt   int64  `json:"expires_at"`
}
