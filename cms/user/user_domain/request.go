package user_domain

type RequestUserInsert struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Address  string `json:"address" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
}

type RequestUserUpdate struct {
	Name    string `json:"name" `
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

type RequsetUserUpdateEmail struct {
	Email string `json:"email" validate:"required"`
}
