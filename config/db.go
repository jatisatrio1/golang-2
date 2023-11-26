package config

import (
	"fmt"
	"rest-api/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "db-go-assignment-2"
)

var (
	db  *gorm.DB
	err error
)

func Connect() {

	psqlInfo := fmt.Sprintf(`
	host=%s 
	port=%d 
	user=%s `+`
	password=%s 
	dbname=%s 
	sslmode=disable`,
		host, port, user, password, dbname)

	//ini koneksi menggunakan database/sql
	// db, err = sql.Open("postgres", psqlInfo)

	//ini koneksi menggunakan gorm
	db, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database")

	db.Debug().AutoMigrate(&model.Item{}, &model.Order{})

}

func GetDB() *gorm.DB {
	return db
}
