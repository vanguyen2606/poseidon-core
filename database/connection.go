package database

import (
	"fmt"

	"github.com/vannguyen2606/poseidon-core/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(env map[string]string) {
	connection, err := gorm.Open(mysql.Open(fmt.Sprintf("%v:%v@/apply_prod", env["SQL_HOST"], env["SQL_PASSWORD"])), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
}
