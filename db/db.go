package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	var err error

	DB, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("❌ Ошибка подключения к базе:", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatal("❌ Ошибка пинга базы:", err)
	}

	log.Println("✅ Подключение к БД установлено")

	// Автоматическое создание таблиц
	createTables := []string{
		`CREATE TABLE IF NOT EXISTS depos (
			id SERIAL PRIMARY KEY,
			name VARCHAR(100) NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			username VARCHAR(50) NOT NULL UNIQUE,
			password TEXT NOT NULL,
			full_name TEXT,
			position TEXT,
			depo_id INT REFERENCES depos(id),
			tabel_num VARCHAR(11),
			phone VARCHAR(20),
			role VARCHAR(20) DEFAULT 'user',
			is_active BOOLEAN DEFAULT TRUE
		);`,
		`CREATE TABLE IF NOT EXISTS tabels (
			id SERIAL PRIMARY KEY,
			full_name VARCHAR(100) NOT NULL,
			tabel_num VARCHAR(10) UNIQUE NOT NULL,
			position VARCHAR(100),
			phone VARCHAR(20),
			depo_id INT REFERENCES depos(id)
		);`,
		`CREATE TABLE IF NOT EXISTS locomotives (
			id SERIAL PRIMARY KEY,
			model VARCHAR(50),
			number VARCHAR(50),
			depo_id INT REFERENCES depos(id),
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);`,
		`CREATE TABLE IF NOT EXISTS dizel_oil_teplovoz (
			id SERIAL PRIMARY KEY,
			depo_id INT REFERENCES depos(id),
			analysis_date DATE,
			repair_type VARCHAR(20),
			locomotive VARCHAR(100),
			section VARCHAR(10),
			flash_point FLOAT,
			viscosity FLOAT,
			contamination FLOAT,
			water_content FLOAT,
			comment TEXT,
			employee_number INT,
			last_oil_date DATE,
			conclusion VARCHAR(20),
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);`,
	}

	for _, tableSQL := range createTables {
		if _, err := DB.Exec(tableSQL); err != nil {
			log.Fatal("❌ Ошибка создания таблицы:", err)
		}
	}
	log.Println("✅ Все таблицы проверены/созданы")
}
