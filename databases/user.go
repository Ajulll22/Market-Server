package databases

type User struct {
	Id_user       int           `gorm:"primaryKey;type:int(11)"`
	Nama_user     string        `gorm:"type:varchar(50)"`
	Email_user    string        `gorm:"type:varchar(50)"`
	Password_user string        `gorm:"type:varchar(100)"`
	Alamat_user   string        `gorm:"type:varchar(100)"`
	Level_user    int           `gorm:"type:varchar(5);default:1"`
	Status_user   int           `gorm:"type:int(5);default:1"`
	Cart          []Cart        `gorm:"foreignKey:Id_user;references:Id_user;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Transaction   []Transaction `gorm:"foreignKey:Id_user;references:Id_user;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
