package mail

import (
	"fmt"
	"net/http"
	"os"

	"github.com/sccicitb/pupr-backend/config"
	"github.com/sccicitb/pupr-backend/constants"
	objects "github.com/sccicitb/pupr-backend/objects/mail"
	"github.com/sccicitb/pupr-backend/utils"
	"google.golang.org/grpc/codes"

	log "github.com/sirupsen/logrus"

	gomail "gopkg.in/gomail.v2"
)

type smtpMailer struct {
	cfg *config.Mailer
}

// Send function to send mail
// Params:
// detail: mail detail
// file: attachment file directory
// Returns *constants.ErrorResponse
func (m *smtpMailer) Send(detail objects.MailDetail) *constants.ErrorResponse {
	result, errs := utils.GenerateMailBody(detail)
	if errs != nil {
		return errs
	}

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", m.cfg.Sender)
	mailer.SetHeader("To", detail.To...)
	for _, cc := range detail.Cc {
		mailer.SetAddressHeader("Cc", cc, cc)
	}
	mailer.SetHeader("Subject", detail.Subject)
	mailer.SetBody("text/html", result)

	dialer := gomail.NewDialer(
		m.cfg.Server,
		int(m.cfg.Port),
		m.cfg.Username,
		m.cfg.Password,
	)

	files := []string{}
	for _, v := range detail.Attachments {
		files = append(files, v.FileDir)

	}

	for _, v := range files {
		mailer.Attach(fmt.Sprintf("%s%s", constants.TempDownloadedFileDir, v))

	}

	if err := dialer.DialAndSend(mailer); err != nil {
		return constants.Error(http.StatusInternalServerError, codes.Internal, constants.DefaultCustomErrorCode, err.Error())
	}

	for _, v := range files {
		err := os.Remove(fmt.Sprintf("%s%s", constants.TempDownloadedFileDir, v))
		if err != nil {
			log.Error(err)
		}
	}
	return nil
}
