package requests

import "encoding/json"

type EditCartRequest struct {
	Jumlah json.Number `json:"jumlah" validate:"required,number"`
}
