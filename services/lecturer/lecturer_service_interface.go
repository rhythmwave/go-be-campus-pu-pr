package lecturer

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type LecturerServiceInterface interface {
	GetList(ctx context.Context, paginationData common.PaginationRequest, req objects.GetLecturerRequest) (objects.LecturerListWithPagination, *constants.ErrorResponse)
	GetSchedule(ctx context.Context, paginationData common.PaginationRequest, studyProgramId, idNationalLecturer, semesterId string) (objects.LecturerScheduleWithPagination, *constants.ErrorResponse)
	GetDetail(ctx context.Context, id string) (objects.GetLecturerDetail, *constants.ErrorResponse)
	Create(ctx context.Context, data objects.CreateLecturer) *constants.ErrorResponse
	Update(ctx context.Context, data objects.UpdateLecturer) *constants.ErrorResponse
	Delete(ctx context.Context, id string) *constants.ErrorResponse
	GetSemesterSummary(ctx context.Context) (objects.GetLecturerSemesterSummary, *constants.ErrorResponse)
	GetProfile(ctx context.Context) (objects.GetLecturerProfile, *constants.ErrorResponse)
	GetAssignedClass(ctx context.Context, semesterId, lecturerId string, classIsActive *bool) ([]objects.GetLecturerAssignedClass, *constants.ErrorResponse)
}

func NewLecturerService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) LecturerServiceInterface {
	return &lecturerService{
		repoCtx,
		infraCtx,
	}
}
