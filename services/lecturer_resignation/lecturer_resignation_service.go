package lecturer_resignation

import (
	"context"

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

type lecturerResignationService struct {
	*repository.RepoCtx
	*infra.InfraCtx
}

func (l lecturerResignationService) GetList(ctx context.Context, paginationData common.PaginationRequest, studyProgramId, idNationalLecturer, semesterId string) (objects.LecturerResignationListWithPagination, *constants.ErrorResponse) {
	var result objects.LecturerResignationListWithPagination

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return result, errs
	}

	tx, err := l.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	if studyProgramId != "" {
		_, errs := l.StudyProgramRepo.GetDetail(ctx, tx, studyProgramId, claims.Role, claims.ID)
		if errs != nil {
			_ = tx.Rollback()
			return result, errs
		}
	}

	modelResult, paginationResult, errs := l.LecturerResignationRepo.GetList(ctx, tx, paginationData, studyProgramId, idNationalLecturer, semesterId)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	resultData := []objects.GetLecturerResignation{}
	for _, v := range modelResult {
		resultData = append(resultData, objects.GetLecturerResignation{
			Id:                    v.Id,
			Name:                  v.Name,
			IdNationalLecturer:    v.IdNationalLecturer,
			FrontTitle:            v.FrontTitle,
			BackDegree:            v.BackDegree,
			SemesterSchoolYear:    appUtils.GenerateSchoolYear(v.SemesterStartYear),
			SemesterType:          v.SemesterType,
			DiktiStudyProgramCode: v.DiktiStudyProgramCode,
			StudyProgramName:      v.StudyProgramName,
			StudyLevelShortName:   v.StudyLevelShortName,
			DiktiStudyProgramType: v.DiktiStudyProgramType,
			ResignDate:            v.ResignDate,
			ResignationNumber:     v.ResignationNumber,
			Purpose:               v.Purpose,
			Remarks:               v.Remarks,
		})
	}

	result = objects.LecturerResignationListWithPagination{
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

func (l lecturerResignationService) Create(ctx context.Context, data objects.CreateLecturerResignation) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := l.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	if !utils.InArrayExist(data.Purpose, appConstants.LecturerResignStatus()) {
		_ = tx.Rollback()
		return appConstants.ErrInvalidResignPurpose
	}

	lecturerData, errs := l.LecturerRepo.GetDetail(ctx, tx, data.LecturerId)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}
	if utils.InArrayExist(utils.NullStringScan(lecturerData.Status), appConstants.LecturerResignStatus()) {
		_ = tx.Rollback()
		return appConstants.ErrLecturerIsResigned
	}

	createData := models.CreateLecturerResignation{
		LecturerId:        data.LecturerId,
		SemesterId:        data.SemesterId,
		ResignDate:        data.ResignDate,
		ResignationNumber: data.ResignationNumber,
		Purpose:           data.Purpose,
		Remarks:           data.Remarks,
		CreatedBy:         claims.ID,
	}
	errs = l.LecturerResignationRepo.Create(ctx, tx, createData)
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
