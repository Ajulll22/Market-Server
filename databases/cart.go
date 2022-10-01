package databases

type Cart struct {
	Id_cart     int `gorm:"primaryKey;type:int(11)"`
	Jumlah      int `gorm:"type:int(11)"`
	Total_harga int `gorm:"type:int(100)"`
	Id_product  int
	Id_user     int
}
