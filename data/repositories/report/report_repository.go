package report

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/sccicitb/pupr-backend/constants"
	appConstants "github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
)

type reportRepository struct {
	*db.DB
}

func (r reportRepository) GetActiveSemesterStudentStatus(ctx context.Context, tx *sqlx.Tx, activeSemesterId string, studyProgramIds []string) ([]models.GetReportStudentStatus, *constants.ErrorResponse) {
	var result []models.GetReportStudentStatus

	err := tx.SelectContext(
		ctx,
		&result,
		getActiveSemesterStudentStatusQuery,
		pq.Array(appConstants.PersistentStudentStatus()),
		pq.Array(appConstants.MomentaryStudentStatus()),
		activeSemesterId,
		pq.Array(studyProgramIds),
	)
	if err != nil {
		return result, constants.ErrorInternalServer(err.Error())
	}

	return result, nil
}

func (r reportRepository) GetStudentStatus(ctx context.Context, tx *sqlx.Tx, semesterId string, studyProgramIds []string) ([]models.GetReportStudentStatus, *constants.ErrorResponse) {
	var result []models.GetReportStudentStatus

	err := tx.SelectContext(
		ctx,
		&result,
		getStudentStatusQuery,
		semesterId,
		pq.Array(studyProgramIds),
	)
	if err != nil {
		return result, constants.ErrorInternalServer(err.Error())
	}

	return result, nil
}

func (r reportRepository) GetStudentClassGrade(ctx context.Context, tx *sqlx.Tx, semesterId string, subjectIds []string) ([]models.GetReportStudentClassGrade, *constants.ErrorResponse) {
	var result []models.GetReportStudentClassGrade

	err := tx.SelectContext(
		ctx,
		&result,
		getStudentClassGradeQuery,
		semesterId,
		pq.Array(subjectIds),
	)
	if err != nil {
		return result, constants.ErrorInternalServer(err.Error())
	}

	return result, nil
}

func (r reportRepository) GetStudentProvince(ctx context.Context, tx *sqlx.Tx, studyProgramIds []string, studentForceFrom, studentForceTo uint32) ([]models.GetReportStudentProvince, *constants.ErrorResponse) {
	var result []models.GetReportStudentProvince

	err := tx.SelectContext(
		ctx,
		&result,
		getStudentProvinceQuery,
		pq.Array(studyProgramIds),
		studentForceFrom,
		studentForceTo,
	)
	if err != nil {
		return result, constants.ErrorInternalServer(err.Error())
	}

	return result, nil
}

func (r reportRepository) GetStudentSchoolProvince(ctx context.Context, tx *sqlx.Tx, studyProgramIds []string, studentForceFrom, studentForceTo uint32) ([]models.GetReportStudentSchoolProvince, *constants.ErrorResponse) {
	var result []models.GetReportStudentSchoolProvince

	err := tx.SelectContext(
		ctx,
		&result,
		getStudentSchoolProvinceQuery,
		pq.Array(studyProgramIds),
		studentForceFrom,
		studentForceTo,
	)
	if err != nil {
		return result, constants.ErrorInternalServer(err.Error())
	}

	return result, nil
}
