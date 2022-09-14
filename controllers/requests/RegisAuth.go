package requests

type RegisRequest struct {
	Nama_user     string `json:"nama_user" validate:"required"`
	Email_user    string `json:"email_user" validate:"required,email"`
	Password_user string `json:"password_user" validate:"required,min=5"`
}
