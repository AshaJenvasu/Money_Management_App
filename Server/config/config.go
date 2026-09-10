package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config โครงสร้างสำหรับเก็บค่าการตั้งค่าทั้งหมดของแอป
type Config struct {
	DBConnString string
}

// LoadConfig ทำหน้าที่อ่านค่าจาก .env แล้วประกอบเป็น Connection String
func LoadConfig() (*Config, error) {
	// 1. โหลดไฟล์ .env (ถ้าไม่มีไฟล์นี้ หรือรันบน Server จริง จะข้ามไปอ่านจาก OS Env แทน)
	err := godotenv.Load()
	if err != nil {
		// Note: ใน Production บน Cloud เราจะตั้งค่า Env ผ่าน OS โดยตรง จึงไม่เจอไฟล์ .env
		// ตรงนี้ใส่ Print เตือนไว้ได้ แต่ไม่จำเป็นต้องสั่ง return error ล่ม
		fmt.Println("Warning: .env file not found, fallback to OS environment variables")
	}

	// 2. ดึงค่าจาก Env Vars
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_SERVER")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSLMODE")

	// 3. ตรวจสอบข้อมูลจำเป็น (Validation Best Practice)
	if dbUser == "" || dbHost == "" || dbName == "" {
		return nil, fmt.Errorf("missing required database environment variables")
	}

	// 4. ประกอบ Connection String สำหรับ PostgreSQL (DSN Format)
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		dbUser, dbPassword, dbHost, dbPort, dbName, dbSSLMode,
	)

	return &Config{
		DBConnString: connStr,
	}, nil
}