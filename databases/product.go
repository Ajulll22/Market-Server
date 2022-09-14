package databases

type Product struct {
	Id_product        int    `gorm:"primaryKey;type:int(11)"`
	Nama_product      string `gorm:"type:varchar(50)"`
	Deskripsi_product string `gorm:"type:text"`
	Gambar_product    string `gorm:"type:varchar(50)"`
	Harga_product     int    `gorm:"type:int(50)"`
	Url_product       string `gorm:"type:varchar(100)"`
	Status_product    bool   `gorm:"default:true"`
	Id_category       int
	Cart              []Cart     `gorm:"foreignKey:Id_product;references:Id_product;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Trx_item          []Trx_item `gorm:"foreignKey:Id_product;references:Id_product;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
