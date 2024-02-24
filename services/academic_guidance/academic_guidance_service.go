package academic_guidance

import (
	"context"

	"github.com/google/uuid"
	"github.com/sccicitb/pupr-backend/constants"
	appConstants "github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
	"github.com/sccicitb/pupr-backend/utils"
)

type academicGuidanceService struct {
	*repository.RepoCtx
	*infra.InfraCtx
}

func (a academicGuidanceService) mapGetSessionList(data []models.GetAcademicGuidanceSession, studentData []models.GetAcademicGuidanceSessionStudent, fileData []models.GetAcademicGuidanceSessionFile) ([]objects.GetAcademicGuidanceSession, *constants.ErrorResponse) {
	var result []objects.GetAcademicGuidanceSession

	studentMap := make(map[string][]objects.GetAcademicGuidanceSessionStudent)
	fileMap := make(map[string][]objects.GetAcademicGuidanceSessionFile)

	for _, v := range studentData {
		studentMap[v.AcademicGuidanceSessionId] = append(studentMap[v.AcademicGuidanceSessionId], objects.GetAcademicGuidanceSessionStudent{
			Id:        v.Id,
			Name:      v.Name,
			NimNumber: v.NimNumber,
		})
	}
	for _, v := range fileData {
		fileUrl, errs := a.Storage.GetURL(v.FilePath, v.FilePathType, nil)
		if errs != nil {
			return result, errs
		}
		fileMap[v.AcademicGuidanceSessionId] = append(fileMap[v.AcademicGuidanceSessionId], objects.GetAcademicGuidanceSessionFile{
			Id:           v.Id,
			Title:        v.Title,
			FileUrl:      fileUrl,
			FilePath:     v.FilePath,
			FilePathType: v.FilePathType,
		})
	}

	for _, v := range data {
		result = append(result, objects.GetAcademicGuidanceSession{
			Id:                 v.Id,
			AcademicGuidanceId: v.AcademicGuidanceId,
			Subject:            v.Subject,
			SessionDate:        v.SessionDate,
			Summary:            v.Summary,
			Files:              fileMap[v.Id],
			Students:           studentMap[v.Id],
		})
	}

	return result, nil
}

func (f academicGuidanceService) GetListStudent(ctx context.Context, paginationData common.PaginationRequest, lecturerId, semesterId string) (objects.AcademicGuidanceStudentListWithPagination, *constants.ErrorResponse) {
	var result objects.AcademicGuidanceStudentListWithPagination

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return result, errs
	}

	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	if claims.Role == appConstants.AppTypeLecturer {
		lecturerId = claims.ID
		if semesterId == "" {
			semesterData, errs := f.SemesterRepo.GetActive(ctx, tx)
			if errs != nil {
				_ = tx.Rollback()
				return result, errs
			}
			semesterId = semesterData.Id
		}
	}

	modelResult, paginationResult, errs := f.AcademicGuidanceRepo.GetListStudent(ctx, tx, paginationData, lecturerId, semesterId)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	resultData := []objects.GetAcademicGuidanceStudent{}
	for _, v := range modelResult {
		resultData = append(resultData, objects.GetAcademicGuidanceStudent{
			Id:                      v.Id,
			NimNumber:               v.NimNumber,
			StudentForce:            v.StudentForce,
			Name:                    v.Name,
			Status:                  v.Status,
			StudyPlanFormIsApproved: v.StudyPlanFormIsApproved,
		})
	}

	result = objects.AcademicGuidanceStudentListWithPagination{
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

func (f academicGuidanceService) GetDetail(ctx context.Context, semesterId, studentId string) (objects.GetAcademicGuidanceDetail, *constants.ErrorResponse) {
	var result objects.GetAcademicGuidanceDetail

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return result, errs
	}

	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	if claims.Role == appConstants.AppTypeStudent {
		studentId = claims.ID
		if semesterId == "" {
			semesterData, errs := f.SemesterRepo.GetActive(ctx, tx)
			if errs != nil {
				_ = tx.Rollback()
				return result, errs
			}
			semesterId = semesterData.Id
		}
	}

	resultData, errs := f.AcademicGuidanceRepo.GetDetailBySemesterIdStudentId(ctx, tx, semesterId, studentId)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	result = objects.GetAcademicGuidanceDetail{
		Id:                 resultData.Id,
		SemesterId:         resultData.SemesterId,
		LecturerId:         resultData.LecturerId,
		LecturerName:       resultData.LecturerName,
		LecturerFrontTitle: resultData.LecturerFrontTitle,
		LecturerBackDegree: resultData.LecturerBackDegree,
		DecisionNumber:     resultData.DecisionNumber,
		DecisionDate:       resultData.DecisionDate,
		TotalStudent:       resultData.TotalStudent,
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, constants.ErrUnknown
	}

	return result, nil
}

