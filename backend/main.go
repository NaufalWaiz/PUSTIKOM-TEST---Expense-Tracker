package main

import (
	"expense-tracker/config"
	"expense-tracker/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("File .env tidak ditemukan, pastikan environment sudah diset manual")
	}

	config.ConnectDB()
	defer config.DB.Close()

	r := gin.Default()
	routes.ExpenseRoutes(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Server gagal dijalankan: %v", err)
	}
}
