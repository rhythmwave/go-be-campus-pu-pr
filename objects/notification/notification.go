package notification

import "time"

type SendNotificationData struct {
	Body             string
	Title            string
	AndroidChannelID string
	Priority         string
	ClickAction      string
	Type             string
	ID               string
	Status           string
	RefURL           string
	PayloadData      map[string]interface{}
}

// SendNotification send notification request data
type SendNotification struct {
	Title        string
	Body         string
	TargetTokens []string
	Data         SendNotificationData
	SaveToTable  bool
	UserIDs      []string
	ExpiryTime   time.Time
}
