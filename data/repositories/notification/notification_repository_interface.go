package notification

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	models "github.com/sccicitb/pupr-backend/data/models/notification"
	"github.com/sccicitb/pupr-backend/infra/db"
)

type NotificationRepositoryInterface interface {
	BulkCreate(ctx context.Context, tx *sqlx.Tx, data []models.Create) *constants.ErrorResponse
}

func NewNotificationRepository(db *db.DB) NotificationRepositoryInterface {
	return &notificationRepository{
		db,
	}
}
