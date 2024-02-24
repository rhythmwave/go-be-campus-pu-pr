package room

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type RoomServiceInterface interface {
	GetList(ctx context.Context, paginationData common.PaginationRequest, req objects.GetRoomRequest) (objects.RoomListWithPagination, *constants.ErrorResponse)
	GetDetail(ctx context.Context, id string) (objects.GetRoomDetail, *constants.ErrorResponse)
	Create(ctx context.Context, data objects.CreateRoom) *constants.ErrorResponse
	Update(ctx context.Context, data objects.UpdateRoom) *constants.ErrorResponse
	Delete(ctx context.Context, id string) *constants.ErrorResponse
	GetSchedule(ctx context.Context, paginationData common.PaginationRequest, roomId string, dayOfWeek uint32, semesterId string) (objects.RoomScheduleWithPagination, *constants.ErrorResponse)
}

func NewRoomService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) RoomServiceInterface {
	return &roomService{
		repoCtx,
		infraCtx,
	}
}
