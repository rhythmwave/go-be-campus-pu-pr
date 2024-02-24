package class

import (
	"context"
	"math"

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

type classService struct {
	*repository.RepoCtx
	*infra.InfraCtx
}

func mapGetList(data []models.GetClass, lecturerData []models.GetClassLecturer) []objects.GetClass {
	results := []objects.GetClass{}

	lecturers := make(map[string][]objects.GetClassLecturer)
	for _, v := range lecturerData {
		lecturers[v.ClassId] = append(lecturers[v.ClassId], objects.GetClassLecturer{
			Id:         v.LecturerId,
			Name:       v.LecturerName,
			FrontTitle: v.LecturerFrontTitle,
			BackDegree: v.LecturerBackDegree,
		})
	}

	for _, v := range data {
		results = append(results, objects.GetClass{
			Id:                          v.Id,
			Name:                        v.Name,
			SubjectId:                   v.SubjectId,
			SubjectCode:                 v.SubjectCode,
			SubjectName:                 v.SubjectName,
			SubjectIsMandatory:          v.SubjectIsMandatory,
			SubjectSemesterPackage:      v.SubjectSemesterPackage,
			MaximumParticipant:          v.MaximumParticipant,
			TotalParticipant:            v.TotalParticipant,
			SubjectTheoryCredit:         v.SubjectTheoryCredit,
			SubjectPracticumCredit:      v.SubjectPracticumCredit,
			SubjectFieldPracticumCredit: v.SubjectFieldPracticumCredit,
			IsActive:                    v.IsActive,
			UnapprovedStudyPlan:         v.UnapprovedStudyPlan,
			TotalMaterial:               v.TotalMaterial,
			TotalWork:                   v.TotalWork,
			TotalDiscussion:             v.TotalDiscussion,
			TotalExam:                   v.TotalExam,
			TotalEvent:                  v.TotalEvent,
			TotalLecturePlan:            v.TotalLecturePlan,
			TotalLectureDone:            v.TotalLectureDone,
			TotalGradedParticipant:      v.TotalGradedParticipant,
			StudyLevelId:                v.StudyLevelId,
			ApplicationDeadline:         v.ApplicationDeadline,
			Lecturers:                   lecturers[v.Id],
			CurriculumId:                v.CurriculumId,
			CurriculumName:              v.CurriculumName,
			SubjectTotalLessonPlan:      v.SubjectTotalLessonPlan,
			StudyProgramId:              v.StudyProgramId,
			StudyProgramName:            v.StudyProgramName,
			SemesterId:                  v.SemesterId,
			SemesterStartYear:           v.SemesterStartYear,
			SchoolYear:                  utils.GenerateSchoolYear(v.SemesterStartYear),
			SemesterType:                v.SemesterType,
		})
	}

	return results
}

func mapGetDetail(resultData models.GetClassDetail, lecturerData []models.GetClassLecturer, gradeComponentData []models.GetSubjectGradeComponent, studentData []models.GetClassParticipant, gradeData []models.GetStudentClassGrade, gradeTypeData []models.GetGradeType) objects.GetClassDetail {
	lecturers := []objects.GetClassDetailLecturer{}
	for _, v := range lecturerData {
		lecturers = append(lecturers, objects.GetClassDetailLecturer{
			Id:                   v.LecturerId,
			Name:                 v.LecturerName,
			FrontTitle:           v.LecturerFrontTitle,
			BackDegree:           v.LecturerBackDegree,
			IsGradingResponsible: v.IsGradingResponsible,
		})
	}

	gradeComponents := []objects.GetClassDetailGradeComponent{}
	for _, v := range gradeComponentData {
		gradeComponents = append(gradeComponents, objects.GetClassDetailGradeComponent{
			Id:         v.Id,
			Name:       v.Name,
			Percentage: v.Percentage,
		})
	}

	gradeMap := make(map[string][]objects.GetClassDetailStudentGrade)
	for _, v := range gradeData {
		gradeMap[v.StudentId] = append(gradeMap[v.StudentId], objects.GetClassDetailStudentGrade{
			ClassGradeComponentId:   v.ClassGradeComponentId,
			ClassGradeComponentName: v.ClassGradeComponentName,
			InitialGrade:            v.InitialGrade,
			FinalGrade:              v.FinalGrade,
		})
	}

	students := []objects.GetClassDetailStudent{}
	for _, v := range studentData {
		students = append(students, objects.GetClassDetailStudent{
			Id:         v.StudentId,
			NimNumber:  v.StudentNimNumber,
			Name:       v.StudentName,
			GradePoint: v.GradePoint,
			GradeCode:  v.GradeCode,
			Grades:     gradeMap[v.StudentId],
		})
	}

	gradeTypes := []objects.GetClassDetailGradeType{}
	for _, v := range gradeTypeData {
		gradeTypes = append(gradeTypes, objects.GetClassDetailGradeType{
			Id:                  v.Id,
			StudyLevelId:        v.StudyLevelId,
			StudyLevelShortName: v.StudyLevelShortName,
			Code:                v.Code,
			GradePoint:          v.GradePoint,
			MinimumGrade:        v.MinimumGrade,
			MaximumGrade:        v.MaximumGrade,
			GradeCategory:       v.GradeCategory,
			GradePointCategory:  v.GradePointCategory,
			Label:               v.Label,
			EnglishLabel:        v.EnglishLabel,
			StartDate:           v.StartDate,
			EndDate:             v.EndDate,
		})
	}

	return objects.GetClassDetail{
		Id:                    resultData.Id,
		Name:                  resultData.Name,
		StudyProgramId:        resultData.StudyProgramId,
		StudyProgramName:      resultData.StudyProgramName,
		DiktiStudyProgramType: resultData.DiktiStudyProgramType,
		StudyLevelShortName:   resultData.StudyLevelShortName,
		CurriculumId:          resultData.CurriculumId,
		CurriculumName:        resultData.CurriculumName,
		CurriculumYear:        resultData.CurriculumYear,
		SemesterId:            resultData.SemesterId,
		SemesterStartYear:     resultData.SemesterStartYear,
		SchoolYear:            appUtils.GenerateSchoolYear(resultData.SemesterStartYear),
		SemesterType:          resultData.SemesterType,
		GradingStartDate:      resultData.GradingStartDate,
		GradingEndDate:        resultData.GradingEndDate,
		SubjectId:             resultData.SubjectId,
		SubjectCode:           resultData.SubjectCode,
		SubjectName:           resultData.SubjectName,
		Scope:                 resultData.Scope,
		IsOnline:              resultData.IsOnline,
		IsOffline:             resultData.IsOffline,
		MinimumParticipant:    resultData.MinimumParticipant,
		MaximumParticipant:    resultData.MaximumParticipant,
		TotalParticipant:      resultData.TotalParticipant,
		Remarks:               appUtils.NullStringScan(resultData.Remarks),
		IsActive:              resultData.IsActive,
		StudyLevelId:          resultData.StudyLevelId,
		ApplicationDeadline:   resultData.ApplicationDeadline,
		IsGradingResponsible:  resultData.IsGradingResponsible,
		Lecturers:             lecturers,
		GradeComponents:       gradeComponents,
		Students:              students,
		GradeTypes:            gradeTypes,
	}
}

func mapGetParticipant(totalLecture uint32, data []models.GetClassParticipant, gradeData []models.GetStudentClassGrade) []objects.GetClassParticipant {
	results := []objects.GetClassParticipant{}

	gradeResult := make(map[string][]objects.GetClassParticipantGrade)
	for _, v := range gradeData {
		gradeResult[v.StudentId] = append(gradeResult[v.StudentId], objects.GetClassParticipantGrade{
			ClassGradeComponentId:   v.ClassGradeComponentId,
			ClassGradeComponentName: v.ClassGradeComponentName,
			InitialGrade:            v.InitialGrade,
			FinalGrade:              v.FinalGrade,
		})
	}

	for _, v := range data {
		var attendancePercentage float64
		if totalLecture != 0 {
			attendancePercentage = math.Round((float64(v.TotalAttendance)/float64(totalLecture))*10000) / 100
		}
		results = append(results, objects.GetClassParticipant{
			StudentId:             v.StudentId,
			StudentNimNumber:      v.StudentNimNumber,
			StudentName:           v.StudentName,
			StudyProgramId:        v.StudyProgramId,
			StudyProgramName:      v.StudyProgramName,
			DiktiStudyProgramCode: v.DiktiStudyProgramCode,
			DiktiStudyProgramType: v.DiktiStudyProgramType,
			StudyLevelShortName:   v.StudyLevelShortName,
			TotalAttendance:       v.TotalAttendance,
			AttendancePercentage:  attendancePercentage,
			TotalSick:             v.TotalSick,
			TotalLeave:            v.TotalLeave,
			TotalAwol:             v.TotalAwol,
			IsAttend:              v.IsAttend,
			IsSick:                v.IsSick,
			IsLeave:               v.IsLeave,
			IsAwol:                v.IsAwol,
			GradePoint:            v.GradePoint,
			GradeCode:             v.GradeCode,
			GradedByAdminId:       v.GradedByAdminId,
			GradedByAdminName:     v.GradedByAdminName,
			GradedByLecturerId:    v.GradedByLecturerId,
			GradedByLecturerName:  v.GradedByLecturerName,
			GradedAt:              v.GradedAt,
			SubjectRepetition:     v.SubjectRepetition,
			SubjectName:           v.SubjectName,
			Grades:                gradeResult[v.StudentId],
		})
	}
	return results
}

func (p classService) GetList(ctx context.Context, paginationData common.PaginationRequest, in objects.GetClassListRequest) (objects.ClassListWithPagination, *constants.ErrorResponse) {
	var result objects.ClassListWithPagination

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return result, errs
	}

	tx, err := p.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	if in.StudyProgramId != "" {
		_, errs = p.StudyProgramRepo.GetDetail(ctx, tx, in.StudyProgramId, claims.Role, claims.ID)
		if errs != nil {
			_ = tx.Rollback()
			return result, errs
		}
	}

	if in.SemesterId == "" {
		activeSemesterData, errs := p.SemesterRepo.GetActive(ctx, tx)
		if errs != nil {
			_ = tx.Rollback()
			return result, errs
		}
		in.SemesterId = activeSemesterData.Id
		if in.FollowSemesterIdParity {
			isOddSemester := activeSemesterData.SemesterType == appConstants.OddSemester
			in.ForOddSemester = &isOddSemester
		}
	} else if in.FollowSemesterIdParity {
		semesterData, errs := p.SemesterRepo.GetById(ctx, tx, in.SemesterId)
		if errs != nil {
			_ = tx.Rollback()
			return result, errs
		}
		isOddSemester := semesterData.SemesterType == appConstants.OddSemester
		in.ForOddSemester = &isOddSemester
	}

	modelResult, paginationResult, errs := p.ClassRepo.GetList(ctx, tx, paginationData, in)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	classIds := []string{}
	for _, v := range modelResult {
		classIds = append(classIds, v.Id)
	}

	lecturerData := []models.GetClassLecturer{}
	if len(classIds) != 0 {
		lecturerData, errs = p.ClassRepo.GetClassLecturerByClassIds(ctx, tx, classIds)
		if errs != nil {
			_ = tx.Rollback()
			return result, errs
		}
	}

	result = objects.ClassListWithPagination{
		Pagination: paginationResult,
		Data:       mapGetList(modelResult, lecturerData),
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, constants.ErrUnknown
	}

	return result, nil
}

