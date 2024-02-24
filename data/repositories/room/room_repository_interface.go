package room

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type RoomRepositoryInterface interface {
	GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, req objects.GetRoomRequest) ([]models.GetRoom, common.Pagination, *constants.ErrorResponse)
	GetDetail(ctx context.Context, tx *sqlx.Tx, id string) (models.GetRoomDetail, *constants.ErrorResponse)
	GetDetailByRoomIds(ctx context.Context, tx *sqlx.Tx, ids []string) ([]models.GetRoomDetail, *constants.ErrorResponse)
	Create(ctx context.Context, tx *sqlx.Tx, data models.CreateRoom) *constants.ErrorResponse
	Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateRoom) *constants.ErrorResponse
	Delete(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse
	GetSchedule(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, roomId string, dayOfWeek uint32, semesterId string) ([]models.GetRoomSchedule, common.Pagination, *constants.ErrorResponse)
	GetScheduleByRoomIds(ctx context.Context, tx *sqlx.Tx, roomIds []string, dayOfWeek uint32, semesterId string) ([]models.GetRoomScheduleDetail, *constants.ErrorResponse)
}

func NewRoomRepository(db *db.DB) RoomRepositoryInterface {
	return &roomRepository{
		db,
	}
}
