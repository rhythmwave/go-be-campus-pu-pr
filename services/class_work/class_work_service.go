package class_work

import (
	"context"
	"time"

	"github.com/sccicitb/pupr-backend/constants"
	appConstants "github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
	"github.com/sccicitb/pupr-backend/utils"
)

type classWorkService struct {
	*repository.RepoCtx
	*infra.InfraCtx
}

func (c classWorkService) GetList(ctx context.Context, paginationData common.PaginationRequest, appType string, classIds []string) (objects.ClassWorkListWithPagination, *constants.ErrorResponse) {
	var result objects.ClassWorkListWithPagination

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return result, errs
	}

	tx, err := c.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	var studentId string
	if appType == appConstants.AppTypeStudent {
		activeSemesterData, errs := c.SemesterRepo.GetActive(ctx, tx)
		if errs != nil {
			_ = tx.Rollback()
			return result, errs
		}

		studentId = claims.ID
		takenClassData, _, errs := c.StudentClassRepo.GetList(ctx, tx, common.PaginationRequest{Limit: constants.DefaultUnlimited}, "", studentId, activeSemesterData.Id, nil)
		if errs != nil {
			_ = tx.Rollback()
			return result, errs
		}

		for _, v := range takenClassData {
			classIds = append(classIds, v.ClassId)
		}
	}

	modelResult, paginationResult, errs := c.ClassWorkRepo.GetList(ctx, tx, paginationData, classIds, studentId)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	resultData := []objects.GetClassWork{}
	for _, v := range modelResult {
		var fileUrl string
		if v.FilePath != nil && v.FilePathType != nil {
			fileUrl, errs = c.Storage.GetURL(*v.FilePath, *v.FilePathType, nil)
			if errs != nil {
				_ = tx.Rollback()
				return result, errs
			}
		}

		var submissionFileUrl string
		if v.SubmissionFilePath != nil && v.SubmissionFilePathType != nil {
			submissionFileUrl, errs = c.Storage.GetURL(*v.SubmissionFilePath, *v.SubmissionFilePathType, nil)
			if errs != nil {
				_ = tx.Rollback()
				return result, errs
			}
		}

		resultData = append(resultData, objects.GetClassWork{
			Id:                     v.Id,
			Title:                  v.Title,
			Abstraction:            v.Abstraction,
			FileUrl:                fileUrl,
			LecturerId:             v.LecturerId,
			LecturerName:           v.LecturerName,
			LecturerFrontTitle:     v.LecturerFrontTitle,
			LecturerBackDegree:     v.LecturerBackDegree,
			StartTime:              v.StartTime,
			EndTime:                v.EndTime,
			TotalSubmission:        v.TotalSubmission,
			SubmissionFileUrl:      submissionFileUrl,
			SubmissionFilePath:     v.SubmissionFilePath,
			SubmissionFilePathType: v.SubmissionFilePathType,
			SubmissionPoint:        v.SubmissionPoint,
		})
	}

	result = objects.ClassWorkListWithPagination{
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

func (c classWorkService) GetSubmission(ctx context.Context, paginationData common.PaginationRequest, classWorkId string) (objects.ClassWorkSubmissionWithPagination, *constants.ErrorResponse) {
	var result objects.ClassWorkSubmissionWithPagination

	tx, err := c.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	classWorkData, errs := c.ClassWorkRepo.GetDetail(ctx, tx, classWorkId)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	modelResult, paginationResult, errs := c.ClassWorkRepo.GetSubmission(ctx, tx, paginationData, classWorkData.ClassId, classWorkId, nil, nil)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	resultData := []objects.GetClassWorkSubmission{}
	for _, v := range modelResult {
		var fileUrl string
		if v.FilePath != nil && v.FilePathType != nil {
			fileUrl, errs = c.Storage.GetURL(*v.FilePath, *v.FilePathType, nil)
			if errs != nil {
				_ = tx.Rollback()
				return result, errs
			}
		}

		resultData = append(resultData, objects.GetClassWorkSubmission{
			Id:               v.Id,
			StudentId:        v.StudentId,
			NimNumber:        v.NimNumber,
			Name:             v.Name,
			StudyProgramName: v.StudyProgramName,
			FileUrl:          fileUrl,
			Point:            v.Point,
		})
	}

	result = objects.ClassWorkSubmissionWithPagination{
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

func (c classWorkService) Create(ctx context.Context, data objects.CreateClassWork) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := c.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	classLecturerData, errs := c.ClassLecturerRepo.GetByClassIdLecturerId(ctx, tx, data.ClassId, claims.ID)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}
	if classLecturerData.Id == "" {
		_ = tx.Rollback()
		return appConstants.ErrUneditableClassWork
	}

	createData := models.CreateClassWork{
		LecturerId:   claims.ID,
		ClassId:      data.ClassId,
		Title:        data.Title,
		Abstraction:  utils.NewNullString(data.Abstraction),
		FilePath:     utils.NewNullString(data.FilePath),
		FilePathType: utils.NewNullString(data.FilePathType),
		StartTime:    data.StartTime,
		EndTime:      data.EndTime,
	}
	errs = c.ClassWorkRepo.Create(ctx, tx, createData)
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

func (c classWorkService) Update(ctx context.Context, data objects.UpdateClassWork) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := c.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	classWorkData, errs := c.ClassWorkRepo.GetDetail(ctx, tx, data.Id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}
	if classWorkData.LecturerId != claims.ID {
		_ = tx.Rollback()
		return appConstants.ErrUneditableClassWork
	}

	updateData := models.UpdateClassWork{
		Id:           data.Id,
		Title:        data.Title,
		Abstraction:  utils.NewNullString(data.Abstraction),
		FilePath:     utils.NewNullString(data.FilePath),
		FilePathType: utils.NewNullString(data.FilePathType),
		StartTime:    data.StartTime,
		EndTime:      data.EndTime,
	}
	errs = c.ClassWorkRepo.Update(ctx, tx, updateData)
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

func (c classWorkService) Delete(ctx context.Context, ids []string) *constants.ErrorResponse {
	tx, err := c.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	errs := c.ClassWorkRepo.Delete(ctx, tx, ids)
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

func (c classWorkService) GradeSubmission(ctx context.Context, classWorkId string, data []objects.GradeClassWorkSubmission) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := c.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	classWorkData, errs := c.ClassWorkRepo.GetDetail(ctx, tx, classWorkId)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}
	if classWorkData.LecturerId != claims.ID {
		_ = tx.Rollback()
		return appConstants.ErrUneditableClassWork
	}

	pointMap := make(map[string]float64)
	ids := []string{}
	for _, v := range data {
		pointMap[v.Id] = v.Point
		ids = append(ids, v.Id)
	}

	paginationData := common.PaginationRequest{
		Page:  constants.DefaultPage,
		Limit: constants.DefaultUnlimited,
	}
	isSubmitted := true
	submissionData, _, errs := c.ClassWorkRepo.GetSubmission(ctx, tx, paginationData, classWorkData.ClassId, classWorkId, ids, &isSubmitted)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	pointData := []models.GradeClassWorkSubmission{}
	for _, v := range submissionData {
		id := utils.NullStringScan(v.Id)
		pointData = append(pointData, models.GradeClassWorkSubmission{
			ClassWorkId:  classWorkId,
			StudentId:    v.StudentId,
			FilePath:     utils.NullStringScan(v.FilePath),
			FilePathType: utils.NullStringScan(v.FilePathType),
			Point:        pointMap[id],
		})
	}

	errs = c.ClassWorkRepo.GradeSubmission(ctx, tx, pointData)
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

func (c classWorkService) Submit(ctx context.Context, classWorkId, filePath, filePathType string) *constants.ErrorResponse {
	now := time.Now()

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := c.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	classWorkData, errs := c.ClassWorkRepo.GetDetail(ctx, tx, classWorkId)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	if now.Before(classWorkData.StartTime) || now.After(classWorkData.EndTime) {
		_ = tx.Rollback()
		return appConstants.ErrClassWorkDeadline
	}

	errs = c.ClassWorkRepo.Submit(ctx, tx, classWorkId, claims.ID, filePath, filePathType)
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
