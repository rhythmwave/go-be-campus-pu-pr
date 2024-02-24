package utils

import (
	"github.com/sccicitb/pupr-backend/constants"
	objects "github.com/sccicitb/pupr-backend/objects/notification"
)

// NotificationInterface interface for fcm infrastructure
type NotificationInterface interface {
	GetProviderName() string
	SendNotification(data objects.SendNotification) *constants.ErrorResponse
}
