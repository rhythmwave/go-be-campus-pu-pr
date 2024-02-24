package officer

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

type officerService struct {
	*repository.RepoCtx
	*infra.InfraCtx
}

func (o officerService) mapGetList(officerData []models.GetOfficer) ([]objects.GetOfficer, *constants.ErrorResponse) {
	results := []objects.GetOfficer{}

	for _, v := range officerData {
		var signatureUrl string
		var errs *constants.ErrorResponse
		if v.SignaturePath != nil && v.SignaturePathType != nil {
			signatureUrl, errs = o.Storage.GetURL(*v.SignaturePath, *v.SignaturePathType, nil)
			if errs != nil {
				return results, errs
			}
		}

		results = append(results, objects.GetOfficer{
			Id:                 v.Id,
			IdNationalLecturer: v.IdNationalLecturer,
			Name:               v.Name,
			Title:              v.Title,
			EnglishTitle:       v.EnglishTitle,
			StudyProgramId:     v.StudyProgramId,
			StudyProgramName:   v.StudyProgramName,
			SignaturePath:      v.SignaturePath,
			SignaturePathType:  v.SignaturePathType,
			SignatureUrl:       signatureUrl,
			ShowSignature:      v.ShowSignature,
			EmployeeNo:         v.EmployeeNo,
		})
	}

	return results, nil
}

func (a officerService) GetList(ctx context.Context, paginationData common.PaginationRequest) (objects.OfficerListWithPagination, *constants.ErrorResponse) {
	var result objects.OfficerListWithPagination

	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	modelResult, paginationResult, errs := a.OfficerRepo.GetList(ctx, tx, paginationData)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	resultData, errs := a.mapGetList(modelResult)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}
	result = objects.OfficerListWithPagination{
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

func (a officerService) Create(ctx context.Context, data objects.CreateOfficer) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	if data.StudyProgramId != "" {
		_, errs = a.StudyProgramRepo.GetDetail(ctx, tx, data.StudyProgramId, claims.Role, claims.ID)
		if errs != nil {
			_ = tx.Rollback()
			return errs
		}
	}

	createData := models.CreateOfficer{
		IdNationalLecturer: utils.NewNullString(data.IdNationalLecturer),
		Name:               data.Name,
		Title:              utils.NewNullString(data.Title),
		EnglishTitle:       utils.NewNullString(data.EnglishTitle),
		StudyProgramId:     utils.NewNullString(data.StudyProgramId),
		SignaturePath:      utils.NewNullString(data.SignaturePath),
		SignaturePathType:  utils.NewNullString(data.SignaturePathType),
		EmployeeNo:         utils.NewNullString(data.EmployeeNo),
		ShowSignature:      data.ShowSignature,
		CreatedBy:          claims.ID,
	}
	errs = a.OfficerRepo.Create(ctx, tx, createData)
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

func (a officerService) Update(ctx context.Context, data objects.UpdateOfficer) *constants.ErrorResponse {
	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs := a.OfficerRepo.GetDetail(ctx, tx, data.Id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	if data.StudyProgramId != "" {
		_, errs = a.StudyProgramRepo.GetDetail(ctx, tx, data.StudyProgramId, claims.Role, claims.ID)
		if errs != nil {
			_ = tx.Rollback()
			return errs
		}
	}

	updateData := models.UpdateOfficer{
		Id:                 data.Id,
		IdNationalLecturer: utils.NewNullString(data.IdNationalLecturer),
		Name:               data.Name,
		Title:              utils.NewNullString(data.Title),
		EnglishTitle:       utils.NewNullString(data.EnglishTitle),
		StudyProgramId:     utils.NewNullString(data.StudyProgramId),
		SignaturePath:      utils.NewNullString(data.SignaturePath),
		SignaturePathType:  utils.NewNullString(data.SignaturePathType),
		EmployeeNo:         utils.NewNullString(data.EmployeeNo),
		ShowSignature:      data.ShowSignature,
		UpdatedBy:          claims.ID,
	}
	errs = a.OfficerRepo.Update(ctx, tx, updateData)
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

func (a officerService) Delete(ctx context.Context, id string) *constants.ErrorResponse {
	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs := a.OfficerRepo.GetDetail(ctx, tx, id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	errs = a.OfficerRepo.Delete(ctx, tx, id)
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
