package database

import (
	"fmt"
	"waysbuck/models"
	"waysbuck/pkg/mysql"
)

// Automatic Migration if Running App
func RunMigration() {
	err := mysql.DB.AutoMigrate(
		&models.Product{},
		&models.Topping{},
		&models.User{},
		&models.Profile{},
		&models.Transaction{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}
