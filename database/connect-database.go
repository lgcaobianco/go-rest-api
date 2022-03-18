package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)



var (
	DB *gorm.DB
	err error
)

func Connect(){
	str := "host=localhost user=dummyroot password=dummypassword dbname=dummyname port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(str))
	if err != nil {
		
	}

}
