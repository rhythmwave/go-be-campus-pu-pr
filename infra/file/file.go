package file

import (
	"io"
	"time"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/utils"
)

type FileCtx struct {
	Firebase utils.StorageInterface
	Local    utils.StorageInterface
	Url      utils.StorageInterface
}

func (f FileCtx) GetURL(path, pathType string, expTime *time.Time) (string, *constants.ErrorResponse) {
	var result string
	var err error
	switch pathType {
	case constants.FirebaseStorageProvider:
		result, err = f.Firebase.GetURL(path, expTime)
		if err != nil {
			return result, constants.ErrorInternalServer(err.Error())
		}
	case constants.LocalStorageProvider:
		result, err = f.Local.GetURL(path, expTime)
		if err != nil {
			return result, constants.ErrorInternalServer(err.Error())
		}
	case constants.UrlStorageProvider:
		result, err = f.Url.GetURL(path, expTime)
		if err != nil {
			return result, constants.ErrorInternalServer(err.Error())
		}
	}

	return result, nil
}

func (f FileCtx) GetBytes(path, pathType string) ([]byte, *constants.ErrorResponse) {
	var result []byte
	var err error
	switch pathType {
	case constants.FirebaseStorageProvider:
		result, err = f.Firebase.GetBytes(path)
		if err != nil {
			return result, constants.ErrorInternalServer(err.Error())
		}
	case constants.LocalStorageProvider:
		result, err = f.Local.GetBytes(path)
		if err != nil {
			return result, constants.ErrorInternalServer(err.Error())
		}
	case constants.UrlStorageProvider:
		result, err = f.Url.GetBytes(path)
		if err != nil {
			return result, constants.ErrorInternalServer(err.Error())
		}
	}

	return result, nil
}

func (f FileCtx) PutFile(data io.Reader, path, pathType, mimeType string) error {
	switch pathType {
	case constants.FirebaseStorageProvider:
		err := f.Firebase.PutFile(data, path, mimeType)
		if err != nil {
			return err
		}
	case constants.LocalStorageProvider:
		err := f.Local.PutFile(data, path, mimeType)
		if err != nil {
			return err
		}
	}

	return nil
}
