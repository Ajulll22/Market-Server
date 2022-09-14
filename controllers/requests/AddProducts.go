package requests

import "encoding/json"

type AddProductRequest struct {
	Nama_product      string      `form:"nama_product" validate:"required"`
	Deskripsi_product string      `form:"deskripsi_product" validate:"required"`
	Id_category       json.Number `form:"id_category" validate:"required,number"`
	Harga_product     json.Number `form:"harga_product" validate:"required,number"`
}
