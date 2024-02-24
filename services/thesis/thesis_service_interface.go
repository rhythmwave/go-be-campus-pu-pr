package thesis

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type ThesisServiceInterface interface {
	GetList(ctx context.Context, pagination common.PaginationRequest, studyProgramId string, nimNumber int64, startSemesterId, status, supervisorLecturerId string) (objects.GetListThesisWithPagination, *constants.ErrorResponse)
	GetDetail(ctx context.Context, id string) (objects.GetDetailThesis, *constants.ErrorResponse)
	Create(ctx context.Context, data objects.CreateThesis) *constants.ErrorResponse
	Update(ctx context.Context, data objects.UpdateThesis) *constants.ErrorResponse
	Delete(ctx context.Context, id string) *constants.ErrorResponse
	GetListThesisDefenseRequest(ctx context.Context, pagination common.PaginationRequest, studyProgramId string, nimNumber int64, startSemesterId string) (objects.GetListThesisDefenseRequestWithPagination, *constants.ErrorResponse)
	RegisterThesisDefense(ctx context.Context, studentId string) *constants.ErrorResponse
	CreateThesisDefense(ctx context.Context, data objects.CreateThesisDefense) *constants.ErrorResponse
	UpdateThesisDefense(ctx context.Context, data objects.UpdateThesisDefense) *constants.ErrorResponse
	GetThesisSupervisorLog(ctx context.Context, pagination common.PaginationRequest, idNationalLecturer, semesterId string) (objects.GetThesisSupervisorLogWithPagination, *constants.ErrorResponse)
}

func NewThesisService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) ThesisServiceInterface {
	return &thesisService{
		repoCtx,
		infraCtx,
	}
}
