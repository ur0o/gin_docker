package database

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var db *gorm.DB

func Init() {
	var err error;
	// TODO: DB情報をgit管理外のファイルに切り出し
	db, err = gorm.Open(mysql.Open("root:pass@tcp(mysql:3306)/gin_docker?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})

	if err != nil {
		panic(err)
	}
}

func GetDB() *gorm.DB {
	return db
}

func Close() {
	d, _ := db.DB()
	d.Close()
}