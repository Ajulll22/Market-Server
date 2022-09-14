package models

type Category struct {
	Id_category   int    `json:"id_category,omitempty" gorm:"primaryKey"`
	Nama_category string `json:"nama_category,omitempty"`
}
