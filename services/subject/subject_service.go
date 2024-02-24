package subject

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

type subjectService struct {
	*repository.RepoCtx
	*infra.InfraCtx
}

func mapGetList(subjectData []models.GetSubject, classData []models.GetClass, filterSubjectWithClassOnly bool) []objects.GetSubject {
	results := []objects.GetSubject{}

	classMap := make(map[string][]objects.GetSubjectClass)
	for _, v := range classData {
		classMap[v.SubjectId] = append(classMap[v.SubjectId], objects.GetSubjectClass{
			Id:   v.Id,
			Name: v.Name,
		})
	}

	for _, v := range subjectData {
		if !filterSubjectWithClassOnly || (filterSubjectWithClassOnly && len(classMap[v.Id]) != 0) {
			results = append(results, objects.GetSubject{
				Id:                            v.Id,
				StudyProgramId:                v.StudyProgramId,
				StudyProgramName:              v.StudyProgramName,
				CurriculumId:                  v.CurriculumId,
				CurriculumName:                v.CurriculumName,
				Code:                          v.Code,
				Name:                          v.Name,
				IsMandatory:                   v.IsMandatory,
				SemesterPackage:               v.SemesterPackage,
				TheoryCredit:                  v.TheoryCredit,
				PracticumCredit:               v.PracticumCredit,
				FieldPracticumCredit:          v.FieldPracticumCredit,
				SubjectPrerequisiteId:         v.SubjectPrerequisiteId,
				PrerequisiteType:              v.PrerequisiteType,
				PrerequisiteMinimumGradePoint: v.PrerequisiteMinimumGradePoint,
				EquivalentStudyProgramId:      v.EquivalentStudyProgramId,
				EquivalentStudyProgramName:    v.EquivalentStudyProgramName,
				EquivalentCurriculumId:        v.EquivalentCurriculumId,
				EquivalentCurriculumName:      v.EquivalentCurriculumName,
				EquivalentSubjectId:           v.EquivalentSubjectId,
				EquivalentSubjectCode:         v.EquivalentSubjectCode,
				EquivalentSubjectName:         v.EquivalentSubjectName,
				SubjectCategoryId:             v.SubjectCategoryId,
				SubjectCategoryName:           v.SubjectCategoryName,
				IsThesis:                      v.IsThesis,
				IsMbkm:                        v.IsMbkm,
				TotalLessonPlan:               v.TotalLessonPlan,
				Classes:                       classMap[v.Id],
			})
		}
	}

	return results
}

func (a subjectService) GetList(ctx context.Context, paginationData common.PaginationRequest, req objects.GetSubjectRequest) (objects.SubjectListWithPagination, *constants.ErrorResponse) {
	var result objects.SubjectListWithPagination

	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	modelResult, paginationResult, errs := a.SubjectRepo.GetList(ctx, tx, paginationData, req)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	subjectIds := []string{}
	for _, v := range modelResult {
		subjectIds = append(subjectIds, v.Id)
	}

	classData := []models.GetClass{}
	if req.ClassSemesterId != "" {
		classData, errs = a.ClassRepo.GetBySubjectIdsSemesterId(ctx, tx, subjectIds, req.ClassSemesterId)
		if errs != nil {
			_ = tx.Rollback()
			return result, errs
		}
	}

	result = objects.SubjectListWithPagination{
		Pagination: paginationResult,
		Data:       mapGetList(modelResult, classData, req.ClassSemesterId != ""),
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, constants.ErrUnknown
	}

	return result, nil
}

