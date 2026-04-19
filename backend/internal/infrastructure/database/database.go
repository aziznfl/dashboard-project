package database

import (
	"database/sql"
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func InitDB(db *sql.DB) error {
	// create tables if not exists
	stmts := []string{
		`CREATE TABLE IF NOT EXISTS users (
		  id INTEGER PRIMARY KEY AUTOINCREMENT,
		  email TEXT NOT NULL UNIQUE,
		  password_hash TEXT NOT NULL,
		  role TEXT NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS payments (
		  id INTEGER PRIMARY KEY AUTOINCREMENT,
		  amount INTEGER NOT NULL,
		  merchant TEXT NOT NULL,
		  status TEXT NOT NULL,
		  created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);`,
	}
	for _, s := range stmts {
		if _, err := db.Exec(s); err != nil {
			return err
		}
	}

	if err := SeedUsers(db); err != nil {
		return err
	}

	if err := SeedPayments(db, 1000); err != nil {
		return err
	}

	return nil
}

func SeedUsers(db *sql.DB) error {
	var cnt int
	row := db.QueryRow("SELECT COUNT(1) FROM users")
	if err := row.Scan(&cnt); err != nil {
		return err
	}
	if cnt == 0 {
		hash, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		if _, err := db.Exec("INSERT INTO users(email, password_hash, role) VALUES (?, ?, ?)", "cs@test.com", string(hash), "cs"); err != nil {
			return err
		}
		if _, err := db.Exec("INSERT INTO users(email, password_hash, role) VALUES (?, ?, ?)", "operation@test.com", string(hash), "operation"); err != nil {
			return err
		}
	}
	return nil
}

func SeedPayments(db *sql.DB, count int) error {
	var payCnt int
	row := db.QueryRow("SELECT COUNT(1) FROM payments")
	if err := row.Scan(&payCnt); err != nil {
		return err
	}

	// Only seed if empty
	if payCnt > 0 {
		return nil
	}

	merchants := []string{"Gojek", "Tokopedia", "Grab", "Shopee", "Amazon", "Netflix", "Spotify", "Apple", "Google", "Facebook"}
	statuses := []string{"completed", "processing", "failed"}

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("INSERT INTO payments(amount, merchant, status, created_at) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < count; i++ {
		amount := int64(r.Intn(990000) + 10000) // 10,000 to 1,000,000
		merchant := merchants[r.Intn(len(merchants))]
		status := statuses[r.Intn(len(statuses))]

		// Random date in the last 30 days
		daysAgo := r.Intn(30)
		createdAt := time.Now().AddDate(0, 0, -daysAgo).Add(time.Duration(r.Intn(24)) * time.Hour)

		if _, err := stmt.Exec(amount, merchant, status, createdAt); err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}
