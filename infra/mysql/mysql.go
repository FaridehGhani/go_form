package mysql

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Connect() *gorm.DB {
	db, err := gorm.Open("mysql", os.Getenv("MYSQL_CONNECTION"))
	if err != nil {
		log.Fatalf("error connectiong to mysql: %s", err)
	}

	return db
}
