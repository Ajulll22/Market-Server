package requests

type LoginRequest struct {
	Email_user    string `json:"email_user" validate:"required,email"`
	Password_user string `json:"password_user" validate:"required"`
}
