package class_exam

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

type classExamService struct {
	*repository.RepoCtx
	*infra.InfraCtx
}

func (c classExamService) GetList(ctx context.Context, paginationData common.PaginationRequest, appType string, classIds []string) (objects.ClassExamListWithPagination, *constants.ErrorResponse) {
	var result objects.ClassExamListWithPagination

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
		studentId = claims.ID
	}

	modelResult, paginationResult, errs := c.ClassExamRepo.GetList(ctx, tx, paginationData, classIds, studentId, nil)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	resultData := []objects.GetClassExam{}
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

		resultData = append(resultData, objects.GetClassExam{
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

	result = objects.ClassExamListWithPagination{
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

func (c classExamService) GetSubmission(ctx context.Context, paginationData common.PaginationRequest, classExamId string) (objects.ClassExamSubmissionWithPagination, *constants.ErrorResponse) {
	var result objects.ClassExamSubmissionWithPagination

	tx, err := c.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	classExamData, errs := c.ClassExamRepo.GetDetail(ctx, tx, classExamId)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	modelResult, paginationResult, errs := c.ClassExamRepo.GetSubmission(ctx, tx, paginationData, classExamData.ClassId, classExamId, nil, nil)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	resultData := []objects.GetClassExamSubmission{}
	for _, v := range modelResult {
		var fileUrl string
		if v.FilePath != nil && v.FilePathType != nil {
			fileUrl, errs = c.Storage.GetURL(*v.FilePath, *v.FilePathType, nil)
			if errs != nil {
				_ = tx.Rollback()
				return result, errs
			}
		}

		resultData = append(resultData, objects.GetClassExamSubmission{
			Id:               v.Id,
			StudentId:        v.StudentId,
			NimNumber:        v.NimNumber,
			Name:             v.Name,
			StudyProgramName: v.StudyProgramName,
			FileUrl:          fileUrl,
			Point:            v.Point,
		})
	}

	result = objects.ClassExamSubmissionWithPagination{
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

func (c classExamService) Create(ctx context.Context, data objects.CreateClassExam) *constants.ErrorResponse {
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
		return appConstants.ErrUneditableClassExam
	}

	createData := models.CreateClassExam{
		LecturerId:   claims.ID,
		ClassId:      data.ClassId,
		Title:        data.Title,
		Abstraction:  utils.NewNullString(data.Abstraction),
		FilePath:     utils.NewNullString(data.FilePath),
		FilePathType: utils.NewNullString(data.FilePathType),
		StartTime:    data.StartTime,
		EndTime:      data.EndTime,
	}
	errs = c.ClassExamRepo.Create(ctx, tx, createData)
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

func (c classExamService) Update(ctx context.Context, data objects.UpdateClassExam) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := c.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	classExamData, errs := c.ClassExamRepo.GetDetail(ctx, tx, data.Id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}
	if classExamData.LecturerId != claims.ID {
		_ = tx.Rollback()
		return appConstants.ErrUneditableClassExam
	}

	updateData := models.UpdateClassExam{
		Id:           data.Id,
		Title:        data.Title,
		Abstraction:  utils.NewNullString(data.Abstraction),
		FilePath:     utils.NewNullString(data.FilePath),
		FilePathType: utils.NewNullString(data.FilePathType),
		StartTime:    data.StartTime,
		EndTime:      data.EndTime,
	}
	errs = c.ClassExamRepo.Update(ctx, tx, updateData)
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

func (c classExamService) Delete(ctx context.Context, ids []string) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := c.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	paginationData := common.PaginationRequest{
		Page:  constants.DefaultPage,
		Limit: constants.DefaultUnlimited,
	}
	classExamData, _, errs := c.ClassExamRepo.GetList(ctx, tx, paginationData, nil, "", ids)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}
	for _, v := range classExamData {
		if v.LecturerId != claims.ID {
			_ = tx.Rollback()
			return appConstants.ErrUneditableClassExam
		}
	}

	errs = c.ClassExamRepo.Delete(ctx, tx, ids)
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

func (c classExamService) GradeSubmission(ctx context.Context, classExamId string, data []objects.GradeClassExamSubmission) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := c.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	classExamData, errs := c.ClassExamRepo.GetDetail(ctx, tx, classExamId)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}
	if classExamData.LecturerId != claims.ID {
		_ = tx.Rollback()
		return appConstants.ErrUneditableClassExam
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
	submissionData, _, errs := c.ClassExamRepo.GetSubmission(ctx, tx, paginationData, classExamData.ClassId, classExamId, ids, &isSubmitted)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	pointData := []models.GradeClassExamSubmission{}
	for _, v := range submissionData {
		id := utils.NullStringScan(v.Id)
		pointData = append(pointData, models.GradeClassExamSubmission{
			ClassExamId:  classExamId,
			StudentId:    v.StudentId,
			FilePath:     utils.NullStringScan(v.FilePath),
			FilePathType: utils.NullStringScan(v.FilePathType),
			Point:        pointMap[id],
		})
	}

	errs = c.ClassExamRepo.GradeSubmission(ctx, tx, pointData)
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

func (c classExamService) Submit(ctx context.Context, classExamId, filePath, filePathType string) *constants.ErrorResponse {
	now := time.Now()
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := c.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	classExamData, errs := c.ClassExamRepo.GetDetail(ctx, tx, classExamId)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	if now.Before(classExamData.StartTime) || now.After(classExamData.EndTime) {
		_ = tx.Rollback()
		return appConstants.ErrClassExamDeadline
	}

	errs = c.ClassExamRepo.Submit(ctx, tx, classExamId, claims.ID, filePath, filePathType)
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
