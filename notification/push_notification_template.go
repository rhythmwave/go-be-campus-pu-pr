package notification

// PushNotificationTemplateData struct standard data template for push notification
type PushNotificationTemplateData struct {
	Title    string
	Body     string
	ImageURL string
}

//PushNotificationTemplate struct push notification template per SDK
type PushNotificationTemplate struct {
	Auth AuthPushNotificationTemplate
	Doc  DocPushNotificationTemplate
	Hris HrisPushNotificationTemplate
}

// AuthPushNotificationTemplate push notification template list for auth SDK
type AuthPushNotificationTemplate struct {
}

// DocPushNotificationTemplate push notification template list for doc SDK
type DocPushNotificationTemplate struct {
}

// AuthPushNotificationTemplate push notification template list for HRIS SDK
type HrisPushNotificationTemplate struct {
}
