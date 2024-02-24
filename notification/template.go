package notification

// NotificationTemplate struct notification template
type NotificationTemplate struct {
	Mail             MailTemplate
	PushNotification PushNotificationTemplate
	Whatsapp         WhatsappTemplate
}