func (f academicGuidanceService) Upsert(ctx context.Context, data objects.UpsertAcademicGuidance) *constants.ErrorResponse {
	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	upsertData := models.UpsertAcademicGuidance{
		SemesterId: data.SemesterId,
		LecturerId: data.LecturerId,
		CreatedBy:  claims.ID,
	}
	academicGuidanceId, errs := f.AcademicGuidanceRepo.Upsert(ctx, tx, upsertData)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	if len(data.StudentIds) != 0 {
		studentData := []models.UpsertAcademicGuidanceStudent{}
		for _, v := range data.StudentIds {
			studentData = append(studentData, models.UpsertAcademicGuidanceStudent{
				AcademicGuidanceId: academicGuidanceId,
				StudentId:          v,
				CreatedBy:          claims.ID,
			})
		}

		errs = f.AcademicGuidanceRepo.UpsertStudent(ctx, tx, studentData)
		if errs != nil {
			_ = tx.Rollback()
			return errs
		}
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return constants.ErrUnknown
	}

	return nil
}

func (f academicGuidanceService) UpsertDecision(ctx context.Context, data objects.UpsertDecisionAcademicGuidance) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	updateData := models.UpsertDecisionAcademicGuidance{
		LecturerId:     data.LecturerId,
		SemesterId:     data.SemesterId,
		DecisionNumber: data.DecisionNumber,
		DecisionDate:   data.DecisionDate,
		CreatedBy:      claims.ID,
	}
	errs = f.AcademicGuidanceRepo.UpsertDecision(ctx, tx, updateData)
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

func (f academicGuidanceService) GetSessionList(ctx context.Context, academicGuidanceId, semesterId, lecturerId string) ([]objects.GetAcademicGuidanceSession, *constants.ErrorResponse) {
	var result []objects.GetAcademicGuidanceSession

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return result, errs
	}

	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	if claims.Role == appConstants.AppTypeLecturer {
		lecturerId = claims.ID
		if semesterId == "" {
			semesterData, errs := f.SemesterRepo.GetActive(ctx, tx)
			if errs != nil {
				_ = tx.Rollback()
				return result, errs
			}
			semesterId = semesterData.Id
		}
	}

	resultData, errs := f.AcademicGuidanceRepo.GetSession(ctx, tx, academicGuidanceId, semesterId, lecturerId)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	var sessionIds []string
	for _, v := range resultData {
		sessionIds = append(sessionIds, v.Id)
	}

	var studentSessionData []models.GetAcademicGuidanceSessionStudent
	var fileSessionData []models.GetAcademicGuidanceSessionFile
	if len(sessionIds) != 0 {
		studentSessionData, errs = f.AcademicGuidanceRepo.GetSessionStudent(ctx, tx, sessionIds)
		if errs != nil {
			_ = tx.Rollback()
			return result, errs
		}
		fileSessionData, errs = f.AcademicGuidanceRepo.GetSessionFile(ctx, tx, sessionIds)
		if errs != nil {
			_ = tx.Rollback()
			return result, errs
		}
	}

	result, errs = f.mapGetSessionList(resultData, studentSessionData, fileSessionData)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, constants.ErrUnknown
	}

	return result, nil
}

