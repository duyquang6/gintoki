package dependency

import (
	"fmt"
	"gintoki/config"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

func NewMySQLConnection() (*sqlx.DB, error) {
	db, err := sqlx.Connect("mysql",
		fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
			config.AppConfig.Database.User, config.AppConfig.Database.Password,
			config.AppConfig.Database.Host, config.AppConfig.Database.Port, config.AppConfig.Database.Name))
	if err != nil {
		log.Fatalln(err)
	}

	return db, nil
}

func NewPostgresSQLConnection() (*sqlx.DB, error) {
	databaseConfig := config.AppConfig.Database
	db, err := sqlx.Connect("postgres",
		fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			databaseConfig.Host, databaseConfig.Port, databaseConfig.User,
			databaseConfig.Password, databaseConfig.Name))
	db.SetConnMaxIdleTime(time.Duration(databaseConfig.MaxIdleTimeConnection * float64(time.Second)))
	db.SetConnMaxLifetime(time.Duration(databaseConfig.MaxLifeTimeConnection * float64(time.Second)))
	db.SetMaxOpenConns(databaseConfig.MaxConnection)
	db.SetMaxIdleConns(databaseConfig.MinConnection)
	if err != nil {
		log.Fatalln(err)
	}

	return db, nil
}

func Close(db interface{}) {
	switch connection := db.(type) {
	case *sqlx.DB:
		connection.Close()
	}
}
