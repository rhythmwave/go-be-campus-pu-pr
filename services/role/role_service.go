package role

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
	"github.com/sccicitb/pupr-backend/utils"
)

type roleService struct {
	*repository.RepoCtx
	*infra.InfraCtx
}

func mapGetList(roleData []models.GetRoleList, studyProgramData []models.GetStudyProgramByRoleIds, permissionData []models.GetPermissionByRoleIds) []objects.GetRole {
	results := []objects.GetRole{}
	tempRole := make(map[string]objects.GetRole)
	tempStudyProgram := make(map[string][]objects.GetStudyProgram)
	tempPermission := make(map[string][]objects.GetPermission)

	for _, v := range roleData {
		tempRole[v.Id] = objects.GetRole{
			Id:          v.Id,
			Name:        v.Name,
			Description: v.Description,
		}
	}
	for _, v := range studyProgramData {
		tempStudyProgram[v.RoleId] = append(tempStudyProgram[v.RoleId], objects.GetStudyProgram{
			Id:                    v.Id,
			Name:                  v.Name,
			StudyLevelShortName:   v.StudyLevelShortName,
			DiktiStudyProgramType: v.DiktiStudyProgramType,
		})
	}
	for _, v := range permissionData {
		tempPermission[v.RoleId] = append(tempPermission[v.RoleId], objects.GetPermission{
			Id:   v.Id,
			Name: v.Name,
		})
	}

	for k, v := range tempStudyProgram {
		if entry, ok := tempRole[k]; ok {
			entry.StudyPrograms = v
			tempRole[k] = entry
		}
	}
	for k, v := range tempPermission {
		if entry, ok := tempRole[k]; ok {
			entry.Permissions = v
			tempRole[k] = entry
		}
	}

	for _, v := range tempRole {
		results = append(results, v)
	}

	return results
}

func (r roleService) GetList(ctx context.Context, paginationData common.PaginationRequest) (objects.RoleListWithPagination, *constants.ErrorResponse) {
	var result objects.RoleListWithPagination

	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	modelResult, paginationResult, errs := r.RoleRepo.GetList(ctx, tx, paginationData)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}
	roleIds := []string{}
	for _, v := range modelResult {
		roleIds = append(roleIds, v.Id)
	}

	studyProgramData, errs := r.StudyProgramRepo.GetByRoleIds(ctx, tx, roleIds)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	permissionData, errs := r.PermissionRepo.GetByRoleIds(ctx, tx, roleIds)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	result = objects.RoleListWithPagination{
		Pagination: paginationResult,
		Data:       mapGetList(modelResult, studyProgramData, permissionData),
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, constants.ErrUnknown
	}

	return result, nil
}

func (r roleService) Create(ctx context.Context, data objects.CreateRole) *constants.ErrorResponse {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	createData := models.CreateRole{
		Name:          data.Name,
		Description:   data.Description,
		StudyPrograms: data.StudyProgramIds,
		Permissions:   data.PermissionIds,
		CreatedBy:     claims.ID,
	}
	errs = r.RoleRepo.Create(ctx, tx, createData)
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

func (r roleService) Update(ctx context.Context, data objects.UpdateRole) *constants.ErrorResponse {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs := r.RoleRepo.GetDetail(ctx, tx, data.Id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	updateData := models.UpdateRole{
		Id:            data.Id,
		Name:          data.Name,
		Description:   data.Description,
		StudyPrograms: data.StudyProgramIds,
		Permissions:   data.PermissionIds,
		UpdatedBy:     claims.ID,
	}
	errs = r.RoleRepo.Update(ctx, tx, updateData)
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

func (r roleService) Delete(ctx context.Context, id string) *constants.ErrorResponse {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs := r.RoleRepo.GetDetail(ctx, tx, id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	errs = r.RoleRepo.Delete(ctx, tx, id)
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
