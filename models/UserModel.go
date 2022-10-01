package models

type User struct {
	Id_user       int    `json:"id_user,omitempty" gorm:"primaryKey"`
	Nama_user     string `json:"nama_user,omitempty"`
	Email_user    string `json:"email_user,omitempty"`
	Alamat_user   string `json:"alamat_user,omitempty"`
	Level_user    int    `json:"level_user,omitempty" gorm:"default:1"`
	Status_user   int    `json:"status_user,omitempty" gorm:"default:1"`
	Password_user string `json:"-"`
}
