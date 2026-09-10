package main

import (
	"log"

	"github.com/AshaJenvasu/Money_Management_App/config"
)

func main() {
	// โหลดค่า Environment Variables ผ่าน LoadConfig ที่เราทำใน config.go
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// เรียกใช้ InitDB หรือ NewDatabaseConnection จากแพ็กเกจ config
	db, err := config.InitDB(cfg.DBConnString)
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	// Best Practice: ใช้ defer สั่งปิด Connection Pool เมื่อ main() ทำงานเสร็จ
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("Error closing database connection: %v", err)
		} else {
			log.Println("Database connection closed gracefully.")
		}
	}()

	//แสดงข้อความเมื่อเชื่อมต่อสำเร็จ
	log.Println("Successfully connected to the PostgreSQL database! 🎉")
}