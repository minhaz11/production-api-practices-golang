package sqlite

import (
	"database/sql"

	"github.com/minhaz11/crud/internal/config"
	_ "modernc.org/sqlite"
)

type Sqlite struct {
	DB *sql.DB
}

func New(cfg *config.Config) (*Sqlite, error) {
	db, err := sql.Open("sqlite", cfg.DbPath)

	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS students (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			email TEXT,
			age INTEGER
		)
	`)

	if err != nil {
		return nil, err
	}

	return &Sqlite{DB: db}, nil
}

func (s *Sqlite) CreateStudent(name string, email string, age int) (int64, error) {
	db := s.DB

	stmt, err := db.Prepare(`INSERT INTO students (name, email, age) VALUES (?, ?, ?)`)

	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(name, email, age)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return id, nil
}
