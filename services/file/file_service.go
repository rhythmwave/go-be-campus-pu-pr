package file

import (
	"context"
	"encoding/base64"
	"fmt"
	"mime"
	"os"
	"path/filepath"
	"strings"
	"time"

	mt "github.com/gabriel-vasile/mimetype"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	fileObject "github.com/sccicitb/pupr-backend/objects/file"
	"github.com/sccicitb/pupr-backend/utils"
)

// var (
// 	filename = fmt.Sprintf("temp/%s", utils.RandomString(40))
// )

type fileService struct {
	*infra.InfraCtx
}

func checkFolderTempExists() {
	if _, err := os.Stat("temp/"); os.IsNotExist(err) {
		fmt.Println("err", err)
		os.Mkdir("temp/", 0777)

	}
}

func (f fileService) UploadBlobTemp(ctx context.Context, data []byte, filePath string) (fileObject.FileSavedResponse, *constants.ErrorResponse) {
	filename := fmt.Sprintf("temp/%s-%s", utils.RandomString(40), time.Now().Format("200601021504050700"))
	checkFolderTempExists()
	var result fileObject.FileSavedResponse
	err := os.WriteFile(filename, data, 0644)
	if err != nil {
		return result, constants.ErrorInternalServer("error init file")
	}

	file, err := os.Open(filename)
	if err != nil {
		return result, constants.ErrorInternalServer("error open file")
	}
	defer file.Close()

	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		return result, constants.ErrorInternalServer("error read buffer")
	}

	var errs *constants.ErrorResponse
	data, errs = utils.ResizeImage(data, uint(f.Config.Server.MaxImageHeight))
	if errs != nil {
		return result, errs
	}

	path, mimeType, err := createFileFromBytes(data, filename)
	if err != nil {
		return result, constants.ErrorInternalServer("error create file")
	}

	result, errs = f.storeFile(ctx, path, filePath, filename, mimeType)
	if errs != nil {
		return result, errs
	}

	if err := os.Remove(path); err != nil {
		return result, constants.ErrorInternalServer("error open put file")
	}
	return result, nil
}

func (f fileService) UploadBase64Temp(ctx context.Context, data string, filePath string) (fileObject.FileSavedResponse, *constants.ErrorResponse) {
	filename := fmt.Sprintf("temp/%s-%s", utils.RandomString(40), time.Now().Format("200601021504050700"))
	checkFolderTempExists()
	var result fileObject.FileSavedResponse

	dec, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return result, constants.ErrorInternalServer("error decode file")
	}

	var errs *constants.ErrorResponse
	dec, errs = utils.ResizeImage(dec, uint(f.Config.Server.MaxImageHeight))
	if errs != nil {
		return result, errs
	}

	path, mimeType, err := createFileFromBytes(dec, filename)
	if err != nil {
		return result, constants.ErrorInternalServer("error create file", err.Error())
	}

	result, errs = f.storeFile(ctx, path, filePath, filename, mimeType)
	if errs != nil {
		return result, errs
	}
	if err := os.Remove(path); err != nil {
		return result, constants.ErrorInternalServer("error open put file")
	}
	return result, nil
}

func (f fileService) storeFile(ctx context.Context, path string, filePath string, fileName string, mimeType string) (fileObject.FileSavedResponse, *constants.ErrorResponse) {
	exp := time.Now().Add(constants.Oneday * 7)
	var result fileObject.FileSavedResponse
	var err error

	fileCreated, err := os.Open(path)
	if err != nil {
		return result, constants.ErrorInternalServer("error open created file")
	}

	savedPath := []string{}
	if filePath != "" {
		savedPath = append(savedPath, filePath)
	}

	savedPath = append(savedPath, filepath.Base(path))
	joinPath := strings.Join(savedPath, "/")
	if err := f.Storage.PutFile(fileCreated, joinPath, f.Config.Server.StorageProvider, mimeType); err != nil {
		fmt.Println(err)
		return result, constants.ErrorInternalServer("error open put file")
	}
	url, errs := f.Storage.GetURL(joinPath, f.Config.Server.StorageProvider, &exp)
	if errs != nil {
		return result, errs
	}
	return fileObject.FileSavedResponse{
		Path:     joinPath,
		PathType: f.Config.Server.StorageProvider,
		MimeType: mimeType,
		URL:      url,
		Size:     getFileSize(path),
	}, nil
}

func (f fileService) GetURL(ctx context.Context, path string) (string, *constants.ErrorResponse) {
	exp := time.Now().Add(constants.Oneday * 7)

	url, errs := f.Storage.GetURL(path, f.Config.Server.StorageProvider, &exp)
	if errs != nil {
		return "", errs
	}
	return url, nil
}

func createFileFromBytes(bytes []byte, filename string) (path string, mimetype string, err error) {
	err = os.WriteFile(filename, bytes, 0644)
	if err != nil {
		return "", "", err
	}

	file, err := os.Open(filename)
	if err != nil {
		return "", "", err
	}
	defer file.Close()

	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		return "", "", err
	}
	fullPath, err := filepath.Abs(filename)
	if err != nil {
		return "", "", err
	}
	// mimeType := http.DetectContentType(buffer)
	mtype, err := mt.DetectFile(fullPath)
	if err != nil {
		return "", "", err
	}
	out, err := mime.ExtensionsByType(mtype.String())
	if err != nil {
		return "", "", err
	}
	if len(out) > 0 {
		ext := out[0]
		if ext == ".moov" {
			ext = ".mov"
		}
		// fmt.Println("ext", ext)

		newPath := fmt.Sprintf("%s%s", fullPath, ext)
		os.Rename(fullPath, newPath)
		fullPath = newPath
	}

	return fullPath, mtype.String(), nil
}

func getFileSize(path string) int64 {
	fi, err := os.Stat(path)
	if err != nil {
		fmt.Println("ERROR GET SIZE", err)
		// Could not obtain stat, handle error
	}
	return fi.Size()
}
