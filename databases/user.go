package databases

type User struct {
	Id_user     int           `gorm:"primaryKey;type:int(11)"`
	Nama_user   string        `gorm:"type:varchar(50)"`
	Email_user  string        `gorm:"type:varchar(50)"`
	Cart        []Cart        `gorm:"foreignKey:Id_user;references:Id_user;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Transaction []Transaction `gorm:"foreignKey:Id_user;references:Id_user;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
