package models

type User struct {
	Id_user       int    `json:"id_user,omitempty" gorm:"primaryKey"`
	Nama_user     string `json:"nama_user,omitempty"`
	Email_user    string `json:"email_user,omitempty"`
	Level_user    int    `json:"level_user,omitempty"`
	Password_user string `json:"-"`
}
