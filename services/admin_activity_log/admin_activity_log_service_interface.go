package admin_activity_log

import (
	"context"
	"net/http"
	"time"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/infra/middleware"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type AdminActivityLogServiceInterface interface {
	GetList(ctx context.Context, paginationData common.PaginationRequest, year, month uint32) (objects.AdminActivityLogWithPagination, *constants.ErrorResponse)
	Create(ctx context.Context, r *http.Request, startTime time.Time, module, action string, body interface{}) *constants.ErrorResponse
}

func NewAdminActivityLogService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx, mw middleware.MiddlewareInterface) AdminActivityLogServiceInterface {
	return &adminActivityLogService{
		repoCtx,
		infraCtx,
		mw,
	}
}