func (f subjectService) GetDetail(ctx context.Context, id string) (objects.GetSubjectDetail, *constants.ErrorResponse) {
	var result objects.GetSubjectDetail

	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	resultData, errs := f.SubjectRepo.GetDetail(ctx, tx, id)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	var syllabusUrl string
	if resultData.SyllabusPath != nil && resultData.SyllabusPathType != nil {
		syllabusUrl, errs = f.Storage.GetURL(*resultData.SyllabusPath, *resultData.SyllabusPathType, nil)
		if errs != nil {
			_ = tx.Rollback()
			return result, errs
		}
	}

	prerequisiteData, errs := f.SubjectPrerequisiteRepo.GetBySubjectId(ctx, tx, id)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	subjectPrerequisites := []objects.GetSubjectDetailPrerequisite{}
	for _, v := range prerequisiteData {
		subjectPrerequisites = append(subjectPrerequisites, objects.GetSubjectDetailPrerequisite{
			Id:                v.PrerequisiteSubjectId,
			Code:              v.PrerequisiteSubjectCode,
			Name:              v.PrerequisiteSubjectName,
			PrerequisiteType:  v.PrerequisiteType,
			MinimumGradePoint: v.MinimumGradePoint,
		})
	}

	result = objects.GetSubjectDetail{
		Id:                           resultData.Id,
		StudyProgramId:               resultData.StudyProgramId,
		StudyProgramName:             resultData.StudyProgramName,
		CurriculumId:                 resultData.CurriculumId,
		CurriculumName:               resultData.CurriculumName,
		Code:                         resultData.Code,
		Name:                         resultData.Name,
		ShortName:                    resultData.ShortName,
		EnglishName:                  resultData.EnglishName,
		EnglishShortName:             resultData.EnglishShortName,
		IsMandatory:                  resultData.IsMandatory,
		Trait:                        resultData.Trait,
		Type:                         resultData.Type,
		SubjectCategoryId:            resultData.SubjectCategoryId,
		SubjectCategoryName:          resultData.SubjectCategoryName,
		CurriculumType:               resultData.CurriculumType,
		TheoryCredit:                 resultData.TheoryCredit,
		PracticumCredit:              resultData.PracticumCredit,
		FieldPracticumCredit:         resultData.FieldPracticumCredit,
		SemesterPackage:              resultData.SemesterPackage,
		RepeatCourseLimit:            resultData.RepeatCourseLimit,
		IsActive:                     resultData.IsActive,
		HasLectureUnit:               resultData.HasLectureUnit,
		HasTeachingMaterial:          resultData.HasTeachingMaterial,
		HasLectureSummary:            resultData.HasLectureSummary,
		SupportingLecturerId:         resultData.SupportingLecturerId,
		SupportingLecturerName:       resultData.SupportingLecturerName,
		StartDate:                    resultData.StartDate,
		EndDate:                      resultData.EndDate,
		MinimumPassingGradePoint:     resultData.MinimumPassingGradePoint,
		MinimumMandatoryCreditTaken:  resultData.MinimumMandatoryCreditTaken,
		MinimumOptionalCreditTaken:   resultData.MinimumOptionalCreditTaken,
		MinimumTotalCreditTaken:      resultData.MinimumTotalCreditTaken,
		MinimumMandatoryCreditPassed: resultData.MinimumMandatoryCreditPassed,
		MinimumOptionalCreditPassed:  resultData.MinimumOptionalCreditPassed,
		MinimumTotalCreditPassed:     resultData.MinimumTotalCreditPassed,
		MinimumGpa:                   resultData.MinimumGpa,
		Abstraction:                  resultData.Abstraction,
		SyllabusPath:                 resultData.SyllabusPath,
		SyllabusPathType:             resultData.SyllabusPathType,
		IsThesis:                     resultData.IsThesis,
		IsMbkm:                       resultData.IsMbkm,
		SyllabusUrl:                  syllabusUrl,
		SubjectPrerequisites:         subjectPrerequisites,
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, constants.ErrUnknown
	}

	return result, nil
}

