package report

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type ReportServiceInterface interface {
	StudentStatus(ctx context.Context, semesterId string) ([]objects.ReportStudentStatus, *constants.ErrorResponse)
	StudentClassGrade(ctx context.Context, pagination common.PaginationRequest, semesterId, studyProgramId string) (objects.ReportStudentClassGradeWithPagination, *constants.ErrorResponse)
	StudentProvince(ctx context.Context, studyProgramId string, studentForceFrom, studentForceTo uint32) ([]objects.ReportStudentProvince, *constants.ErrorResponse)
	StudentSchoolProvince(ctx context.Context, studyProgramId string, studentForceFrom, studentForceTo uint32) ([]objects.ReportStudentSchoolProvince, *constants.ErrorResponse)
}

func NewReportService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) ReportServiceInterface {
	return &reportService{
		repoCtx,
		infraCtx,
	}
}
