package student_leave

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

type studentLeaveService struct {
	*repository.RepoCtx
	*infra.InfraCtx
}

func (l studentLeaveService) GetListRequest(ctx context.Context, paginationData common.PaginationRequest, appType, studyProgramId string, isApproved *bool) (objects.StudentLeaveRequestListWithPagination, *constants.ErrorResponse) {
	var result objects.StudentLeaveRequestListWithPagination

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return result, errs
	}

	tx, err := l.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	if studyProgramId != "" && appType == appConstants.AppTypeAdmin {
		_, errs := l.StudyProgramRepo.GetDetail(ctx, tx, studyProgramId, claims.Role, claims.ID)
		if errs != nil {
			_ = tx.Rollback()
			return result, errs
		}
	}

	var studentId string
	if appType == appConstants.AppTypeStudent {
		studentId = claims.ID
	}

	modelResult, paginationResult, errs := l.StudentLeaveRepo.GetListRequest(ctx, tx, paginationData, studyProgramId, studentId, isApproved)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	resultData := []objects.GetStudentLeaveRequest{}
	for _, v := range modelResult {
		resultData = append(resultData, objects.GetStudentLeaveRequest{
			Id:                         v.Id,
			NimNumber:                  v.NimNumber,
			Name:                       v.Name,
			DiktiStudyProgramCode:      v.DiktiStudyProgramCode,
			StudyProgramName:           v.StudyProgramName,
			StudyLevelShortName:        v.StudyLevelShortName,
			DiktiStudyProgramType:      v.DiktiStudyProgramType,
			StartDate:                  v.StartDate,
			TotalLeaveDurationSemester: v.TotalLeaveDurationSemester,
			PermitNumber:               v.PermitNumber,
			Purpose:                    v.Purpose,
			Remarks:                    v.Remarks,
			IsApproved:                 v.IsApproved,
			SemesterType:               v.SemesterType,
			SemesterSchoolYear:         appUtils.GenerateSchoolYear(v.SemesterStartYear),
		})
	}

	result = objects.StudentLeaveRequestListWithPagination{
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

func (l studentLeaveService) GetList(ctx context.Context, paginationData common.PaginationRequest, studyProgramId, semesterId string) (objects.StudentLeaveListWithPagination, *constants.ErrorResponse) {
	var result objects.StudentLeaveListWithPagination

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

	modelResult, paginationResult, errs := l.StudentLeaveRepo.GetList(ctx, tx, paginationData, studyProgramId, semesterId)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	resultData := []objects.GetStudentLeave{}
	for _, v := range modelResult {
		resultData = append(resultData, objects.GetStudentLeave{
			Id:                    v.Id,
			NimNumber:             v.NimNumber,
			Name:                  v.Name,
			DiktiStudyProgramCode: v.DiktiStudyProgramCode,
			StudyProgramName:      v.StudyProgramName,
			StudyLevelShortName:   v.StudyLevelShortName,
			DiktiStudyProgramType: v.DiktiStudyProgramType,
			SemesterSchoolYear:    appUtils.GenerateSchoolYear(v.SemesterStartYear),
			SemesterType:          v.SemesterType,
			PermitNumber:          v.PermitNumber,
			Purpose:               v.Purpose,
			Remarks:               v.Remarks,
		})
	}

	result = objects.StudentLeaveListWithPagination{
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

func (l studentLeaveService) Create(ctx context.Context, data objects.CreateStudentLeave) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := l.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	studentId := data.StudentId
	if studentId == "" {
		studentId = claims.ID
	}

	createData := models.CreateStudentLeave{
		StudentId:                  studentId,
		TotalLeaveDurationSemester: data.TotalLeaveDurationSemester,
		StartDate:                  data.StartDate,
		PermitNumber:               utils.NewNullString(data.PermitNumber),
		Purpose:                    data.Purpose,
		Remarks:                    data.Remarks,
	}
	errs = l.StudentLeaveRepo.Create(ctx, tx, createData)
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

func (l studentLeaveService) Update(ctx context.Context, data objects.UpdateStudentLeave) *constants.ErrorResponse {
	tx, err := l.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	leaveData, errs := l.StudentLeaveRepo.GetDetailRequest(ctx, tx, data.Id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	if utils.NullBooleanScan(leaveData.IsApproved) {
		_ = tx.Rollback()
		return appConstants.ErrLeaveActiveUneditable
	}

	updateData := models.UpdateStudentLeave{
		Id:           data.Id,
		PermitNumber: utils.NewNullString(data.PermitNumber),
		Purpose:      data.Purpose,
		Remarks:      data.Remarks,
	}
	errs = l.StudentLeaveRepo.Update(ctx, tx, updateData)
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

func (l studentLeaveService) Approve(ctx context.Context, id string, isApproved bool) *constants.ErrorResponse {
	tx, err := l.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	leaveData, errs := l.StudentLeaveRepo.GetDetailRequest(ctx, tx, id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	if leaveData.IsApproved != nil {
		_ = tx.Rollback()
		return appConstants.ErrLeaveActiveUneditable
	}

	errs = l.StudentLeaveRepo.Approve(ctx, tx, id, isApproved)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	errs = l.StudentRepo.BulkUpdateStatus(ctx, tx, models.BulkUpdateStatusStudent{
		Ids:    []string{id},
		Status: appConstants.StudentStatusAktif,
	})
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

func (l studentLeaveService) End(ctx context.Context, id string) *constants.ErrorResponse {
	tx, err := l.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	leaveData, errs := l.StudentLeaveRepo.GetDetailRequest(ctx, tx, id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	if !utils.NullBooleanScan(leaveData.IsApproved) {
		_ = tx.Rollback()
		return appConstants.ErrLeaveIsNotActive
	}

	errs = l.StudentLeaveRepo.End(ctx, tx, id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	errs = l.StudentRepo.BulkUpdateStatus(ctx, tx, models.BulkUpdateStatusStudent{
		Ids:    []string{id},
		Status: appConstants.StudentStatusAktif,
	})
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

func (l studentLeaveService) Delete(ctx context.Context, id string) *constants.ErrorResponse {
	tx, err := l.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	leaveData, errs := l.StudentLeaveRepo.GetDetailRequest(ctx, tx, id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	if utils.NullBooleanScan(leaveData.IsApproved) {
		_ = tx.Rollback()
		return appConstants.ErrLeaveActiveUneditable
	}

	errs = l.StudentLeaveRepo.Delete(ctx, tx, id)
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
