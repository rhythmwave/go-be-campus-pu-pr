package lecturer_mutation

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
	"github.com/sccicitb/pupr-backend/utils"
	appUtils "github.com/sccicitb/pupr-backend/utils"
)

type lecturerMutationService struct {
	*repository.RepoCtx
	*infra.InfraCtx
}

func (l lecturerMutationService) GetList(ctx context.Context, paginationData common.PaginationRequest, studyProgramId, idNationalLecturer, semesterId string) (objects.LecturerMutationListWithPagination, *constants.ErrorResponse) {
	var result objects.LecturerMutationListWithPagination

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

	modelResult, paginationResult, errs := l.LecturerMutationRepo.GetList(ctx, tx, paginationData, studyProgramId, idNationalLecturer, semesterId)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	resultData := []objects.GetLecturerMutation{}
	for _, v := range modelResult {
		resultData = append(resultData, objects.GetLecturerMutation{
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
			MutationDate:          v.MutationDate,
			DecisionNumber:        v.DecisionNumber,
			Destination:           v.Destination,
		})
	}

	result = objects.LecturerMutationListWithPagination{
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

func (l lecturerMutationService) Create(ctx context.Context, data objects.CreateLecturerMutation) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := l.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs = l.LecturerRepo.GetDetail(ctx, tx, data.LecturerId)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	createData := models.CreateLecturerMutation{
		LecturerId:     data.LecturerId,
		SemesterId:     data.SemesterId,
		MutationDate:   data.MutationDate,
		DecisionNumber: data.DecisionNumber,
		Destination:    data.Destination,
		CreatedBy:      claims.ID,
	}
	errs = l.LecturerMutationRepo.Create(ctx, tx, createData)
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
