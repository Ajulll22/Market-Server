package models

type Trx_item struct {
	Id_item     int     `json:"id_item,omitempty" gorm:"primaryKey"`
	Jumlah_item int     `json:"jumlah_item,omitempty"`
	Harga_item  int     `json:"harga_item,omitempty"`
	Id_product  int     `json:"id_product,omitempty"`
	Id_trx      int     `json:"id_trx,omitempty"`
	Product     Product `gorm:"foreignKey:Id_product;references:Id_product;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"product,omitempty"`
}
