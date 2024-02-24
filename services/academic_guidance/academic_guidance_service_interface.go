package academic_guidance

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type AcademicGuidanceServiceInterface interface {
	GetListStudent(ctx context.Context, paginationData common.PaginationRequest, lecturerId, semesterId string) (objects.AcademicGuidanceStudentListWithPagination, *constants.ErrorResponse)
	GetDetail(ctx context.Context, studentId, semesterId string) (objects.GetAcademicGuidanceDetail, *constants.ErrorResponse)
	Upsert(ctx context.Context, data objects.UpsertAcademicGuidance) *constants.ErrorResponse
	UpsertDecision(ctx context.Context, data objects.UpsertDecisionAcademicGuidance) *constants.ErrorResponse
	GetSessionList(ctx context.Context, academicGuidanceId, semesterId, lecturerId string) ([]objects.GetAcademicGuidanceSession, *constants.ErrorResponse)
	CreateSession(ctx context.Context, data objects.CreateAcademicGuidanceSession) *constants.ErrorResponse
	UpdateSession(ctx context.Context, data objects.UpdateAcademicGuidanceSession) *constants.ErrorResponse
	DeleteSession(ctx context.Context, id string) *constants.ErrorResponse
}

func NewAcademicGuidanceService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) AcademicGuidanceServiceInterface {
	return &academicGuidanceService{
		repoCtx,
		infraCtx,
	}
}
