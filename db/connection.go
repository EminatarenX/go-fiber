package main

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	
)


func Connection () (*gorm.DB, error) {
	
	dsn := "root:naog7412@tcp(localhost:3306)/act2?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	db.Debug()

	return db, nil

}

func main() {
	Connection()
}