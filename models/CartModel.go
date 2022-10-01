package models

type Cart struct {
	Id_cart     int     `json:"id_cart,omitempty" gorm:"primaryKey"`
	Jumlah      int     `json:"jumlah,omitempty"`
	Total_harga int     `json:"total_harga"`
	Id_product  int     `json:"id_product,omitempty"`
	Id_user     int     `json:"id_user,omitempty"`
	Product     Product `gorm:"foreignKey:Id_product;references:Id_product;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"product,omitempty"`
}
