package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"url_short/internal/storage"

	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

func New() (*Storage, error) {
	const op = "storage.postgres.New"

	db, err := sql.Open("postgres", storage.GetDbConfig())
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	stmt, err := db.Prepare(
		`CREATE TABLE IF NOT EXISTS url(
			id SERIAL NOT NULL UNIQUE,
			alias TEXT NOT NULL UNIQUE,
			url TEXT NOT NULL UNIQUE);
		`,
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if _, err = stmt.Exec(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) SaveURL(urlToSave, alias string) error {
	const op = "storage.postgres.SaveURL"

	_, err := s.db.Exec("INSERT INTO url(alias, url) VALUES($1, $2)", alias, urlToSave)
	if err != nil {
		return fmt.Errorf("%s: %w", op, storage.ErrUrlExists)
	}

	return nil
}

func (s *Storage) GetUrl(alias string) (string, error) {
	var op = "storage.postgres.GetUrl"
	var newUrl string

	if err := s.db.QueryRow("SELECT url FROM url WHERE alias=$1", alias).Scan(&newUrl); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", storage.ErrUrlNotFound
		}
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return newUrl, nil
}

func (s *Storage) Delete(alias string) error {
	var op = "storage.postgres.Delete"

	if _, err := s.db.Exec("DELETE FROM url WHERE alias=$1", alias); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
