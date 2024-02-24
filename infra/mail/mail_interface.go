package mail

import (
	"github.com/sccicitb/pupr-backend/config"
	"github.com/sccicitb/pupr-backend/constants"
	objects "github.com/sccicitb/pupr-backend/objects/mail"
)

// MailInterface interface for mail service
type MailInterface interface {
	Send(detail objects.MailDetail) *constants.ErrorResponse
}

// NewSmtpMailer function to connect mail to MailInterface as smtp
// Params:
// cfg: mailer config
// Returns MailInterface
func NewSmtpMailer(cfg *config.Mailer) MailInterface {
	return &smtpMailer{
		cfg: cfg,
	}
}