func (f classService) GetDetail(ctx context.Context, id string) (objects.GetClassDetail, *constants.ErrorResponse) {
	var result objects.GetClassDetail

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return result, errs
	}

	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	var lecturerId string
	if claims.Role == appConstants.AppTypeLecturer {
		lecturerId = claims.ID
	}
	resultData, errs := f.ClassRepo.GetDetail(ctx, tx, id, lecturerId)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	lecturerData, errs := f.ClassRepo.GetClassLecturerByClassIds(ctx, tx, []string{id})
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	gradeComponentData, errs := f.SubjectGradeComponentRepo.GetBySubjectId(ctx, tx, resultData.SubjectId)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	studentData, _, errs := f.StudentClassRepo.GetClassParticipant(ctx, tx, common.PaginationRequest{Page: constants.DefaultPage, Limit: constants.DefaultUnlimited}, []string{id}, "", nil, "")
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}
	studentIds := []string{}
	for _, v := range studentData {
		studentIds = append(studentIds, v.StudentId)
	}

	gradeData := []models.GetStudentClassGrade{}
	if len(studentIds) != 0 {
		gradeData, errs = f.StudentClassRepo.GetStudentClassGradeByClassIdAndStudentIds(ctx, tx, id, studentIds)
		if errs != nil {
			_ = tx.Rollback()
			return result, errs
		}
	}

	gradeTypeData, _, errs := f.GradeTypeRepo.GetList(ctx, tx, common.PaginationRequest{Page: constants.DefaultPage, Limit: constants.DefaultUnlimited}, resultData.StudyLevelId)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	result = mapGetDetail(resultData, lecturerData, gradeComponentData, studentData, gradeData, gradeTypeData)

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, constants.ErrUnknown
	}

	return result, nil
}

