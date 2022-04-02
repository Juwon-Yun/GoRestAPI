package dataBase

import (
	"github.com/Juwon-Yun/goRest/constants"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {
	URL := constants.Constant()
	_, err := gorm.Open(mysql.Open(URL), &gorm.Config{})

	if err != nil {
		panic("could not connect to the dataBase")
	}
}
