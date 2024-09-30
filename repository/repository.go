package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type UserRepository struct {
	DB *sql.DB
}

func New(path string) (*UserRepository, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	stmt, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS feedback(
	    client_id INTEGER PRIMARY KEY,
	    rating INT NOT NULL CHECK(rating BETWEEN 1 AND 5),
		comment TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP);
	`)
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec()
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	return &UserRepository{DB: db}, nil
}

func (r *UserRepository) FindById(id int) (bool, error) {
	var exists bool
	query := "SELECT EXISTS (SELECT * FROM feedback WHERE client_id = ?)"
	err := r.DB.QueryRow(query, id).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}
