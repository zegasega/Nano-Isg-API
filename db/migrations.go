package db

import (
	"isg_API/models"
	"log"
)

func Migrate() {
	err := DB.AutoMigrate(models.User{}, models.Project{}, models.Personel{}, models.Isg_Egitim{})

	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Database migration completed successfully")
}
