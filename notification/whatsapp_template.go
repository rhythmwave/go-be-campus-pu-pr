package notification

// WhatsappTemplateData struct standard data template for whatsapp
type WhatsappTemplateData struct {
	Message  string
	MediaURL string
	Link     string
}

// WhatsappTemplate struct whatsapp template per SDK
type WhatsappTemplate struct {
	Auth AuthWhatsappTemplate
	Doc  DocWhatsappTemplate
	Hris HrisWhatsappTemplate
}

// AuthWhatsappTemplate whatsapp template list for auth SDK
type AuthWhatsappTemplate struct {
	// message should provide one string data for OTP
	AccountActivation WhatsappTemplateData
	// message should provide one string data for OTP
	ForgotPassword WhatsappTemplateData
}

// DocWhatsappTemplate whatsapp template list for doc SDK
type DocWhatsappTemplate struct {
	SignatureRequest   WhatsappTemplateData
	SignatureCompleted WhatsappTemplateData
}

// AuthWhatsappTemplate whatsapp template list for HRIS SDK
type HrisWhatsappTemplate struct {
	RequestToInvite    WhatsappTemplateData
	EmployeeInvitation WhatsappTemplateData
}
