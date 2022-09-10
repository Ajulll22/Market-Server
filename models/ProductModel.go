package models

type Product struct {
	Id_product        int      `json:"id_product,omitempty"`
	Nama_product      string   `json:"nama_product,omitempty"`
	Deskripsi_product string   `json:"deskripsi_product,omitempty"`
	Gambar_product    string   `json:"gambar_product,omitempty"`
	Url_product       string   `json:"url_product,omitempty"`
	Status_product    bool     `json:"status_product,omitempty"`
	Id_category       int      `json:"id_category,omitempty"`
	Category          Category `gorm:"foreignKey:Id_category;references:Id_category;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"category,omitempty"`
}
