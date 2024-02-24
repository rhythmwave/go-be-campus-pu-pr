package utils

import (
	"io"
	"time"
)

// StorageInterface interface for AWS S3 and Firebase Storage
type StorageInterface interface {
	PutFile(data io.Reader, path, mimeType string) error
	DeleteFile(path string) error
	GetURL(path string, expTime *time.Time) (string, error)
	GetBytes(path string) ([]byte, error)
	GetProvider() string
}
