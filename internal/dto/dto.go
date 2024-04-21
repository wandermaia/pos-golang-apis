package dto

type CreateProductInput struct {
	Name  string  `jason:"name"`
	Price float64 `jason:"price"`
}

type CreateUserInput struct {
	Name     string `jason:"name"`
	Email    string `jason:"email"`
	Password string `jason:"password"`
}

type GetJWTInput struct {
	Email    string `jason:"email"`
	Password string `jason:"password"`
}

type GetJWTOutput struct {
	AccessToken string `json:"access_token"`
}
