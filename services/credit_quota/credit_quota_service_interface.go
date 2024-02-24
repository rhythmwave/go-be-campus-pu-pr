package credit_quota

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type CreditQuotaServiceInterface interface {
	GetList(ctx context.Context, paginationData common.PaginationRequest) (objects.CreditQuotaListWithPagination, *constants.ErrorResponse)
	Create(ctx context.Context, data objects.CreateCreditQuota) *constants.ErrorResponse
	Update(ctx context.Context, data objects.UpdateCreditQuota) *constants.ErrorResponse
	Delete(ctx context.Context, id string) *constants.ErrorResponse
}

func NewCreditQuotaService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) CreditQuotaServiceInterface {
	return &creditQuotaService{
		repoCtx,
		infraCtx,
	}
}
