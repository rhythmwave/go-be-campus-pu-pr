package yudicium

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type YudiciumServiceInterface interface {
	Apply(ctx context.Context, data objects.ApplyYudicium) *constants.ErrorResponse
	GetListStudent(ctx context.Context, pagination common.PaginationRequest, req objects.GetListStudentYudiciumRequest) (objects.GetListStudentYudiciumWithPagination, *constants.ErrorResponse)
	Do(ctx context.Context, data objects.DoYudicium) *constants.ErrorResponse
}

func NewYudiciumService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) YudiciumServiceInterface {
	return &yudiciumService{
		repoCtx,
		infraCtx,
	}
}
