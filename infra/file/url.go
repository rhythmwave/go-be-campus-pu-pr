package file

import (
	"errors"
	"io"
	"time"

	"github.com/sccicitb/pupr-backend/config"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/utils"
)

func NewUrlStorage(cfg *config.Config) utils.StorageInterface {
	return &urlStorage{
		cfg,
	}
}

type urlStorage struct {
	*config.Config
}

func (u urlStorage) PutFile(data io.Reader, path, mimeType string) error {
	return errors.New("Unimplemented")
}

func (u urlStorage) DeleteFile(path string) error {
	return errors.New("Unimplemented")
}

func (u urlStorage) GetURL(path string, expTime *time.Time) (string, error) {
	return path, nil
}

func (u urlStorage) GetBytes(path string) ([]byte, error) {
	return []byte{}, errors.New("Unimplemented")
}

func (u urlStorage) GetProvider() string {
	return constants.UrlStorageProvider
}
