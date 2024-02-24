package mail

// MailAttachment struct of mail attachmets
type MailAttachment struct {
	FileDir  string
	FileName string
}

// MailDetail struct for mail detail
type MailDetail struct {
	LogoURL     string
	To          []string
	Cc          []string
	Subject     string
	Html        string
	Template    string
	Attachments []MailAttachment
	Data        interface{}
}
