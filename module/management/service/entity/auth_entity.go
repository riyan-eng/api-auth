package entity

type LoginEntity struct {
	Name               string `json:"name"`
	AccessToken        string `json:"access_token"`
	AccessTokenExpired int64  `json:"access_token_expired"`
}

type Register struct {
	UserName string
	Password string
}
