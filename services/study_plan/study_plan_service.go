package study_plan

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
	appUtils "github.com/sccicitb/pupr-backend/utils"
)

type studyPlanService struct {
	*repository.RepoCtx
	*infra.InfraCtx
}

func mapGetDetail(studentData models.GetStudent, semesterData models.GetSemesterDetail, studyPlanData models.GetStudyPlan, creditQuotaData []models.GetCreditQuota, classData []models.GetStudentClass, scheduleData []models.GetLectureDetail) objects.GetStudentStudyPlanDetail {
	scheduleMap := make(map[string][]objects.GetStudentStudyPlanDetailClassSchedule)

	for _, v := range scheduleData {
		classId := utils.NullStringScan(v.ClassId)
		scheduleMap[classId] = append(scheduleMap[classId], objects.GetStudentStudyPlanDetailClassSchedule{
			Date:      v.LecturePlanDate,
			StartTime: v.LecturePlanStartTime,
			EndTime:   v.LecturePlanEndTime,
			RoomId:    v.RoomId,
			RoomName:  v.RoomName,
		})
	}

	classes := []objects.GetStudentStudyPlanDetailClass{}
	for _, v := range classData {
		classes = append(classes, objects.GetStudentStudyPlanDetailClass{
			Id:                          v.ClassId,
			Name:                        v.ClassName,
			SubjectId:                   v.SubjectId,
			SubjectName:                 v.SubjectName,
			SubjectCode:                 v.SubjectCode,
			SubjectTheoryCredit:         v.SubjectTheoryCredit,
			SubjectPracticumCredit:      v.SubjectPracticumCredit,
			SubjectFieldPracticumCredit: v.SubjectFieldPracticumCredit,
			TotalLectureDone:            v.TotalLectureDone,
			TotalAttendance:             v.TotalAttendance,
			ActiveLectureId:             v.ActiveLectureId,
			ActiveLectureHasAttend:      v.ActiveLectureHasAttend,
			GradePoint:                  v.GradePoint,
			GradeCode:                   v.GradeCode,
			SubjectIsMandatory:          v.SubjectIsMandatory,
			Schedules:                   scheduleMap[v.ClassId],
		})
	}

	return objects.GetStudentStudyPlanDetail{
		StudyPlanInputStartDate:            semesterData.StudyPlanInputStartDate,
		StudyPlanInputEndDate:              semesterData.StudyPlanInputEndDate,
		Id:                                 studyPlanData.Id,
		IsSubmitted:                        studyPlanData.IsSubmitted,
		IsApproved:                         studyPlanData.IsApproved,
		StudentId:                          studentData.Id,
		StudentNimNumber:                   studentData.NimNumber,
		StudentName:                        studentData.Name,
		StudyProgramId:                     studentData.StudyProgramId,
		StudyProgramName:                   studentData.StudyProgramName,
		SemesterId:                         semesterData.Id,
		SemesterSchoolYear:                 appUtils.GenerateSchoolYear(semesterData.SemesterStartYear),
		SemesterType:                       semesterData.SemesterType,
		MaximumCredit:                      appUtils.GetMaximumCredit(creditQuotaData, studentData.PreviousSemesterGradePoint),
		AcademicGuidanceLecturerId:         studentData.AcademicGuidanceLecturerId,
		AcademicGuidanceLecturerName:       studentData.AcademicGuidanceLecturerName,
		AcademicGuidanceLecturerFrontTitle: studentData.AcademicGuidanceLecturerFrontTitle,
		AcademicGuidanceLecturerBackDegree: studentData.AcademicGuidanceLecturerBackDegree,
		TotalMandatoryCredit:               studyPlanData.TotalMandatoryCredit,
		TotalOptionalCredit:                studyPlanData.TotalOptionalCredit,
		GradePoint:                         studyPlanData.GradePoint,
		Gpa:                                studentData.Gpa,
		IsThesis:                           studyPlanData.IsThesis,
		Classes:                            classes,
	}
}

