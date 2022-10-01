package models

import "time"

type Transaction struct {
	Id_trx     int        `json:"id_trx,omitempty" gorm:"primaryKey"`
	Total_trx  int        `json:"total_trx,omitempty"`
	Alamat_trx string     `json:"alamat_trx,omitempty"`
	Status_trx int        `json:"status_trx,omitempty" gorm:"default:1"`
	Created_at time.Time  `json:"created_at,omitempty"`
	Id_user    int        `json:"id_user,omitempty"`
	Trx_item   []Trx_item `gorm:"foreignKey:Id_trx;references:Id_trx;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"trx_item,omitempty"`
}
