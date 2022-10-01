package requests

type EditProfilRequest struct {
	Alamat_user string `json:"alamat_user" validate:"required"`
}