func (f academicGuidanceService) CreateSession(ctx context.Context, data objects.CreateAcademicGuidanceSession) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	if claims.Role == appConstants.AppTypeLecturer {
		if data.SemesterId == "" {
			semesterData, errs := f.SemesterRepo.GetActive(ctx, tx)
			if errs != nil {
				_ = tx.Rollback()
				return errs
			}
			data.SemesterId = semesterData.Id
		}
		academicGuidanceData, errs := f.AcademicGuidanceRepo.GetDetailBySemesterIdLecturerId(ctx, tx, data.SemesterId, claims.ID)
		if errs != nil {
			_ = tx.Rollback()
			return errs
		}
		data.AcademicGuidanceId = academicGuidanceData.Id
	}

	sessionId := uuid.New().String()
	createData := models.UpsertAcademicGuidanceSession{
		Id:                 sessionId,
		AcademicGuidanceId: data.AcademicGuidanceId,
		Subject:            data.Subject,
		SessionDate:        data.SessionDate,
		Summary:            data.Summary,
	}
	errs = f.AcademicGuidanceRepo.UpsertSession(ctx, tx, createData)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	var studentData []models.UpsertAcademicGuidanceSessionStudent
	for _, v := range data.StudentIds {
		studentData = append(studentData, models.UpsertAcademicGuidanceSessionStudent{
			AcademicGuidanceSessionId: sessionId,
			StudentId:                 v,
		})
	}
	if len(studentData) != 0 {
		errs = f.AcademicGuidanceRepo.UpsertSessionStudent(ctx, tx, studentData)
		if errs != nil {
			_ = tx.Rollback()
			return errs
		}
	}

	var fileData []models.UpsertAcademicGuidanceSessionFile
	for _, v := range data.Files {
		fileData = append(fileData, models.UpsertAcademicGuidanceSessionFile{
			AcademicGuidanceSessionId: sessionId,
			FilePath:                  v.FilePath,
			FilePathType:              v.FilePathType,
		})
	}
	if len(fileData) != 0 {
		errs = f.AcademicGuidanceRepo.UpsertSessionFile(ctx, tx, fileData)
		if errs != nil {
			_ = tx.Rollback()
			return errs
		}
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return constants.ErrUnknown
	}

	return nil
}

func (f academicGuidanceService) UpdateSession(ctx context.Context, data objects.UpdateAcademicGuidanceSession) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	sessionData, errs := f.AcademicGuidanceRepo.GetSessionById(ctx, tx, data.Id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	if claims.Role == appConstants.AppTypeLecturer {
		academicGuidanceData, errs := f.AcademicGuidanceRepo.GetDetail(ctx, tx, sessionData.AcademicGuidanceId)
		if errs != nil {
			_ = tx.Rollback()
			return errs
		}
		if academicGuidanceData.LecturerId != claims.ID {
			_ = tx.Rollback()
			return constants.ErrEligbleAccess
		}
	}

	createData := models.UpsertAcademicGuidanceSession{
		Id:                 sessionData.Id,
		AcademicGuidanceId: sessionData.AcademicGuidanceId,
		Subject:            data.Subject,
		SessionDate:        data.SessionDate,
		Summary:            data.Summary,
	}
	errs = f.AcademicGuidanceRepo.UpsertSession(ctx, tx, createData)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	errs = f.AcademicGuidanceRepo.DeleteSessionStudentExcludingStudentIds(ctx, tx, sessionData.Id, data.StudentIds)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	var studentData []models.UpsertAcademicGuidanceSessionStudent
	for _, v := range data.StudentIds {
		studentData = append(studentData, models.UpsertAcademicGuidanceSessionStudent{
			AcademicGuidanceSessionId: sessionData.Id,
			StudentId:                 v,
		})
	}
	if len(studentData) != 0 {
		errs = f.AcademicGuidanceRepo.UpsertSessionStudent(ctx, tx, studentData)
		if errs != nil {
			_ = tx.Rollback()
			return errs
		}
	}

	var fileData []models.UpsertAcademicGuidanceSessionFile
	var filePaths []string
	for _, v := range data.Files {
		filePaths = append(filePaths, v.FilePath)
		fileData = append(fileData, models.UpsertAcademicGuidanceSessionFile{
			AcademicGuidanceSessionId: sessionData.Id,
			Title:                     v.Title,
			FilePath:                  v.FilePath,
			FilePathType:              v.FilePathType,
		})
	}
	errs = f.AcademicGuidanceRepo.DeleteSessionFileExcludingFilePaths(ctx, tx, sessionData.Id, filePaths)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	if len(fileData) != 0 {
		errs = f.AcademicGuidanceRepo.UpsertSessionFile(ctx, tx, fileData)
		if errs != nil {
			_ = tx.Rollback()
			return errs
		}
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return constants.ErrUnknown
	}

	return nil
}

func (f academicGuidanceService) DeleteSession(ctx context.Context, id string) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	sessionData, errs := f.AcademicGuidanceRepo.GetSessionById(ctx, tx, id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	if claims.Role == appConstants.AppTypeLecturer {
		academicGuidanceData, errs := f.AcademicGuidanceRepo.GetDetail(ctx, tx, sessionData.AcademicGuidanceId)
		if errs != nil {
			_ = tx.Rollback()
			return errs
		}
		if academicGuidanceData.LecturerId != claims.ID {
			_ = tx.Rollback()
			return constants.ErrEligbleAccess
		}
	}

	errs = f.AcademicGuidanceRepo.DeleteSession(ctx, tx, id)
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
