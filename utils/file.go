package utils

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"mime"
	"net/http"
	"os"
	"path/filepath"

	"github.com/nfnt/resize"
	"github.com/sccicitb/pupr-backend/constants"
)

// ConvertFileToBase64 function to convert file to it's base64
// Params:
// file: file directory
// Returns file base64 string
func ConvertFileToBase64(file string) string {
	f, _ := os.Open(file)

	// Read entire JPG into byte slice.
	reader := bufio.NewReader(f)
	content, _ := io.ReadAll(reader)

	// Encode as base64.
	encoded := base64.StdEncoding.EncodeToString(content)
	return encoded
}

func CreateFileFromBytes(bytes []byte, filename string) (path string, mimetype string, err error) {
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
	mimeType := http.DetectContentType(buffer)
	out, err := mime.ExtensionsByType(mimeType)
	if err != nil {
		return "", "", err
	}
	if len(out) > 0 {
		newPath := fmt.Sprintf("%s%s", fullPath, out[0])
		os.Rename(fullPath, newPath)
		fullPath = newPath
	}

	return fullPath, mimeType, nil
}

func GetFileSize(path string) int64 {
	fi, err := os.Stat(path)
	if err != nil {
		fmt.Println("ERROR GET SIZE", err)
		// Could not obtain stat, handle error
	}
	return fi.Size()
}

func ResizeImage(data []byte, maxHeight uint) ([]byte, *constants.ErrorResponse) {
	contentType := http.DetectContentType(data)
	if contentType != constants.JpegContentType {
		return data, nil
	}

	if maxHeight == 0 {
		maxHeight = constants.DefaultMaxImageHeight
	}

	var result []byte
	image, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return result, constants.ErrorInternalServer(err.Error())
	}

	newImage := resize.Resize(0, maxHeight, image, resize.Lanczos3)

	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	err = jpeg.Encode(w, newImage, nil)
	if err != nil {
		return result, constants.ErrorInternalServer(err.Error())
	}
	result = b.Bytes()

	return result, nil
}
