package databases

type Cart struct {
	Id_cart     int    `gorm:"primaryKey;type:int(11)"`
	Jumlah      int    `gorm:"type:int(11)"`
	Total_harga int    `gorm:"type:int(100)"`
	Keterangan  string `gorm:"type:varchar(75)"`
	Status_cart bool   `gorm:"default:true"`
	Id_product  int
	Id_user     int
}
