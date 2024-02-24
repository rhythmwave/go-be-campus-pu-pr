package graduation_predicate

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type GraduationPredicateServiceInterface interface {
	GetList(ctx context.Context, paginationData common.PaginationRequest) (objects.GraduationPredicateListWithPagination, *constants.ErrorResponse)
	Create(ctx context.Context, data objects.CreateGraduationPredicate) *constants.ErrorResponse
	Update(ctx context.Context, data objects.UpdateGraduationPredicate) *constants.ErrorResponse
	Delete(ctx context.Context, id string) *constants.ErrorResponse
}

func NewGraduationPredicateService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) GraduationPredicateServiceInterface {
	return &graduationPredicateService{
		repoCtx,
		infraCtx,
	}
}
