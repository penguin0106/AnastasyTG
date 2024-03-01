package database

import (
	"database/sql"
	_ "github.com/lib/pq"
)

// DB представляет собой структуру базы данных
type DB struct {
	*sql.DB
}

// NewDB создает новое подключение к базе данных
func NewDB(connectionString string) (*DB, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

// Close закрывает соединение с базой данных
func (db *DB) Close() {
	db.DB.Close()
}

// CreateTables создает необходимые таблицы, если они не существуют
func (db *DB) CreateTables() error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(255) UNIQUE NOT NULL,
		password VARCHAR(255) NOT NULL,
		role VARCHAR(50) NOT NULL
	);`)
	if err != nil {
		return err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS products (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		price NUMERIC(10, 2) NOT NULL
	);`)
	if err != nil {
		return err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS orders (
		id SERIAL PRIMARY KEY,
		user_id INT NOT NULL,
		product_ids INT[] NOT NULL,
		status VARCHAR(50) NOT NULL,
		FOREIGN KEY (user_id) REFERENCES users(id)
	);`)
	if err != nil {
		return err
	}

	return nil
}
