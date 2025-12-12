package main

import (
	"github.com/Shakti-Tamang/crud-app/initializers"
	"github.com/Shakti-Tamang/crud-app/models"
)

func init() {
	initializers.ConnectDB()
	initializers.LoadEnvVariable()
}

func main() {
	initializers.DB.AutoMigrate(
		&models.Item{},
		&models.UserData{},
	)
}
