package lecturer_leave

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	appConstants "github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
	"github.com/sccicitb/pupr-backend/utils"
	appUtils "github.com/sccicitb/pupr-backend/utils"
)

type lecturerLeaveService struct {
	*repository.RepoCtx
	*infra.InfraCtx
}

func (l lecturerLeaveService) GetList(ctx context.Context, paginationData common.PaginationRequest, studyProgramId, idNationalLecturer, semesterId string, isActive bool) (objects.LecturerLeaveListWithPagination, *constants.ErrorResponse) {
	var result objects.LecturerLeaveListWithPagination

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return result, errs
	}

	tx, err := l.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	if studyProgramId != "" {
		_, errs := l.StudyProgramRepo.GetDetail(ctx, tx, studyProgramId, claims.Role, claims.ID)
		if errs != nil {
			_ = tx.Rollback()
			return result, errs
		}
	}

	modelResult, paginationResult, errs := l.LecturerLeaveRepo.GetList(ctx, tx, paginationData, studyProgramId, idNationalLecturer, semesterId, isActive)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	resultData := []objects.GetLecturerLeave{}
	for _, v := range modelResult {
		var fileUrl string
		if v.FilePath != nil && v.FilePathType != nil {
			fileUrl, errs = l.Storage.GetURL(*v.FilePath, *v.FilePathType, nil)
			if errs != nil {
				_ = tx.Rollback()
				return result, errs
			}
		}

		resultData = append(resultData, objects.GetLecturerLeave{
			Id:                    v.Id,
			Name:                  v.Name,
			IdNationalLecturer:    v.IdNationalLecturer,
			FrontTitle:            v.FrontTitle,
			BackDegree:            v.BackDegree,
			SemesterSchoolYear:    appUtils.GenerateSchoolYear(v.SemesterStartYear),
			SemesterType:          v.SemesterType,
			DiktiStudyProgramCode: v.DiktiStudyProgramCode,
			StudyProgramName:      v.StudyProgramName,
			StudyLevelShortName:   v.StudyLevelShortName,
			DiktiStudyProgramType: v.DiktiStudyProgramType,
			StartDate:             v.StartDate,
			EndDate:               v.EndDate,
			PermitNumber:          v.PermitNumber,
			Purpose:               v.Purpose,
			Remarks:               v.Remarks,
			FileUrl:               fileUrl,
			FilePath:              v.FilePath,
			FilePathType:          v.FilePathType,
		})
	}

	result = objects.LecturerLeaveListWithPagination{
		Pagination: paginationResult,
		Data:       resultData,
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, constants.ErrUnknown
	}

	return result, nil
}

func (l lecturerLeaveService) Create(ctx context.Context, data objects.CreateLecturerLeave) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := l.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	existLeave, errs := l.LecturerLeaveRepo.GetDetailByLecturerIdAndStartDate(ctx, tx, data.LecturerId, data.StartDate)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}
	if existLeave.Id != "" {
		_ = tx.Rollback()
		return appConstants.ErrIdenticalLeave
	}

	createData := models.CreateLecturerLeave{
		LecturerId:   data.LecturerId,
		SemesterId:   data.SemesterId,
		StartDate:    data.StartDate,
		PermitNumber: data.PermitNumber,
		Purpose:      data.Purpose,
		Remarks:      data.Remarks,
		FilePath:     utils.NewNullString(data.FilePath),
		FilePathType: utils.NewNullString(data.FilePathType),
		CreatedBy:    claims.ID,
	}
	errs = l.LecturerLeaveRepo.Create(ctx, tx, createData)
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

func (l lecturerLeaveService) Update(ctx context.Context, data objects.UpdateLecturerLeave) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := l.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	leaveData, errs := l.LecturerLeaveRepo.GetDetail(ctx, tx, data.Id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	if leaveData.IsActive {
		_ = tx.Rollback()
		return appConstants.ErrLeaveActiveUneditable
	}

	updateData := models.UpdateLecturerLeave{
		Id:           data.Id,
		PermitNumber: data.PermitNumber,
		Purpose:      data.Purpose,
		Remarks:      data.Remarks,
		FilePath:     utils.NewNullString(data.FilePath),
		FilePathType: utils.NewNullString(data.FilePathType),
		UpdatedBy:    claims.ID,
	}
	errs = l.LecturerLeaveRepo.Update(ctx, tx, updateData)
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

func (l lecturerLeaveService) End(ctx context.Context, id string) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := l.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	leaveData, errs := l.LecturerLeaveRepo.GetDetail(ctx, tx, id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	if !leaveData.IsActive {
		_ = tx.Rollback()
		return appConstants.ErrLeaveIsNotActive
	}

	errs = l.LecturerLeaveRepo.End(ctx, tx, id, claims.ID)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	errs = l.LecturerRepo.UpdateStatus(ctx, tx, []string{id}, appConstants.LecturerStatusAktif)
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

func (l lecturerLeaveService) Delete(ctx context.Context, id string) *constants.ErrorResponse {
	tx, err := l.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	leaveData, errs := l.LecturerLeaveRepo.GetDetail(ctx, tx, id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	if leaveData.IsActive {
		_ = tx.Rollback()
		return appConstants.ErrLeaveActiveUneditable
	}

	errs = l.LecturerLeaveRepo.Delete(ctx, tx, id)
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
