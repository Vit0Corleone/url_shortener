package storage

import (
	"errors"
	"fmt"
	"os"
)

var (
	ErrUrlNotFound = errors.New("url not found")
	ErrUrlExists   = errors.New("url exists")
)

func GetDbConfig() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=%s password=%s",
		os.Getenv("STORAGE_HOST"),
		os.Getenv("STORAGE_PORT"),
		os.Getenv("STORAGE_USER"),
		os.Getenv("STORAGE_DBNAME"),
		os.Getenv("STORAGE_SSLMODE"),
		os.Getenv("STORAGE_PASSWORD"),
	)
}
