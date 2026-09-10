package config

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib" // Blank import เพื่อ register driver ของ pgx
)

// InitDB ทำหน้าที่สร้าง Connection Pool และเช็คการเชื่อมต่อกับ Database
func InitDB(connStr string) (*sql.DB, error) {
	// 1. สร้าง Connection Pool
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// 2. ตั้งค่า Connection Pool Best Practices
	db.SetMaxOpenConns(25)                 // จำนวน Connection สูงสุดที่เปิดได้พร้อมกัน
	db.SetMaxIdleConns(25)                 // จำนวน Connection ที่สแตนด์บายรอใช้งาน
	db.SetConnMaxLifetime(5 * time.Minute) // อายุสูงสุดของ Connection ก่อนจะถูกรีเซ็ต

	// 3. ยิง Ping เพื่อทดสอบว่าต่อ Database ติดจริงหรือไม่
	if err := db.Ping(); err != nil {
		db.Close() // ปิด pool ทันทีถ้าเชื่อมต่อไม่สำเร็จ เพื่อคืน resource
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}