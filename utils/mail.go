package utils

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/sccicitb/pupr-backend/constants"
	objects "github.com/sccicitb/pupr-backend/objects/mail"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
)

// ParseAttachmentsFromMultipartForm function to get mail attachment data from multipart/form-data
// Params:
// files: request file
// Returns list mail attachment and *constants.ErrorResponse
func ParseAttachmentsFromMultipartForm(files []*multipart.FileHeader) ([]objects.MailAttachment, *constants.ErrorResponse) {
	results := []objects.MailAttachment{}

	for i := range files {
		file, err := files[i].Open()
		if err != nil {
			return results, constants.Error(http.StatusInternalServerError, codes.Internal, constants.DefaultCustomErrorCode, err.Error())
		}
		defer func() {
			err := file.Close()
			if err != nil {
				log.Errorln(err)
			}
		}()

		out, err := os.Create(fmt.Sprintf("%s%s", constants.TempDownloadedFileDir, files[i].Filename))
		if err != nil {
			return results, constants.Error(http.StatusInternalServerError, codes.Internal, constants.DefaultCustomErrorCode, err.Error())
		}
		defer func() {
			err := out.Close()
			if err != nil {
				log.Errorln(err)
			}
		}()

		_, err = io.Copy(out, file)
		if err != nil {
			return results, constants.Error(http.StatusInternalServerError, codes.Internal, constants.DefaultCustomErrorCode, err.Error())
		}
		results = append(results, objects.MailAttachment{
			FileDir:  fmt.Sprintf("%s%s", constants.TempDownloadedFileDir, files[i].Filename),
			FileName: files[i].Filename,
		})
	}

	return results, nil
}

// GenerateMailBody function to generate mail body from mail detail
// Params:
// detail: mail detail
// Returns string and *constants.ErrorResponse
func GenerateMailBody(detail objects.MailDetail) (string, *constants.ErrorResponse) {
	mainTemplate := constants.MailMainTemplate
	logoTemplate := constants.MailLogoTemplate
	footerTemplate := constants.MailFooterTemplate
	defaultBodyTemplate := constants.DefaultBodyTemplate

	var result string

	if detail.Template != "" && detail.Html != "" {
		return result, constants.ErrEmailTemplate
	}
	if detail.Template != "" {
		templ := template.New("").Funcs(template.FuncMap{
			"CurrencyFormat": CommaSeparated,
		})
		t := template.Must(templ.ParseFiles(mainTemplate, logoTemplate, footerTemplate, detail.Template))
		var tpl bytes.Buffer
		if err := t.ExecuteTemplate(&tpl, "layout", detail.Data); err != nil {
			return result, constants.Error(http.StatusInternalServerError, codes.Internal, constants.DefaultCustomErrorCode, err.Error())
		}

		result = tpl.String()
	} else if detail.Html != "" {
		detail.Data = struct {
			Year int
		}{
			Year: time.Now().Year(),
		}

		t := template.Must(template.ParseFiles(mainTemplate, logoTemplate, footerTemplate, defaultBodyTemplate))
		var tpl bytes.Buffer
		if err := t.ExecuteTemplate(&tpl, "layout", detail.Data); err != nil {
			return result, constants.Error(http.StatusInternalServerError, codes.Internal, constants.DefaultCustomErrorCode, err.Error())
		}

		result = tpl.String()

		result = strings.ReplaceAll(result, "{BODY_TEMPLATE}", detail.Html)
	}
	result = strings.ReplaceAll(result, "{LOGO_URL}", detail.LogoURL)

	return result, nil
}
