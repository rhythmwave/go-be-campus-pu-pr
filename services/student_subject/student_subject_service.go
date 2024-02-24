package student_subject

import (
	"context"
	"sort"

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

type studentSubjectService struct {
	*repository.RepoCtx
	*infra.InfraCtx
}

func mapGetPdfDetail(studentData models.GetStudent, studentSubjectData []models.GetStudentSubject) objects.GetTranscriptDetail {
	var result objects.GetTranscriptDetail

	subjectMap := make(map[uint32][]objects.GetTranscriptDetailSemesterSubject)
	var totalTheoryCredit uint32
	var totalPracticumCredit uint32
	for _, v := range studentSubjectData {
		theoryCredit := v.SubjectTheoryCredit
		practicumCredit := (v.SubjectPracticumCredit + v.SubjectFieldPracticumCredit)

		totalTheoryCredit += theoryCredit
		totalPracticumCredit += practicumCredit

		subjectMap[v.SemesterPackage] = append(subjectMap[v.SemesterPackage], objects.GetTranscriptDetailSemesterSubject{
			SubjectCode:        v.SubjectCode,
			SubjectName:        v.SubjectName,
			SubjectEnglishName: v.SubjectEnglishName,
			TheoryCredit:       theoryCredit,
			PracticumCredit:    practicumCredit,
			GradeCode:          v.GradeCode,
		})
	}

	semesterPackages := make([]int, 0, len(subjectMap))

	for k := range subjectMap {
		semesterPackages = append(semesterPackages, int(k))
	}
	sort.Ints(semesterPackages)

	var semesters []objects.GetTranscriptDetailSemester
	for _, v := range semesterPackages {
		semesters = append(semesters, objects.GetTranscriptDetailSemester{
			SemesterPackage: uint32(v),
			Subjects:        subjectMap[uint32(v)],
		})
	}

	result = objects.GetTranscriptDetail{
		NimNumber:           studentData.NimNumber,
		Name:                studentData.Name,
		BirthRegencyName:    studentData.BirthRegencyName,
		BirthDate:           studentData.BirthDate,
		GraduationDate:      studentData.GraduationDate,
		DiplomaNumber:       studentData.DiplomaNumber,
		StudyProgramName:    studentData.StudyProgramName,
		StudyLevelName:      studentData.StudyLevelName,
		StudyLevelShortName: studentData.StudyLevelShortName,
		TotalCredit:         totalTheoryCredit + totalPracticumCredit,
		Gpa:                 studentData.Gpa,
		GraduationPredicate: studentData.GraduationPredicate,
		TheoryCredit:        totalTheoryCredit,
		PracticumCredit:     totalPracticumCredit,
		ThesisTitle:         studentData.ThesisTitle,
		ThesisEnglishTitle:  studentData.ThesisEnglishTitle,
		Semesters:           semesters,
	}

	return result
}

func (s studentSubjectService) GetDetail(ctx context.Context, studentId string) (objects.GetDetailStudentSubject, *constants.ErrorResponse) {
	var result objects.GetDetailStudentSubject

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return result, errs
	}

	tx, err := s.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	if studentId == "" {
		studentId = claims.ID
	}

	studentData, errs := s.StudentRepo.GetDetail(ctx, tx, studentId, "")
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	studentSubjectData, _, errs := s.StudentRepo.GetStudentSubject(ctx, tx, common.PaginationRequest{Page: constants.DefaultPage, Limit: constants.DefaultUnlimited}, studentId)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	subjectResult := []objects.GetDetailStudentSubjectSubject{}
	for _, v := range studentSubjectData {
		subjectResult = append(subjectResult, objects.GetDetailStudentSubjectSubject{
			SemesterType:                v.GradeSemesterType,
			SemesterStartYear:           v.GradeSemesterStartYear,
			SemesterSchoolYear:          appUtils.GenerateSchoolYear(v.GradeSemesterStartYear),
			SubjectCode:                 v.SubjectCode,
			SubjectName:                 v.SubjectName,
			SubjectTheoryCredit:         v.SubjectTheoryCredit,
			SubjectPracticumCredit:      v.SubjectPracticumCredit,
			SubjectFieldPracticumCredit: v.SubjectFieldPracticumCredit,
			GradePoint:                  v.GradePoint,
			GradeCode:                   v.GradeCode,
		})
	}

	result = objects.GetDetailStudentSubject{
		Id:               studentData.Id,
		Name:             studentData.Name,
		NimNumber:        studentData.NimNumber,
		StudyProgramName: studentData.StudyProgramName,
		TotalCredit:      studentData.TotalCredit,
		Gpa:              studentData.Gpa,
		Subjects:         subjectResult,
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, constants.ErrUnknown
	}

	return result, nil
}

func (s studentSubjectService) GetPdfDetail(ctx context.Context, studentId string) (objects.GetTranscriptDetail, *constants.ErrorResponse) {
	var result objects.GetTranscriptDetail

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return result, errs
	}

	tx, err := s.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	if studentId == "" {
		studentId = claims.ID
	}

	studentData, errs := s.StudentRepo.GetDetail(ctx, tx, studentId, "")
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	studentSubjectData, _, errs := s.StudentRepo.GetStudentSubject(ctx, tx, common.PaginationRequest{Page: constants.DefaultPage, Limit: constants.DefaultUnlimited}, studentId)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	result = mapGetPdfDetail(studentData, studentSubjectData)

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, constants.ErrUnknown
	}

	return result, nil
}

func (s studentSubjectService) ConvertMbkmGrade(ctx context.Context, data objects.ConvertMbkmGrade) *constants.ErrorResponse {
	tx, err := s.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	activeSemesterData, errs := s.SemesterRepo.GetActive(ctx, tx)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	classData, errs := s.StudentClassRepo.GetStudentClassByStudentIdClassId(ctx, tx, []models.StudentIdClassId{
		{
			StudentId: data.StudentId,
			ClassId:   data.MbkmClassId,
		},
	})
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}
	if len(classData) == 0 {
		_ = tx.Rollback()
		return utils.ErrDataNotFound("class")
	}
	sourceClass := classData[0]
	if !sourceClass.SubjectIsMbkm {
		_ = tx.Rollback()
		return appConstants.ErrSubjectNotMbkm
	}
	if sourceClass.GradeCode == nil {
		_ = tx.Rollback()
		return appConstants.ErrClassNotGraded
	}

	availableCredit := sourceClass.TotalCredit - sourceClass.MbkmUsedCredit

	destinationSubjectData, errs := s.SubjectRepo.GetDetailByIds(ctx, tx, data.DestinationSubjectIds)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	var totalConvertedCredit uint32
	for _, v := range destinationSubjectData {
		totalConvertedCredit += (v.TheoryCredit + v.PracticumCredit + v.FieldPracticumCredit)
	}

	if availableCredit < totalConvertedCredit {
		_ = tx.Rollback()
		return appConstants.ErrInsufficientCreditConversion
	}

	errs = s.StudentClassRepo.UpdateMbkmConvertedCredit(ctx, tx, sourceClass.Id, totalConvertedCredit)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	convertData := models.ConvertStudentGrade{
		GradeSemesterId: activeSemesterData.Id,
		GradePoint:      sourceClass.GradePoint,
		GradeCode:       *sourceClass.GradeCode,
		MbkmSubjectId:   sourceClass.SubjectId,
	}
	errs = s.StudentRepo.ConvertGrade(ctx, tx, data.StudentId, data.DestinationSubjectIds, convertData)
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
