package student_class

import (
	"context"
	"database/sql"
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

type studentClassService struct {
	*repository.RepoCtx
	*infra.InfraCtx
}

func (s studentClassService) GetList(ctx context.Context, pagination common.PaginationRequest, studyPlanId, studentId, semesterId, appType string, isMbkm *bool) (objects.StudentClassListWithPagination, *constants.ErrorResponse) {
	var result objects.StudentClassListWithPagination

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return result, errs
	}

	tx, err := s.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	if appType == appConstants.AppTypeStudent {
		studentId = claims.ID
	}

	modelResult, paginationResult, errs := s.StudentClassRepo.GetList(ctx, tx, pagination, studyPlanId, studentId, semesterId, isMbkm)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	resultData := []objects.GetStudentClass{}
	if len(modelResult) != 0 {
		classIds := []string{}
		for _, v := range modelResult {
			classIds = append(classIds, v.ClassId)
		}

		scheduleData, errs := s.LectureRepo.GetByClassIds(ctx, tx, classIds)
		if errs != nil {
			_ = tx.Rollback()
			return result, errs
		}

		resultData = mapGetList(modelResult, scheduleData)
	}

	result = objects.StudentClassListWithPagination{
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

func (s studentClassService) TransferStudentClass(ctx context.Context, data objects.TransferStudentClass) *constants.ErrorResponse {
	tx, err := s.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	params := []models.StudentIdClassId{}
	classIds := []string{data.SourceClassId}
	destinationClassMap := make(map[string]string)
	for _, v := range data.Data {
		classIds = append(classIds, v.DestinationClassId)
		destinationClassMap[v.StudentId] = v.DestinationClassId
		params = append(params, models.StudentIdClassId{
			ClassId:   data.SourceClassId,
			StudentId: v.StudentId,
		})
	}

	classes, errs := s.ClassRepo.GetDetailByIds(ctx, tx, classIds, "")
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}
	if len(classes) == 0 {
		_ = tx.Rollback()
		return utils.ErrDataNotFound("class")
	}

	subjectId := classes[0].SubjectId
	for _, v := range classes {
		if subjectId != v.SubjectId {
			_ = tx.Rollback()
			return appConstants.ErrDifferentSubjectClasses
		}
	}

	studentData, errs := s.StudentClassRepo.GetStudentClassByStudentIdClassId(ctx, tx, params)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	newClassData := []models.CreateStudentClass{}
	for _, v := range studentData {
		newClassData = append(newClassData, models.CreateStudentClass{
			Id:                  v.Id,
			StudyPlanId:         v.StudyPlanId,
			CurriculumId:        v.CurriculumId,
			StudentCurriculumId: v.StudentCurriculumId,
			ClassId:             destinationClassMap[v.StudentId],
		})
	}
	errs = s.StudentClassRepo.BulkUpdateClass(ctx, tx, newClassData)
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

func (s studentClassService) ReshuffleStudentClass(ctx context.Context, data []objects.ReshuffleStudentClass) *constants.ErrorResponse {
	tx, err := s.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	params := []models.StudentIdClassId{}
	classIds := []string{}
	destinationClassMap := make(map[string]string)
	for _, v := range data {
		classIds = append(classIds, v.DestinationClassId)
		for _, w := range v.Students {
			classIds = append(classIds, w.SourceClassId)
			params = append(params, models.StudentIdClassId{
				StudentId: w.StudentId,
				ClassId:   w.SourceClassId,
			})
			destinationClassMap[w.StudentId] = v.DestinationClassId
		}
	}

	classes, errs := s.ClassRepo.GetDetailByIds(ctx, tx, classIds, "")
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}
	if len(classes) == 0 {
		_ = tx.Rollback()
		return utils.ErrDataNotFound("class")
	}

	subjectId := classes[0].SubjectId
	for _, v := range classes {
		if subjectId != v.SubjectId {
			_ = tx.Rollback()
			return appConstants.ErrDifferentSubjectClasses
		}
	}

	studentData, errs := s.StudentClassRepo.GetStudentClassByStudentIdClassId(ctx, tx, params)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	newClassData := []models.CreateStudentClass{}
	for _, v := range studentData {
		newClassData = append(newClassData, models.CreateStudentClass{
			Id:                  v.Id,
			StudyPlanId:         v.StudyPlanId,
			CurriculumId:        v.CurriculumId,
			StudentCurriculumId: v.StudentCurriculumId,
			ClassId:             destinationClassMap[v.StudentId],
		})
	}
	errs = s.StudentClassRepo.BulkUpdateClass(ctx, tx, newClassData)
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

func (s studentClassService) MergeStudentClass(ctx context.Context, data objects.MergeStudentClass) *constants.ErrorResponse {
	tx, err := s.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	oldClassId := []string{}
	classIds := []string{data.DestinationClassId}
	for _, v := range data.SourceClassIds {
		classIds = append(classIds, v)
		if v != data.DestinationClassId {
			oldClassId = append(oldClassId, v)
		}
	}

	classes, errs := s.ClassRepo.GetDetailByIds(ctx, tx, classIds, "")
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}
	if len(classes) == 0 {
		_ = tx.Rollback()
		return utils.ErrDataNotFound("class")
	}

	subjectId := classes[0].SubjectId
	for _, v := range classes {
		if subjectId != v.SubjectId {
			_ = tx.Rollback()
			return appConstants.ErrDifferentSubjectClasses
		}
	}

	paginationData := common.PaginationRequest{
		Page:  constants.DefaultPage,
		Limit: constants.DefaultUnlimited,
	}
	oldParticipants, _, errs := s.StudentClassRepo.GetClassParticipant(ctx, tx, paginationData, data.SourceClassIds, "", nil, "")
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	newClassData := []models.CreateStudentClass{}
	for _, v := range oldParticipants {
		newClassData = append(newClassData, models.CreateStudentClass{
			Id:                  v.StudentClassId,
			StudyPlanId:         v.StudyPlanId,
			CurriculumId:        v.CurriculumId,
			StudentCurriculumId: v.StudentCurriculumId,
			ClassId:             data.DestinationClassId,
		})
	}
	errs = s.StudentClassRepo.BulkUpdateClass(ctx, tx, newClassData)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	errs = s.ClassRepo.InactivateClasses(ctx, tx, oldClassId)
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

func (s studentClassService) BulkGradeStudentClass(ctx context.Context, classId string, data []objects.GradeStudentClass) *constants.ErrorResponse {
	now := time.Now()

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := s.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	classData, errs := s.ClassRepo.GetDetail(ctx, tx, classId, claims.ID)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}
	if classData.GradingStartDate != nil && classData.GradingEndDate != nil {
		gradingStartDate := *classData.GradingStartDate
		gradingEndDate := *classData.GradingEndDate
		if gradingStartDate.After(now) || gradingEndDate.Before(now) {
			_ = tx.Rollback()
			return appConstants.ErrNotGradingTime
		}
	}

	var gradedByAdminId sql.NullString
	var gradedByLecturerId sql.NullString
	if claims.Role == appConstants.AppTypeAdmin {
		gradedByAdminId = utils.NewNullString(claims.ID)
	}
	if claims.Role == appConstants.AppTypeLecturer {
		if !utils.NullBooleanScan(classData.IsGradingResponsible) {
			_ = tx.Rollback()
			return appConstants.ErrUngradeableClass
		}
		gradedByLecturerId = utils.NewNullString(claims.ID)
	}

	gradeData := []models.GradeStudentClass{}
	for _, v := range data {
		gradeData = append(gradeData, models.GradeStudentClass{
			ClassId:               classId,
			StudentId:             v.StudentId,
			ClassGradeComponentId: v.ClassGradeComponentId,
			InitialGrade:          v.InitialGrade,
			GradedByAdminId:       gradedByAdminId,
			GradedByLecturerId:    gradedByLecturerId,
		})
	}

	utils.PrintStruct(gradeData)
	errs = s.StudentClassRepo.BulkGradeStudentClass(ctx, tx, gradeData)
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
