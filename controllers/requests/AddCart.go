package requests

import "encoding/json"

type AddCartRequest struct {
	Id_product json.Number `json:"id_product" validate:"required,number"`
}
