package databases

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Conect(url string) {
	connection, err := gorm.Open(mysql.Open(url), &gorm.Config{PrepareStmt: true})

	if err != nil {
		log.Fatalln(err)
	}

	DB = connection
	connection.AutoMigrate(&User{}, &Category{}, &Product{}, &Cart{}, &Transaction{}, &Trx_item{})
}
