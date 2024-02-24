package report

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
)

type ReportRepositoryInterface interface {
	GetActiveSemesterStudentStatus(ctx context.Context, tx *sqlx.Tx, activeSemesterId string, studyProgramIds []string) ([]models.GetReportStudentStatus, *constants.ErrorResponse)
	GetStudentStatus(ctx context.Context, tx *sqlx.Tx, semesterId string, studyProgramIds []string) ([]models.GetReportStudentStatus, *constants.ErrorResponse)
	GetStudentClassGrade(ctx context.Context, tx *sqlx.Tx, semesterId string, subjectIds []string) ([]models.GetReportStudentClassGrade, *constants.ErrorResponse)
	GetStudentProvince(ctx context.Context, tx *sqlx.Tx, studyProgramIds []string, studentForceFrom, studentForceTo uint32) ([]models.GetReportStudentProvince, *constants.ErrorResponse)
	GetStudentSchoolProvince(ctx context.Context, tx *sqlx.Tx, studyProgramIds []string, studentForceFrom, studentForceTo uint32) ([]models.GetReportStudentSchoolProvince, *constants.ErrorResponse)
}

func NewReportRepository(db *db.DB) ReportRepositoryInterface {
	return &reportRepository{
		db,
	}
}
