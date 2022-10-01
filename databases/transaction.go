package databases

import "time"

type Transaction struct {
	Id_trx     int    `gorm:"primaryKey;type:int(11)"`
	Total_trx  int    `gorm:"type:int(100)"`
	Alamat_trx string `gorm:"type:varchar(100)"`
	Status_trx int    `gorm:"type:int(5)"`
	Created_at time.Time
	Id_user    int
	Trx_item   []Trx_item `gorm:"foreignKey:Id_trx;references:Id_trx;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
