package request

type Register struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=32"`
	Name     string `json:"name" validate:"required,min=3,max=108"`
}

type Login struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}

type Logout struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}

type Refresh struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
	AccessToken  string `json:"accessToken" validate:"required"`
}
