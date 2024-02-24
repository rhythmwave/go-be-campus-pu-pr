package permission

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type permissionService struct {
	*repository.RepoCtx
	*infra.InfraCtx
}

func (p permissionService) GetList(ctx context.Context, paginationData common.PaginationRequest) (objects.PermissionListWithPagination, *constants.ErrorResponse) {
	var result objects.PermissionListWithPagination

	tx, err := p.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	modelResult, paginationResult, errs := p.PermissionRepo.GetList(ctx, tx, paginationData)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	resultData := []objects.GetPermission{}
	for _, v := range modelResult {
		resultData = append(resultData, objects.GetPermission{
			Id:   v.Id,
			Name: v.Name,
		})
	}

	result = objects.PermissionListWithPagination{
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
