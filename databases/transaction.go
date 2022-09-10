package databases

type Transaction struct {
	Id_trx    int `gorm:"primaryKey;type:int(11)"`
	Total_trx int `gorm:"type:int(50)"`
	Id_user   int
	Trx_item  []Trx_item `gorm:"foreignKey:Id_trx;references:Id_trx;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
