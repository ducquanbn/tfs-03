package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBConnect() *gorm.DB {

	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/film"))
	if err != nil {
		fmt.Println(err.Error())
	}
	return db
}
