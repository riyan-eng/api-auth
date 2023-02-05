package entity

type LoginEntity struct {
	Name               string `json:"name"`
	AccessToken        string `json:"access_token"`
	AccessTokenExpired int64  `json:"access_token_expired"`
	RefreshToken       string `json:"refresh_token"`
}

type Refresh struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type Register struct {
	UserName string
	Password string
}
