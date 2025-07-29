package data

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./clinic.db")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// doctors table
	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS doctors (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email TEXT UNIQUE NOT NULL,
			password TEXT NOT NULL,
			name TEXT NOT NULL,
			skill TEXT NOT NULL,
			title TEXT NOT NULL
		)
	`)
	if err != nil {
		log.Fatal("Failed to create doctors table:", err)
	}

	// patients table
	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS patients (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			doctor_id INTEGER,
			name TEXT NOT NULL,
			age INTEGER NOT NULL,
			diagnosis TEXT NOT NULL,
			contact TEXT NOT NULL,
			district TEXT NOT NULL,
			FOREIGN KEY (doctor_id) REFERENCES doctors(id)
		)
	`)
	if err != nil {
		log.Fatal("Failed to create patients table:", err)
	}

	// appointments table
	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS appointments (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			doctor_name TEXT NOT NULL,
			patient_name TEXT NOT NULL,
			diagnosis TEXT NOT NULL,
			time DATETIME DEFAULT CURRENT_TIMESTAMP,
			type TEXT NOT NULL,
			platform TEXT NOT NULL,
			FOREIGN KEY (doctor_name) REFERENCES doctors(name),
			FOREIGN KEY (patient_name) REFERENCES patients(name),
			FOREIGN KEY (diagnosis) REFERENCES patients(diagnosis)
		)
	`)
	if err != nil {
		log.Fatal("Faile to create appointments table:", err)
	}
}
