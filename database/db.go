package database

import (
	"github.com/jmoiron/sqlx"
	"log"
	"sync"
)

var once sync.Once

// type global
var (
	instance *sqlx.DB
)

func NewDB() *sqlx.DB {
	once.Do(func() {
		db, err := sqlx.Connect("mysql", "mysqluser:mysqlpw@tcp(127.0.0.1:3306)/inventory?charset=utf8mb4&parseTime=True&loc=Local")
		if err != nil {
			log.Fatalln(err)
		}
		instance = db
	})
	return instance
}