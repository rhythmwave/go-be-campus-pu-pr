package notification

// MailTemplateData struct standard data template for mail
type MailTemplateData struct {
	Subject  string
	Template string
}

// MailTemplate struct mail template per SDK
type MailTemplate struct {
	LogoURL string
	Auth    AuthMailTemplate
	Doc     DocMailTemplate
	Hris    HrisMailTemplate
}

// AuthMailTemplate mail template list for auth SDK
type AuthMailTemplate struct {
	// template need object OTP(string) and Year(int)
	LoginViaOTP MailTemplateData
	// template need object OTP(string) and Year(int)
	Registration MailTemplateData
	// template need object OTP(string) and Year(int)
	ForgotPassword MailTemplateData
	// template need object OTP(string) and Year(int)
	AdminForgotPassword MailTemplateData
}

// DocMailTemplate mail template list for doc SDK
type DocMailTemplate struct {
}

// AuthMailTemplate mail template list for HRIS SDK
type HrisMailTemplate struct {
	RequestEmployeeInvitation MailTemplateData
	EmployeeInvitation        MailTemplateData
	ErrorInvitation           MailTemplateData
}
