package exam_supervisor

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

type examSupervisorService struct {
	*repository.RepoCtx
	*infra.InfraCtx
}

func (p examSupervisorService) GetList(ctx context.Context, paginationData common.PaginationRequest, studyProgramId string) (objects.ExamSupervisorListWithPagination, *constants.ErrorResponse) {
	var result objects.ExamSupervisorListWithPagination

	tx, err := p.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	modelResult, paginationResult, errs := p.ExamSupervisorRepo.GetList(ctx, tx, paginationData, studyProgramId)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	resultData := []objects.GetExamSupervisor{}
	for _, v := range modelResult {
		resultData = append(resultData, objects.GetExamSupervisor{
			Id:                 v.Id,
			Name:               v.Name,
			IdNationalLecturer: v.IdNationalLecturer,
		})
	}

	result = objects.ExamSupervisorListWithPagination{
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

func (f examSupervisorService) GetDetail(ctx context.Context, id string) (objects.GetExamSupervisorDetail, *constants.ErrorResponse) {
	var result objects.GetExamSupervisorDetail

	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	resultData, errs := f.ExamSupervisorRepo.GetDetail(ctx, tx, id)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	result = objects.GetExamSupervisorDetail{
		Id:                 resultData.Id,
		IdNationalLecturer: resultData.IdNationalLecturer,
		Name:               resultData.Name,
		FrontTitle:         resultData.FrontTitle,
		BackDegree:         resultData.BackDegree,
		StudyProgramId:     resultData.StudyProgramId,
		StudyProgramName:   resultData.StudyProgramName,
		IdNumber:           resultData.IdNumber,
		BirthDate:          resultData.BirthDate,
		BirthRegencyId:     resultData.BirthRegencyId,
		BirthRegencyName:   resultData.BirthRegencyName,
		BirthCountryId:     resultData.BirthCountryId,
		BirthCountryName:   resultData.BirthCountryName,
		Sex:                resultData.Sex,
		BloodType:          resultData.BloodType,
		Religion:           resultData.Religion,
		MaritalStatus:      resultData.MaritalStatus,
		Address:            resultData.Address,
		RegencyId:          resultData.RegencyId,
		RegencyName:        resultData.RegencyName,
		CountryId:          resultData.CountryId,
		CountryName:        resultData.CountryName,
		PostalCode:         resultData.PostalCode,
		PhoneNumber:        resultData.PhoneNumber,
		Fax:                resultData.Fax,
		MobilePhoneNumber:  resultData.MobilePhoneNumber,
		OfficePhoneNumber:  resultData.OfficePhoneNumber,
		EmployeeType:       resultData.EmployeeType,
		EmployeeStatus:     resultData.EmployeeStatus,
		SkCpnsNumber:       resultData.SkCpnsNumber,
		SkCpnsDate:         resultData.SkCpnsDate,
		TmtCpnsDate:        resultData.TmtCpnsDate,
		CpnsCategory:       resultData.CpnsCategory,
		CpnsDurationMonth:  resultData.CpnsDurationMonth,
		PrePositionDate:    resultData.PrePositionDate,
		SkPnsNumber:        resultData.SkPnsNumber,
		SkPnsDate:          resultData.SkPnsDate,
		TmtPnsDate:         resultData.TmtPnsDate,
		PnsCategory:        resultData.PnsCategory,
		PnsOathDate:        resultData.PnsOathDate,
		JoinDate:           resultData.JoinDate,
		EndDate:            resultData.EndDate,
		TaspenNumber:       resultData.TaspenNumber,
		FormerInstance:     resultData.FormerInstance,
		Remarks:            resultData.Remarks,
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, constants.ErrUnknown
	}

	return result, nil
}

func (a examSupervisorService) Create(ctx context.Context, data objects.CreateExamSupervisor) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	createData := models.CreateExamSupervisor{
		IdNationalLecturer: data.IdNationalLecturer,
		Name:               data.Name,
		FrontTitle:         utils.NewNullString(data.FrontTitle),
		BackDegree:         utils.NewNullString(data.BackDegree),
		StudyProgramId:     utils.NewNullString(data.StudyProgramId),
		IdNumber:           utils.NewNullString(data.IdNumber),
		BirthDate:          utils.NewNullTime(data.BirthDate),
		BirthRegencyId:     utils.NewNullInt32(int32(data.BirthRegencyId)),
		Sex:                utils.NewNullString(data.Sex),
		BloodType:          utils.NewNullString(data.BloodType),
		Religion:           utils.NewNullString(data.Religion),
		MaritalStatus:      utils.NewNullString(data.MaritalStatus),
		Address:            utils.NewNullString(data.Address),
		RegencyId:          utils.NewNullInt32(int32(data.RegencyId)),
		PostalCode:         utils.NewNullString(data.PostalCode),
		PhoneNumber:        utils.NewNullString(data.PhoneNumber),
		Fax:                utils.NewNullString(data.Fax),
		MobilePhoneNumber:  utils.NewNullString(data.MobilePhoneNumber),
		OfficePhoneNumber:  utils.NewNullString(data.OfficePhoneNumber),
		EmployeeType:       utils.NewNullString(data.EmployeeType),
		EmployeeStatus:     utils.NewNullString(data.EmployeeStatus),
		SkCpnsNumber:       utils.NewNullString(data.SkCpnsNumber),
		SkCpnsDate:         utils.NewNullTime(data.SkCpnsDate),
		TmtCpnsDate:        utils.NewNullTime(data.TmtCpnsDate),
		CpnsCategory:       utils.NewNullString(data.CpnsCategory),
		CpnsDurationMonth:  utils.NewNullInt32(int32(data.CpnsDurationMonth)),
		PrePositionDate:    utils.NewNullTime(data.PrePositionDate),
		SkPnsNumber:        utils.NewNullString(data.SkPnsNumber),
		SkPnsDate:          utils.NewNullTime(data.SkPnsDate),
		TmtPnsDate:         utils.NewNullTime(data.TmtPnsDate),
		PnsCategory:        utils.NewNullString(data.PnsCategory),
		PnsOathDate:        utils.NewNullTime(data.PnsOathDate),
		JoinDate:           utils.NewNullTime(data.JoinDate),
		EndDate:            utils.NewNullTime(data.EndDate),
		TaspenNumber:       utils.NewNullString(data.TaspenNumber),
		FormerInstance:     utils.NewNullString(data.FormerInstance),
		Remarks:            utils.NewNullString(data.Remarks),
		CreatedBy:          claims.ID,
	}
	errs = a.ExamSupervisorRepo.Create(ctx, tx, createData)
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

func (a examSupervisorService) Update(ctx context.Context, data objects.UpdateExamSupervisor) *constants.ErrorResponse {
	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs := a.ExamSupervisorRepo.GetDetail(ctx, tx, data.Id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	updateData := models.UpdateExamSupervisor{
		Id:                 data.Id,
		IdNationalLecturer: data.IdNationalLecturer,
		Name:               data.Name,
		FrontTitle:         utils.NewNullString(data.FrontTitle),
		BackDegree:         utils.NewNullString(data.BackDegree),
		StudyProgramId:     utils.NewNullString(data.StudyProgramId),
		IdNumber:           utils.NewNullString(data.IdNumber),
		BirthDate:          utils.NewNullTime(data.BirthDate),
		BirthRegencyId:     utils.NewNullInt32(int32(data.BirthRegencyId)),
		Sex:                utils.NewNullString(data.Sex),
		BloodType:          utils.NewNullString(data.BloodType),
		Religion:           utils.NewNullString(data.Religion),
		MaritalStatus:      utils.NewNullString(data.MaritalStatus),
		Address:            utils.NewNullString(data.Address),
		RegencyId:          utils.NewNullInt32(int32(data.RegencyId)),
		PostalCode:         utils.NewNullString(data.PostalCode),
		PhoneNumber:        utils.NewNullString(data.PhoneNumber),
		Fax:                utils.NewNullString(data.Fax),
		MobilePhoneNumber:  utils.NewNullString(data.MobilePhoneNumber),
		OfficePhoneNumber:  utils.NewNullString(data.OfficePhoneNumber),
		EmployeeType:       utils.NewNullString(data.EmployeeType),
		EmployeeStatus:     utils.NewNullString(data.EmployeeStatus),
		SkCpnsNumber:       utils.NewNullString(data.SkCpnsNumber),
		SkCpnsDate:         utils.NewNullTime(data.SkCpnsDate),
		TmtCpnsDate:        utils.NewNullTime(data.TmtCpnsDate),
		CpnsCategory:       utils.NewNullString(data.CpnsCategory),
		CpnsDurationMonth:  utils.NewNullInt32(int32(data.CpnsDurationMonth)),
		PrePositionDate:    utils.NewNullTime(data.PrePositionDate),
		SkPnsNumber:        utils.NewNullString(data.SkPnsNumber),
		SkPnsDate:          utils.NewNullTime(data.SkPnsDate),
		TmtPnsDate:         utils.NewNullTime(data.TmtPnsDate),
		PnsCategory:        utils.NewNullString(data.PnsCategory),
		PnsOathDate:        utils.NewNullTime(data.PnsOathDate),
		JoinDate:           utils.NewNullTime(data.JoinDate),
		EndDate:            utils.NewNullTime(data.EndDate),
		TaspenNumber:       utils.NewNullString(data.TaspenNumber),
		FormerInstance:     utils.NewNullString(data.FormerInstance),
		Remarks:            utils.NewNullString(data.Remarks),
		UpdatedBy:          claims.ID,
	}
	errs = a.ExamSupervisorRepo.Update(ctx, tx, updateData)
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

func (a examSupervisorService) Delete(ctx context.Context, id string) *constants.ErrorResponse {
	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs := a.ExamSupervisorRepo.GetDetail(ctx, tx, id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	errs = a.ExamSupervisorRepo.Delete(ctx, tx, id)
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