func (a subjectService) Create(ctx context.Context, data objects.CreateSubject) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	createData := models.CreateSubject{
		CurriculumId:                 data.CurriculumId,
		Code:                         data.Code,
		Name:                         data.Name,
		ShortName:                    utils.NewNullString(data.ShortName),
		EnglishName:                  utils.NewNullString(data.EnglishName),
		EnglishShortName:             utils.NewNullString(data.EnglishShortName),
		IsMandatory:                  data.IsMandatory,
		Trait:                        data.Trait,
		Type:                         utils.NewNullString(data.Type),
		SubjectCategoryId:            data.SubjectCategoryId,
		CurriculumType:               data.CurriculumType,
		TheoryCredit:                 data.TheoryCredit,
		PracticumCredit:              data.PracticumCredit,
		FieldPracticumCredit:         data.FieldPracticumCredit,
		SemesterPackage:              data.SemesterPackage,
		RepeatCourseLimit:            data.RepeatCourseLimit,
		IsActive:                     data.IsActive,
		HasLectureUnit:               data.HasLectureUnit,
		HasTeachingMaterial:          data.HasTeachingMaterial,
		HasLectureSummary:            data.HasLectureSummary,
		SupportingLecturerId:         utils.NewNullString(data.SupportingLecturerId),
		StartDate:                    utils.NewNullTime(data.StartDate),
		EndDate:                      utils.NewNullTime(data.EndDate),
		MinimumPassingGradePoint:     data.MinimumPassingGradePoint,
		MinimumMandatoryCreditTaken:  utils.NewNullInt32(int32(data.MinimumMandatoryCreditTaken)),
		MinimumOptionalCreditTaken:   utils.NewNullInt32(int32(data.MinimumOptionalCreditTaken)),
		MinimumTotalCreditTaken:      utils.NewNullInt32(int32(data.MinimumTotalCreditTaken)),
		MinimumMandatoryCreditPassed: utils.NewNullInt32(int32(data.MinimumMandatoryCreditPassed)),
		MinimumOptionalCreditPassed:  utils.NewNullInt32(int32(data.MinimumOptionalCreditPassed)),
		MinimumTotalCreditPassed:     utils.NewNullInt32(int32(data.MinimumTotalCreditPassed)),
		MinimumGpa:                   utils.NewNullFloat64(&data.MinimumGpa),
		Abstraction:                  utils.NewNullString(data.Abstraction),
		SyllabusPath:                 utils.NewNullString(data.SyllabusPath),
		SyllabusPathType:             utils.NewNullString(data.SyllabusPathType),
		IsThesis:                     utils.NewNullBoolean(data.IsThesis),
		IsMbkm:                       data.IsMbkm,
		CreatedBy:                    claims.ID,
	}
	errs = a.SubjectRepo.Create(ctx, tx, createData)
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

