package config

import (
	"log"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
)

var(
	db * gorm.DB
)

func Connect(){
	d,err:=gorm.Open("mysql","devUser:devPassword@tcp(127.0.0.1:3306)/bookstore-management?charset=utf8&parseTime=True&loc=Local")
	if err != nil{
		log.Fatal("Database connection Failed!",err)
	}

	db=d
}

func GetDB() *gorm.DB{
	return db
}