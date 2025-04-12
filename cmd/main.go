package main

import (
	"isg_API/db"
	"isg_API/routes"
	"log"
)

func main() {
	db.InitDB()
	db.Migrate()
	r := routes.SetupRoutes()

	log.Println("Server starting on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
