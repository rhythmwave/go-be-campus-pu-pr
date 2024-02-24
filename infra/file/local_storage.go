package file

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/sccicitb/pupr-backend/config"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/utils"
)

func NewLocalStorage(cfg *config.Config) utils.StorageInterface {
	return &localStorage{
		cfg,
	}
}

type localStorage struct {
	*config.Config
}

func (u localStorage) checkFolderLocalExists() {
	if _, err := os.Stat(u.Server.LocalStoragePath); os.IsNotExist(err) {
		os.Mkdir(u.Server.LocalStoragePath, 0777)
	}
}

func (u localStorage) PutFile(data io.Reader, path, mimeType string) error {
	u.checkFolderLocalExists()

	fileName := fmt.Sprintf("%s/%s", u.Server.LocalStoragePath, path)

	buf := new(bytes.Buffer)
	buf.ReadFrom(data)

	err := os.WriteFile(fileName, buf.Bytes(), 0644)
	if err != nil {
		return err
	}

	return nil
}

func (u localStorage) DeleteFile(path string) error {
	return errors.New("Unimplemented")
}

func (u localStorage) GetURL(path string, expTime *time.Time) (string, error) {
	url := fmt.Sprintf("%s/%s/%s", u.Server.AppUrl, u.Server.LocalStoragePath, path)
	return url, nil
}

func (u localStorage) GetBytes(path string) ([]byte, error) {
	return []byte{}, errors.New("Unimplemented")
}

func (l localStorage) GetProvider() string {
	return constants.LocalStorageProvider
}
