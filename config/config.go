package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase() (*gorm.DB, error) {
	var err error
	dsn := "root:@rif123#@tcp(127.0.0.1:3306)/crowfunding_miniproject?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed Connect Databases")
	}
	return db, nil

}