func (a classService) Create(ctx context.Context, data objects.CreateClass) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	createData := models.CreateClass{
		SubjectId:           data.SubjectId,
		SemesterId:          data.SemesterId,
		Name:                data.Name,
		Scope:               utils.NewNullString(data.Scope),
		IsOnline:            utils.NewNullBoolean(data.IsOnline),
		IsOffline:           utils.NewNullBoolean(data.IsOffline),
		MinimumParticipant:  utils.NewNullInt32(int32(data.MinimumParticipant)),
		MaximumParticipant:  utils.NewNullInt32(int32(data.MaximumParticipant)),
		Remarks:             utils.NewNullString(data.Remarks),
		ApplicationDeadline: utils.NewNullTime(data.ApplicationDeadline),
		CreatedBy:           claims.ID,
	}
	classId, errs := a.ClassRepo.Create(ctx, tx, createData)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	lecturerData := []models.UpsertClassLecturer{}
	for _, v := range data.Lecturers {
		lecturerData = append(lecturerData, models.UpsertClassLecturer{
			ClassId:              classId,
			LecturerId:           v.Id,
			IsGradingResponsible: v.IsGradingResponsible,
			CreatedBy:            claims.ID,
		})
	}

	if len(lecturerData) != 0 {
		errs = a.ClassRepo.UpsertClassLecturer(ctx, tx, lecturerData)
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

func (a classService) Update(ctx context.Context, data objects.UpdateClass) *constants.ErrorResponse {
	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs := a.ClassRepo.GetDetail(ctx, tx, data.Id, "")
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	updateData := models.UpdateClass{
		Id:                  data.Id,
		SubjectId:           data.SubjectId,
		Name:                data.Name,
		Scope:               data.Scope,
		IsOnline:            data.IsOnline,
		IsOffline:           data.IsOffline,
		MinimumParticipant:  data.MinimumParticipant,
		MaximumParticipant:  utils.NewNullInt32(int32(data.MaximumParticipant)),
		Remarks:             utils.NewNullString(data.Remarks),
		ApplicationDeadline: utils.NewNullTime(data.ApplicationDeadline),
		UpdatedBy:           claims.ID,
	}
	errs = a.ClassRepo.Update(ctx, tx, updateData)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	lecturerData := []models.UpsertClassLecturer{}
	lecturerIds := []string{}
	for _, v := range data.Lecturers {
		lecturerIds = append(lecturerIds, v.Id)
		lecturerData = append(lecturerData, models.UpsertClassLecturer{
			ClassId:              data.Id,
			LecturerId:           v.Id,
			IsGradingResponsible: v.IsGradingResponsible,
			CreatedBy:            claims.ID,
		})
	}

	errs = a.ClassRepo.DeleteClassLecturerExcludingLecturerIds(ctx, tx, data.Id, lecturerIds)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	if len(lecturerData) != 0 {
		errs = a.ClassRepo.UpsertClassLecturer(ctx, tx, lecturerData)
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

func (a classService) UpdateActivation(ctx context.Context, id string, isActive bool) *constants.ErrorResponse {
	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs := a.ClassRepo.GetDetail(ctx, tx, id, "")
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	errs = a.ClassRepo.UpdateActivation(ctx, tx, id, isActive)
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

func (a classService) Delete(ctx context.Context, id string) *constants.ErrorResponse {
	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs := a.ClassRepo.GetDetail(ctx, tx, id, "")
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	errs = a.ClassRepo.Delete(ctx, tx, id)
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

func (a classService) Duplicate(ctx context.Context, fromSemesterId, toSemesterId string) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	if fromSemesterId == toSemesterId {
		_ = tx.Rollback()
		return appConstants.ErrIdenticalSemesterId
	}

	errs = a.ClassRepo.Duplicate(ctx, tx, fromSemesterId, toSemesterId, claims.ID)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	errs = a.ClassRepo.DuplicateLecturer(ctx, tx, fromSemesterId, toSemesterId, claims.ID)
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

func (a classService) BulkUpdateMaximumParticipant(ctx context.Context, data []objects.UpdateClassMaximumParticipant) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	if len(data) == 0 {
		_ = tx.Rollback()
		return utils.ErrDataNotFound("class")
	}
	ids := []string{}
	maximumParticipantMap := make(map[string]uint32)
	for _, v := range data {
		ids = append(ids, v.Id)
		maximumParticipantMap[v.Id] = v.MaximumParticipant
	}

	classData, errs := a.ClassRepo.GetDetailByIds(ctx, tx, ids, "")
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	upsertData := []models.CreateClass{}
	for _, v := range classData {
		upsertData = append(upsertData, models.CreateClass{
			Id:                 v.Id,
			SubjectId:          v.SubjectId,
			SemesterId:         v.SemesterId,
			Name:               v.Name,
			Scope:              utils.NewNullString(utils.NullStringScan(v.Scope)),
			IsOnline:           utils.NewNullBoolean(utils.NullBooleanScan(v.IsOnline)),
			IsOffline:          utils.NewNullBoolean(utils.NullBooleanScan(v.IsOffline)),
			MinimumParticipant: utils.NewNullInt32(int32(utils.NullUint32Scan(v.MinimumParticipant))),
			MaximumParticipant: utils.NewNullInt32(int32(maximumParticipantMap[v.Id])),
			Remarks:            utils.NewNullString(utils.NullStringScan(v.Remarks)),
			CreatedBy:          claims.ID,
		})
	}

	errs = a.ClassRepo.UpsertMaximumParticipant(ctx, tx, upsertData)
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

func (p classService) GetClassParticipantList(ctx context.Context, paginationData common.PaginationRequest, classId, lectureId string, isGraded *bool, studentId string) (objects.ClassParticipantWithPagination, *constants.ErrorResponse) {
	var result objects.ClassParticipantWithPagination

	tx, err := p.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	classData, errs := p.ClassRepo.GetDetail(ctx, tx, classId, "")
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	modelResult, paginationResult, errs := p.StudentClassRepo.GetClassParticipant(ctx, tx, paginationData, []string{classId}, lectureId, isGraded, studentId)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	studentIds := []string{}
	for _, v := range modelResult {
		studentIds = append(studentIds, v.StudentId)
	}

	gradeData, errs := p.StudentClassRepo.GetStudentClassGradeByClassIdAndStudentIds(ctx, tx, classId, studentIds)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	result = objects.ClassParticipantWithPagination{
		Pagination: paginationResult,
		Data:       mapGetParticipant(classData.TotalLectureDone, modelResult, gradeData),
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, constants.ErrUnknown
	}

	return result, nil
}