func (s studyPlanService) BulkCreate(ctx context.Context, data objects.BulkCreateStudyPlan) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := s.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	if claims.Role == appConstants.AppTypeStudent {
		data.StudentIds = []string{claims.ID}
		semesterData, errs := s.SemesterRepo.GetActive(ctx, tx)
		if errs != nil {
			_ = tx.Rollback()
			return errs
		}
		data.SemesterId = semesterData.Id
	}

	errs = s.StudentRepo.UpdateActiveSemesterPackage(ctx, tx, data.StudentIds)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	previousSemesterData, errs := s.SemesterRepo.GetPreviousSemester(ctx, tx, data.SemesterId)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	curriculumData, errs := s.CurriculumRepo.GetActive(ctx, tx)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	curriculums := make(map[string]string)
	for _, v := range curriculumData {
		curriculums[v.StudyProgramId] = v.Id
	}

	activeStudents, errs := s.StudentRepo.GetActive(ctx, tx, previousSemesterData.Id, data.StudentIds)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	studentIds := []string{}
	for _, v := range activeStudents {
		studentIds = append(studentIds, v.Id)
	}
	creditQuotaData, _, errs := s.CreditQuotaRepo.GetList(ctx, tx, common.PaginationRequest{Page: constants.DefaultPage, Limit: constants.DefaultUnlimited})
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	studyPlanData := []models.CreateStudyPlan{}
	studentClassData := []models.CreateStudentClass{}
	deleteStudentClassData := []models.DeleteStudentClassExcludingClassIds{}

	existingStudyPlanData, errs := s.StudyPlanRepo.GetByStudentIdsAndSemesterId(ctx, tx, studentIds, data.SemesterId)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	existingStudyPlanMap := make(map[string]models.GetStudyPlan)
	for _, v := range existingStudyPlanData {
		existingStudyPlanMap[v.StudentId] = v
	}

	for _, v := range activeStudents {
		if curriculums[v.StudyProgramId] == "" {
			_ = tx.Rollback()
			return appConstants.ErrNoActiveCurriculum
		}

		var studyPlanId string
		if existingStudyPlanMap[v.Id].Id != "" {
			if existingStudyPlanMap[v.Id].IsApproved {
				_ = tx.Rollback()
				return appConstants.ErrApprovedStudyPlan
			}
			studyPlanId = existingStudyPlanMap[v.Id].Id
			deleteStudentClassData = append(deleteStudentClassData, models.DeleteStudentClassExcludingClassIds{
				StudyPlanId:      studyPlanId,
				ExcludedClassIds: data.ClassIds,
			})
		} else {
			studyPlanId = uuid.New().String()
		}

		studyPlanData = append(studyPlanData, models.CreateStudyPlan{
			Id:              studyPlanId,
			StudentId:       v.Id,
			SemesterId:      data.SemesterId,
			SemesterPackage: v.CurrentSemesterPackage,
			MaximumCredit:   appUtils.GetMaximumCredit(creditQuotaData, v.PreviousSemesterGradePoint),
			IsSubmitted:     claims.Role == appConstants.AppTypeStudent,
			IsThesis:        data.IsThesis,
		})

		for _, w := range data.ClassIds {
			studentClassData = append(studentClassData, models.CreateStudentClass{
				StudyPlanId:         studyPlanId,
				CurriculumId:        curriculums[v.StudyProgramId],
				StudentCurriculumId: v.CurriculumId,
				ClassId:             w,
			})
		}
	}

	if len(studyPlanData) != 0 {
		errs = s.StudyPlanRepo.BulkCreate(ctx, tx, studyPlanData)
		if errs != nil {
			_ = tx.Rollback()
			return errs
		}
	}

	if !data.IsThesis {
		if len(deleteStudentClassData) != 0 {
			errs = s.StudentClassRepo.BulkDeleteExcludingClassIds(ctx, tx, deleteStudentClassData)
			if errs != nil {
				_ = tx.Rollback()
				return errs
			}
		}
		if len(studentClassData) != 0 {
			errs = s.StudentClassRepo.BulkCreate(ctx, tx, studentClassData)
			if errs != nil {
				_ = tx.Rollback()
				return errs
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return constants.ErrUnknown
	}

	return nil
}

func (s studyPlanService) BulkApprove(ctx context.Context, studyPlanIds []string, isApproved bool) *constants.ErrorResponse {
	tx, err := s.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	errs := s.StudyPlanRepo.BulkApprove(ctx, tx, studyPlanIds, isApproved)
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

func (s studyPlanService) GetList(ctx context.Context, paginationData common.PaginationRequest, studentId, semesterId string) (objects.GetStudyPlanWithPagination, *constants.ErrorResponse) {
	var result objects.GetStudyPlanWithPagination

	tx, err := s.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	modelResult, paginationResult, errs := s.StudyPlanRepo.GetList(ctx, tx, paginationData, studentId, semesterId)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	resultData := []objects.GetListStudyPlan{}
	for _, v := range modelResult {
		resultData = append(resultData, objects.GetListStudyPlan{
			SemesterId:            v.SemesterId,
			SemesterStartYear:     v.SemesterStartYear,
			SchoolYear:            appUtils.GenerateSchoolYear(v.SemesterStartYear),
			SemesterType:          v.SemesterType,
			TotalMandatoryCredit:  v.TotalMandatoryCredit,
			TotalOptionalCredit:   v.TotalOptionalCredit,
			GradePoint:            v.GradePoint,
			StudentId:             v.StudentId,
			StudentNimNumber:      v.StudentNimNumber,
			StudentName:           v.StudentName,
			StudyProgramId:        v.StudyProgramId,
			StudyProgramName:      v.StudyProgramName,
			DiktiStudyProgramCode: v.DiktiStudyProgramCode,
			DiktiStudyProgramType: v.DiktiStudyProgramType,
			StudyLevelShortName:   v.StudyLevelShortName,
			IsThesis:              v.IsThesis,
		})
	}

	result = objects.GetStudyPlanWithPagination{
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

func (s studyPlanService) GetDetail(ctx context.Context, studentId, semesterId string) (objects.GetStudentStudyPlanDetail, *constants.ErrorResponse) {
	var result objects.GetStudentStudyPlanDetail

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return result, errs
	}

	tx, err := s.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	var semesterData models.GetSemesterDetail
	if semesterId == "" {
		semesterData, errs = s.SemesterRepo.GetActive(ctx, tx)
		if errs != nil {
			_ = tx.Rollback()
			return result, errs
		}
		semesterId = semesterData.Id
	} else {
		semesterData, errs = s.SemesterRepo.GetById(ctx, tx, semesterId)
		if errs != nil {
			_ = tx.Rollback()
			return result, errs
		}
	}
	previousSemesterData, errs := s.SemesterRepo.GetPreviousSemester(ctx, tx, semesterData.Id)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	if studentId == "" {
		studentId = claims.ID
	}
	studentData, errs := s.StudentRepo.GetDetail(ctx, tx, studentId, previousSemesterData.Id)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	creditQuotaData, _, errs := s.CreditQuotaRepo.GetList(ctx, tx, common.PaginationRequest{Page: constants.DefaultPage, Limit: constants.DefaultUnlimited})
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	studyPlanData, errs := s.StudyPlanRepo.GetByStudentIdAndSemesterId(ctx, tx, studentId, semesterId)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	studentClassData := []models.GetStudentClass{}
	scheduleData := []models.GetLectureDetail{}
	if studyPlanData.Id != "" {
		studentClassData, _, errs = s.StudentClassRepo.GetList(ctx, tx, common.PaginationRequest{Page: constants.DefaultPage, Limit: constants.DefaultUnlimited}, studyPlanData.Id, "", "", nil)
		if errs != nil {
			_ = tx.Rollback()
			return result, errs
		}

		classIds := []string{}
		for _, v := range studentClassData {
			classIds = append(classIds, v.ClassId)
		}

		scheduleData, errs = s.LectureRepo.GetByClassIds(ctx, tx, classIds)
		if errs != nil {
			_ = tx.Rollback()
			return result, errs
		}
	}

	result = mapGetDetail(studentData, semesterData, studyPlanData, creditQuotaData, studentClassData, scheduleData)

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, constants.ErrUnknown
	}

	return result, nil
}
