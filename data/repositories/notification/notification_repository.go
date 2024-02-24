package notification

import (
	"context"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	models "github.com/sccicitb/pupr-backend/data/models/notification"
	"github.com/sccicitb/pupr-backend/infra/db"

	"google.golang.org/grpc/codes"
)

type notificationRepository struct {
	*db.DB
}

func (n notificationRepository) BulkCreate(ctx context.Context, tx *sqlx.Tx, data []models.Create) *constants.ErrorResponse {
	_, err := tx.NamedExecContext(
		ctx,
		bulkCreateQuery,
		data,
	)
	if err != nil {
		return constants.Error(http.StatusInternalServerError, codes.Internal, constants.DefaultCustomErrorCode, err.Error())
	}

	return nil
}
