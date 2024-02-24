package study_program

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

type studyProgramService struct {
	*repository.RepoCtx
	*infra.InfraCtx
}

func (f studyProgramService) GetList(ctx context.Context, paginationData common.PaginationRequest, majorId string, withAccessToken bool) (objects.StudyProgramListWithPagination, *constants.ErrorResponse) {
	var result objects.StudyProgramListWithPagination

	var userId string
	var userRole string
	if withAccessToken {
		claims, errs := utils.GetJWTClaimsFromContext(ctx)
		if errs != nil {
			return result, errs
		}
		userId = claims.ID
		userRole = claims.Role
	}

	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	modelResult, paginationResult, errs := f.StudyProgramRepo.GetList(ctx, tx, paginationData, majorId, userRole, userId)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	resultData := []objects.GetStudyProgram{}
	for _, v := range modelResult {
		resultData = append(resultData, objects.GetStudyProgram{
			Id:                    v.Id,
			Name:                  v.Name,
			StudyLevelShortName:   v.StudyLevelShortName,
			StudyLevelName:        v.StudyLevelName,
			DiktiStudyProgramType: v.DiktiStudyProgramType,
			DiktiStudyProgramCode: v.DiktiStudyProgramCode,
			Accreditation:         v.Accreditation,
			ActiveCurriculumYear:  v.ActiveCurriculumYear,
			Degree:                v.Degree,
			ShortDegree:           v.ShortDegree,
			EnglishDegree:         v.EnglishDegree,
		})
	}

	result = objects.StudyProgramListWithPagination{
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

func (f studyProgramService) GetDetail(ctx context.Context, id string) (objects.GetStudyProgramDetail, *constants.ErrorResponse) {
	var result objects.GetStudyProgramDetail

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return result, errs
	}

	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	resultData, errs := f.StudyProgramRepo.GetDetail(ctx, tx, id, claims.Role, claims.ID)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	result = objects.GetStudyProgramDetail{
		Id:                            resultData.Id,
		DiktiStudyProgramId:           resultData.DiktiStudyProgramId,
		DiktiStudyProgramCode:         resultData.DiktiStudyProgramCode,
		DiktiStudyProgramName:         resultData.DiktiStudyProgramName,
		DiktiStudyProgramType:         resultData.DiktiStudyProgramType,
		StudyLevelShortName:           resultData.StudyLevelShortName,
		StudyLevelName:                resultData.StudyLevelName,
		Name:                          resultData.Name,
		EnglishName:                   resultData.EnglishName,
		ShortName:                     resultData.ShortName,
		EnglishShortName:              resultData.EnglishShortName,
		AdministrativeUnit:            resultData.AdministrativeUnit,
		FacultyId:                     resultData.FacultyId,
		FacultyName:                   resultData.FacultyName,
		MajorId:                       resultData.MajorId,
		MajorName:                     resultData.MajorName,
		Address:                       resultData.Address,
		PhoneNumber:                   resultData.PhoneNumber,
		Fax:                           resultData.Fax,
		Email:                         resultData.Email,
		Website:                       resultData.Website,
		ContactPerson:                 resultData.ContactPerson,
		CuriculumReviewFrequency:      utils.NullStringScan(resultData.CuriculumReviewFrequency),
		CuriculumReviewMethod:         utils.NullStringScan(resultData.CuriculumReviewMethod),
		EstablishmentDate:             utils.NullTimeScan(resultData.EstablishmentDate),
		IsActive:                      resultData.IsActive,
		StartSemester:                 resultData.StartSemester,
		OperationalPermitNumber:       utils.NullStringScan(resultData.OperationalPermitNumber),
		OperationalPermitDate:         utils.NullTimeScan(resultData.OperationalPermitDate),
		OperationalPermitDueDate:      utils.NullTimeScan(resultData.OperationalPermitDueDate),
		HeadLecturerId:                resultData.HeadLecturerId,
		HeadLecturerName:              resultData.HeadLecturerName,
		HeadLecturerMobilePhoneNumber: resultData.HeadLecturerMobilePhoneNumber,
		OperatorName:                  resultData.OperatorName,
		OperatorPhoneNumber:           resultData.OperatorPhoneNumber,
		MinimumGraduationCredit:       resultData.MinimumGraduationCredit,
		MinimumThesisCredit:           resultData.MinimumThesisCredit,
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, constants.ErrUnknown
	}

	return result, nil
}

func (f studyProgramService) Create(ctx context.Context, data objects.CreateStudyProgram) *constants.ErrorResponse {
	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	createData := models.CreateStudyProgram{
		DiktiStudyProgramId:      data.DiktiStudyProgramId,
		Name:                     data.Name,
		EnglishName:              utils.NewNullString(data.EnglishName),
		ShortName:                utils.NewNullString(data.ShortName),
		EnglishShortName:         utils.NewNullString(data.EnglishShortName),
		AdministrativeUnit:       utils.NewNullString(data.AdministrativeUnit),
		MajorId:                  data.MajorId,
		Address:                  utils.NewNullString(data.Address),
		PhoneNumber:              utils.NewNullString(data.PhoneNumber),
		Fax:                      utils.NewNullString(data.Fax),
		Email:                    utils.NewNullString(data.Email),
		Website:                  utils.NewNullString(data.Website),
		ContactPerson:            utils.NewNullString(data.ContactPerson),
		CuriculumReviewFrequency: utils.NewNullString(data.CuriculumReviewFrequency),
		CuriculumReviewMethod:    utils.NewNullString(data.CuriculumReviewMethod),
		EstablishmentDate:        utils.NewNullTime(data.EstablishmentDate),
		IsActive:                 data.IsActive,
		StartSemester:            utils.NewNullString(data.StartSemester),
		OperationalPermitNumber:  utils.NewNullString(data.OperationalPermitNumber),
		OperationalPermitDate:    utils.NewNullTime(data.OperationalPermitDate),
		OperationalPermitDueDate: utils.NewNullTime(data.OperationalPermitDueDate),
		HeadLecturerId:           utils.NewNullString(data.HeadLecturerId),
		OperatorName:             utils.NewNullString(data.OperatorName),
		OperatorPhoneNumber:      utils.NewNullString(data.OperatorPhoneNumber),
		CreatedBy:                claims.ID,
	}
	errs = f.StudyProgramRepo.Create(ctx, tx, createData)
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

func (f studyProgramService) Update(ctx context.Context, data objects.UpdateStudyProgram) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs = f.StudyProgramRepo.GetDetail(ctx, tx, data.Id, claims.Role, claims.ID)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	updateData := models.UpdateStudyProgram{
		Id:                       data.Id,
		DiktiStudyProgramId:      data.DiktiStudyProgramId,
		Name:                     data.Name,
		EnglishName:              utils.NewNullString(data.EnglishName),
		ShortName:                utils.NewNullString(data.ShortName),
		EnglishShortName:         utils.NewNullString(data.EnglishShortName),
		AdministrativeUnit:       utils.NewNullString(data.AdministrativeUnit),
		MajorId:                  data.MajorId,
		Address:                  utils.NewNullString(data.Address),
		PhoneNumber:              utils.NewNullString(data.PhoneNumber),
		Fax:                      utils.NewNullString(data.Fax),
		Email:                    utils.NewNullString(data.Email),
		Website:                  utils.NewNullString(data.Website),
		ContactPerson:            utils.NewNullString(data.ContactPerson),
		CuriculumReviewFrequency: utils.NewNullString(data.CuriculumReviewFrequency),
		CuriculumReviewMethod:    utils.NewNullString(data.CuriculumReviewMethod),
		EstablishmentDate:        utils.NewNullTime(data.EstablishmentDate),
		IsActive:                 data.IsActive,
		StartSemester:            utils.NewNullString(data.StartSemester),
		OperationalPermitNumber:  utils.NewNullString(data.OperationalPermitNumber),
		OperationalPermitDate:    utils.NewNullTime(data.OperationalPermitDate),
		OperationalPermitDueDate: utils.NewNullTime(data.OperationalPermitDueDate),
		HeadLecturerId:           utils.NewNullString(data.HeadLecturerId),
		OperatorName:             utils.NewNullString(data.OperatorName),
		OperatorPhoneNumber:      utils.NewNullString(data.OperatorPhoneNumber),
		MinimumGraduationCredit:  utils.NewNullInt32(int32(data.MinimumGraduationCredit)),
		MinimumThesisCredit:      utils.NewNullInt32(int32(data.MinimumThesisCredit)),
		UpdatedBy:                claims.ID,
	}
	errs = f.StudyProgramRepo.Update(ctx, tx, updateData)
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

func (f studyProgramService) UpdateDegree(ctx context.Context, data objects.UpdateDegreeStudyProgram) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs = f.StudyProgramRepo.GetDetail(ctx, tx, data.StudyProgramId, claims.Role, claims.ID)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	updateData := models.UpdateDegreeStudyProgram{
		Id:            data.StudyProgramId,
		Degree:        utils.NewNullString(data.Degree),
		ShortDegree:   utils.NewNullString(data.ShortDegree),
		EnglishDegree: utils.NewNullString(data.EnglishDegree),
		UpdatedBy:     claims.ID,
	}
	errs = f.StudyProgramRepo.UpdateDegree(ctx, tx, updateData)
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

func (f studyProgramService) Delete(ctx context.Context, id string) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs = f.StudyProgramRepo.GetDetail(ctx, tx, id, claims.Role, claims.ID)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	errs = f.StudyProgramRepo.Delete(ctx, tx, id)
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
