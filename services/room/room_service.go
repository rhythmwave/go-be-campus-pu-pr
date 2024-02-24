package room

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
	"github.com/sccicitb/pupr-backend/utils"
)

type roomService struct {
	*repository.RepoCtx
	*infra.InfraCtx
}

func (a roomService) GetList(ctx context.Context, paginationData common.PaginationRequest, req objects.GetRoomRequest) (objects.RoomListWithPagination, *constants.ErrorResponse) {
	var result objects.RoomListWithPagination

	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	modelResult, paginationResult, errs := a.RoomRepo.GetList(ctx, tx, paginationData, req)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	result = objects.RoomListWithPagination{
		Pagination: paginationResult,
		Data:       mapGetList(modelResult),
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, constants.ErrUnknown
	}

	return result, nil
}

func (f roomService) GetDetail(ctx context.Context, id string) (objects.GetRoomDetail, *constants.ErrorResponse) {
	var result objects.GetRoomDetail

	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	resultData, errs := f.RoomRepo.GetDetail(ctx, tx, id)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	result = objects.GetRoomDetail{
		Id:               resultData.Id,
		BuildingId:       resultData.BuildingId,
		BuildingCode:     resultData.BuildingCode,
		BuildingName:     resultData.BuildingName,
		Code:             resultData.Code,
		Name:             resultData.Name,
		Capacity:         resultData.Capacity,
		ExamCapacity:     resultData.ExamCapacity,
		Purpose:          resultData.Purpose,
		IsUsable:         resultData.IsUsable,
		Area:             resultData.Area,
		PhoneNumber:      resultData.PhoneNumber,
		Facility:         resultData.Facility,
		Remarks:          resultData.Remarks,
		Owner:            resultData.Owner,
		Location:         resultData.Location,
		StudyProgramId:   resultData.StudyProgramId,
		StudyProgramName: resultData.StudyProgramName,
		IsLaboratory:     resultData.IsLaboratory,
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, constants.ErrUnknown
	}

	return result, nil
}

func (a roomService) Create(ctx context.Context, data objects.CreateRoom) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	if data.StudyProgramId != "" {
		_, errs = a.StudyProgramRepo.GetDetail(ctx, tx, data.StudyProgramId, claims.Role, claims.ID)
		if errs != nil {
			_ = tx.Rollback()
			return errs
		}
	}

	createData := models.CreateRoom{
		BuildingId:     data.BuildingId,
		Code:           data.Code,
		Name:           utils.NewNullString(data.Name),
		Capacity:       utils.NewNullInt32(int32(data.Capacity)),
		ExamCapacity:   utils.NewNullInt32(int32(data.ExamCapacity)),
		IsUsable:       data.IsUsable,
		Area:           utils.NewNullFloat64(&data.Area),
		PhoneNumber:    utils.NewNullString(data.PhoneNumber),
		Facility:       utils.NewNullString(data.Facility),
		Remarks:        utils.NewNullString(data.Remarks),
		Purpose:        data.Purpose,
		Owner:          utils.NewNullString(data.Owner),
		Location:       utils.NewNullString(data.Location),
		StudyProgramId: utils.NewNullString(data.StudyProgramId),
		IsLaboratory:   data.IsLaboratory,
		CreatedBy:      claims.ID,
	}
	errs = a.RoomRepo.Create(ctx, tx, createData)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return constants.ErrUnknown
	}

	return nil
}

func (a roomService) Update(ctx context.Context, data objects.UpdateRoom) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs = a.RoomRepo.GetDetail(ctx, tx, data.Id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	if data.StudyProgramId != "" {
		_, errs = a.StudyProgramRepo.GetDetail(ctx, tx, data.StudyProgramId, claims.Role, claims.ID)
		if errs != nil {
			_ = tx.Rollback()
			return errs
		}
	}

	updateData := models.UpdateRoom{
		Id:             data.Id,
		Code:           data.Code,
		Name:           utils.NewNullString(data.Name),
		Capacity:       utils.NewNullInt32(int32(data.Capacity)),
		ExamCapacity:   utils.NewNullInt32(int32(data.ExamCapacity)),
		IsUsable:       data.IsUsable,
		Area:           utils.NewNullFloat64(&data.Area),
		PhoneNumber:    utils.NewNullString(data.PhoneNumber),
		Facility:       utils.NewNullString(data.Facility),
		Remarks:        utils.NewNullString(data.Remarks),
		Purpose:        data.Purpose,
		Owner:          utils.NewNullString(data.Owner),
		Location:       utils.NewNullString(data.Location),
		StudyProgramId: utils.NewNullString(data.StudyProgramId),
		UpdatedBy:      claims.ID,
	}
	errs = a.RoomRepo.Update(ctx, tx, updateData)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return constants.ErrUnknown
	}

	return nil
}

func (a roomService) Delete(ctx context.Context, id string) *constants.ErrorResponse {
	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs := a.RoomRepo.GetDetail(ctx, tx, id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	errs = a.RoomRepo.Delete(ctx, tx, id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return constants.ErrUnknown
	}

	return nil
}

func (a roomService) GetSchedule(ctx context.Context, paginationData common.PaginationRequest, roomId string, dayOfWeek uint32, semesterId string) (objects.RoomScheduleWithPagination, *constants.ErrorResponse) {
	var result objects.RoomScheduleWithPagination

	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	modelResult, paginationResult, errs := a.RoomRepo.GetSchedule(ctx, tx, paginationData, roomId, dayOfWeek, semesterId)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	roomIds := []string{}
	for _, v := range modelResult {
		roomIds = append(roomIds, v.RoomId)
	}

	scheduleData := []models.GetRoomScheduleDetail{}
	if len(roomIds) != 0 {
		scheduleData, errs = a.RoomRepo.GetScheduleByRoomIds(ctx, tx, roomIds, dayOfWeek, semesterId)
		if errs != nil {
			_ = tx.Rollback()
			return result, errs
		}

	}

	result = objects.RoomScheduleWithPagination{
		Pagination: paginationResult,
		Data:       mapGetSchedule(modelResult, scheduleData),
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, constants.ErrUnknown
	}

	return result, nil
}
