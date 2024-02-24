package admin_activity_log

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/infra/middleware"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
	"github.com/sccicitb/pupr-backend/utils"
)

type adminActivityLogService struct {
	*repository.RepoCtx
	*infra.InfraCtx
	middleware.MiddlewareInterface
}

func (a adminActivityLogService) GetList(ctx context.Context, paginationData common.PaginationRequest, year, month uint32) (objects.AdminActivityLogWithPagination, *constants.ErrorResponse) {
	var result objects.AdminActivityLogWithPagination

	tx, err := a.DBLog.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	modelResult, paginationResult, errs := a.AdminActivityLogRepo.GetList(ctx, tx, paginationData, year, month)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	resultData := []objects.GetAdminActivityLog{}
	for _, v := range modelResult {
		resultData = append(resultData, objects.GetAdminActivityLog{
			Id:            v.Id,
			AdminId:       v.AdminId,
			AdminName:     v.AdminName,
			AdminUsername: v.AdminUsername,
			Module:        v.Module,
			Action:        v.Action,
			IpAddress:     v.IpAddress,
			UserAgent:     v.UserAgent,
			ExecutionTime: v.ExecutionTime,
			MemoryUsage:   v.MemoryUsage,
			CreatedAt:     v.CreatedAt,
		})
	}

	result = objects.AdminActivityLogWithPagination{
		Pagination: paginationResult,
		Data:       resultData,
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, constants.ErrUnknown
	}

	return result, nil
}

func (a adminActivityLogService) Create(ctx context.Context, r *http.Request, startTime time.Time, module, action string, body interface{}) *constants.ErrorResponse {
	tx, err := a.DBLog.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	if body != nil {
		action = fmt.Sprintf("%s\n%s", action, utils.StructToJson(body))
	}

	createData := models.CreateAdminActivityLog{
		AdminId:       claims.ID,
		AdminName:     claims.Name,
		AdminUsername: claims.Email,
		Module:        module,
		Action:        action,
		IpAddress:     utils.GetIP(r),
		UserAgent:     r.UserAgent(),
		ExecutionTime: time.Since(startTime).Seconds(),
		MemoryUsage:   utils.GetMemoryUsage().Alloc,
	}

	errs = a.AdminActivityLogRepo.Create(ctx, tx, createData)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return constants.ErrUnknown
	}

	return nil
}
