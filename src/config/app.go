package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Connection() *gorm.DB {
	db, err := gorm.Open("mysql", "user007:justpas/1@/bookstore?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		panic(err)
	}
	return db
}
