package databases

type Trx_item struct {
	Id_item     int `gorm:"primaryKey;type:int(11)"`
	Jumlah_item int `gorm:"type:int(11)"`
	Harga_item  int `gorm:"type:int(50)"`
	Id_product  int
	Id_trx      int
}
