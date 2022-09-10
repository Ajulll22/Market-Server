package databases

type Category struct {
	Id_category   int       `gorm:"primaryKey;type:int(11)"`
	Nama_category string    `gorm:"type:varchar(50)"`
	Product       []Product `gorm:"foreignKey:Id_category;references:Id_category;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
