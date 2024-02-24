package report

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

type reportService struct {
	*repository.RepoCtx
	*infra.InfraCtx
}

func (r reportService) StudentStatus(ctx context.Context, semesterId string) ([]objects.ReportStudentStatus, *constants.ErrorResponse) {
	var result []objects.ReportStudentStatus

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return result, errs
	}

	tx, err := r.DB.Begin(ctx)
	if err != nil {
		_ = tx.Rollback()
		return result, constants.ErrUnknown
	}

	studyProgramData, _, errs := r.StudyProgramRepo.GetList(ctx, tx, common.PaginationRequest{Limit: constants.DefaultUnlimited}, "", claims.Role, claims.ID)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	var studyProgramIds []string
	for _, v := range studyProgramData {
		studyProgramIds = append(studyProgramIds, v.Id)
	}

	activeSemesterData, errs := r.SemesterRepo.GetActive(ctx, tx)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	var resultData []models.GetReportStudentStatus
	if semesterId == "" || activeSemesterData.Id == semesterId {
		resultData, errs = r.ReportRepo.GetActiveSemesterStudentStatus(ctx, tx, activeSemesterData.Id, studyProgramIds)
		if errs != nil {
			_ = tx.Rollback()
			return result, errs
		}
	} else {
		resultData, errs = r.ReportRepo.GetStudentStatus(ctx, tx, semesterId, studyProgramIds)
		if errs != nil {
			_ = tx.Rollback()
			return result, errs
		}
	}

	result = mapStudentStatus(resultData)

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, constants.ErrUnknown
	}

	return result, nil
}

func (r reportService) StudentClassGrade(ctx context.Context, pagination common.PaginationRequest, semesterId, studyProgramId string) (objects.ReportStudentClassGradeWithPagination, *constants.ErrorResponse) {
	var result objects.ReportStudentClassGradeWithPagination

	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	modelResult, paginationResult, errs := r.SubjectRepo.GetList(ctx, tx, pagination, objects.GetSubjectRequest{StudyProgramId: studyProgramId})
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	var subjectIds []string
	for _, v := range modelResult {
		subjectIds = append(subjectIds, v.Id)
	}

	var gradeData []models.GetReportStudentClassGrade
	if len(subjectIds) != 0 {
		gradeData, errs = r.ReportRepo.GetStudentClassGrade(ctx, tx, semesterId, subjectIds)
		if errs != nil {
			_ = tx.Rollback()
			return result, errs
		}
	}

	result = objects.ReportStudentClassGradeWithPagination{
		Pagination: paginationResult,
		Data:       mapStudentClassGrade(modelResult, gradeData),
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, constants.ErrUnknown
	}

	return result, nil
}

func (r reportService) StudentProvince(ctx context.Context, studyProgramId string, studentForceFrom, studentForceTo uint32) ([]objects.ReportStudentProvince, *constants.ErrorResponse) {
	var result []objects.ReportStudentProvince

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return result, errs
	}

	tx, err := r.DB.Begin(ctx)
	if err != nil {
		_ = tx.Rollback()
		return result, constants.ErrUnknown
	}

	var studyProgramIds []string
	if studyProgramId != "" {
		studyProgramIds = append(studyProgramIds, studyProgramId)
	} else {
		studyProgramData, _, errs := r.StudyProgramRepo.GetList(ctx, tx, common.PaginationRequest{Limit: constants.DefaultUnlimited}, "", claims.Role, claims.ID)
		if errs != nil {
			_ = tx.Rollback()
			return result, errs
		}

		for _, v := range studyProgramData {
			studyProgramIds = append(studyProgramIds, v.Id)
		}
	}

	resultData, errs := r.ReportRepo.GetStudentProvince(ctx, tx, studyProgramIds, studentForceFrom, studentForceTo)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	result = mapStudentProvince(resultData)

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, constants.ErrUnknown
	}

	return result, nil
}

func (r reportService) StudentSchoolProvince(ctx context.Context, studyProgramId string, studentForceFrom, studentForceTo uint32) ([]objects.ReportStudentSchoolProvince, *constants.ErrorResponse) {
	var result []objects.ReportStudentSchoolProvince

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return result, errs
	}

	tx, err := r.DB.Begin(ctx)
	if err != nil {
		_ = tx.Rollback()
		return result, constants.ErrUnknown
	}

	var studyProgramIds []string
	if studyProgramId != "" {
		studyProgramIds = append(studyProgramIds, studyProgramId)
	} else {
		studyProgramData, _, errs := r.StudyProgramRepo.GetList(ctx, tx, common.PaginationRequest{Limit: constants.DefaultUnlimited}, "", claims.Role, claims.ID)
		if errs != nil {
			_ = tx.Rollback()
			return result, errs
		}

		for _, v := range studyProgramData {
			studyProgramIds = append(studyProgramIds, v.Id)
		}
	}

	resultData, errs := r.ReportRepo.GetStudentSchoolProvince(ctx, tx, studyProgramIds, studentForceFrom, studentForceTo)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	result = mapStudentSchoolProvince(resultData)

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, constants.ErrUnknown
	}

	return result, nil
}
