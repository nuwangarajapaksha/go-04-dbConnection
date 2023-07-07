package models

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:ass34@tcp(localhost:3307)/go_test_db"))
	if err != nil {
		panic(err)
	}
	database.AutoMigrate(&Product{})
	DB = database
}

func ConnectionDb() {
	db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	insert, err := db.Query("INSERT INTO test VALUES ( 2, 'TEST' )")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
}
