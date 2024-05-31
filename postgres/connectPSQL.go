package postgress

import (
	struct_test "example.com/n11/struct"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDB() (*gorm.DB,error){
	var err error
	db, err = gorm.Open(postgres.Open(
		`postgres://postgres:salom@localhost:5432/vazifa1?sslmode=disable`))
	
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&struct_test.User{})

	return db, nil
}