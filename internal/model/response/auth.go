package response

type Register struct {
	UserId  string `json:"userId"`
	Message string `json:"message"`
}

type Auth struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	ExpiresAt    string `json:"expiresAt"`
}