func (a subjectService) Update(ctx context.Context, data objects.UpdateSubject) *constants.ErrorResponse {
	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs := a.SubjectRepo.GetDetail(ctx, tx, data.Id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	updateData := models.UpdateSubject{
		Id:                           data.Id,
		CurriculumId:                 data.CurriculumId,
		Code:                         data.Code,
		Name:                         data.Name,
		ShortName:                    utils.NewNullString(data.ShortName),
		EnglishName:                  utils.NewNullString(data.EnglishName),
		EnglishShortName:             utils.NewNullString(data.EnglishShortName),
		IsMandatory:                  data.IsMandatory,
		Trait:                        data.Trait,
		Type:                         utils.NewNullString(data.Type),
		SubjectCategoryId:            data.SubjectCategoryId,
		CurriculumType:               data.CurriculumType,
		TheoryCredit:                 data.TheoryCredit,
		PracticumCredit:              data.PracticumCredit,
		FieldPracticumCredit:         data.FieldPracticumCredit,
		SemesterPackage:              data.SemesterPackage,
		RepeatCourseLimit:            data.RepeatCourseLimit,
		IsActive:                     data.IsActive,
		HasLectureUnit:               data.HasLectureUnit,
		HasTeachingMaterial:          data.HasTeachingMaterial,
		HasLectureSummary:            data.HasLectureSummary,
		SupportingLecturerId:         utils.NewNullString(data.SupportingLecturerId),
		StartDate:                    utils.NewNullTime(data.StartDate),
		EndDate:                      utils.NewNullTime(data.EndDate),
		MinimumPassingGradePoint:     data.MinimumPassingGradePoint,
		MinimumMandatoryCreditTaken:  utils.NewNullInt32(int32(data.MinimumMandatoryCreditTaken)),
		MinimumOptionalCreditTaken:   utils.NewNullInt32(int32(data.MinimumOptionalCreditTaken)),
		MinimumTotalCreditTaken:      utils.NewNullInt32(int32(data.MinimumTotalCreditTaken)),
		MinimumMandatoryCreditPassed: utils.NewNullInt32(int32(data.MinimumMandatoryCreditPassed)),
		MinimumOptionalCreditPassed:  utils.NewNullInt32(int32(data.MinimumOptionalCreditPassed)),
		MinimumTotalCreditPassed:     utils.NewNullInt32(int32(data.MinimumTotalCreditPassed)),
		MinimumGpa:                   utils.NewNullFloat64(&data.MinimumGpa),
		Abstraction:                  utils.NewNullString(data.Abstraction),
		SyllabusPath:                 utils.NewNullString(data.SyllabusPath),
		SyllabusPathType:             utils.NewNullString(data.SyllabusPathType),
		IsThesis:                     utils.NewNullBoolean(data.IsThesis),
		IsMbkm:                       data.IsMbkm,
		UpdatedBy:                    claims.ID,
	}
	errs = a.SubjectRepo.Update(ctx, tx, updateData)
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

func (a subjectService) Delete(ctx context.Context, id string) *constants.ErrorResponse {
	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs := a.SubjectRepo.GetDetail(ctx, tx, id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	errs = a.SubjectRepo.Delete(ctx, tx, id)
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

func (a subjectService) SetPrerequisiteSubject(ctx context.Context, subjectId string, prerequisites []objects.SetPrerequisiteSubject) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs = a.SubjectRepo.GetDetail(ctx, tx, subjectId)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	if len(prerequisites) == 0 {
		errs = a.SubjectPrerequisiteRepo.DeleteAllBySubjectId(ctx, tx, subjectId)
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

	prerequisiteSubjectIds := []string{}
	upsertData := []models.CreateSubjectPrerequisite{}
	for _, v := range prerequisites {
		prerequisiteSubjectIds = append(prerequisiteSubjectIds, v.Id)
		upsertData = append(upsertData, models.CreateSubjectPrerequisite{
			SubjectId:             subjectId,
			PrerequisiteSubjectId: v.Id,
			PrerequisiteType:      v.PrerequisiteType,
			MinimumGradePoint:     utils.NewNullFloat64(&v.MinimumGradePoint),
			CreatedBy:             claims.ID,
		})
	}

	errs = a.SubjectPrerequisiteRepo.DeleteAllBySubjectIdExcludingPrerequisiteSubjectId(ctx, tx, subjectId, prerequisiteSubjectIds)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	errs = a.SubjectPrerequisiteRepo.Upsert(ctx, tx, upsertData)
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

func (a subjectService) SetEquivalentSubject(ctx context.Context, data objects.SetEquivalentSubject) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	createData := models.CreateSubjectEquivalence{
		SubjectId:           data.SubjectId,
		EquivalentSubjectId: data.EquivalentSubjectId,
		CreatedBy:           claims.ID,
	}
	errs = a.SubjectEquivalenceRepo.Upsert(ctx, tx, createData)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}
	if data.IsViceVersa {
		viceVersaData := models.CreateSubjectEquivalence{
			SubjectId:           data.EquivalentSubjectId,
			EquivalentSubjectId: data.SubjectId,
			CreatedBy:           claims.ID,
		}
		errs = a.SubjectEquivalenceRepo.Upsert(ctx, tx, viceVersaData)
		if errs != nil {
			_ = tx.Rollback()
			return errs
		}
	} else {
		errs = a.SubjectEquivalenceRepo.DeleteBySubjectIdAndEquivalentSubjectId(ctx, tx, data.EquivalentSubjectId, data.SubjectId)
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

func (a subjectService) DeleteEquivalentSubject(ctx context.Context, subjectId, equivalentSubjectId string) *constants.ErrorResponse {
	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	errs := a.SubjectEquivalenceRepo.DeleteBySubjectIdAndEquivalentSubjectId(ctx, tx, subjectId, equivalentSubjectId)
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
