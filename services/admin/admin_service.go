package admin

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

type adminService struct {
	*repository.RepoCtx
	*infra.InfraCtx
}

func mapGetList(adminData []models.GetAdmin, studyProgramData []models.GetStudyProgramByRoleIds, permissionData []models.GetPermissionByRoleIds) []objects.GetAdmin {
	results := []objects.GetAdmin{}
	tempStudyProgram := make(map[string][]objects.GetStudyProgram)
	tempPermission := make(map[string][]objects.GetPermission)

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

	for _, v := range adminData {
		results = append(results, objects.GetAdmin{
			Id:            v.Id,
			Username:      v.Username,
			Name:          v.Name,
			RoleId:        v.RoleId,
			RoleName:      v.RoleName,
			StudyPrograms: tempStudyProgram[v.RoleId],
			Permissions:   tempPermission[v.RoleId],
		})
	}

	return results
}

func (a adminService) GetList(ctx context.Context, paginationData common.PaginationRequest) (objects.AdminListWithPagination, *constants.ErrorResponse) {
	var result objects.AdminListWithPagination

	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	modelResult, paginationResult, errs := a.AdminRepo.GetList(ctx, tx, paginationData)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	roleIds := []string{}
	for _, v := range modelResult {
		roleIds = append(roleIds, v.RoleId)
	}

	studyProgramData, errs := a.StudyProgramRepo.GetByRoleIds(ctx, tx, roleIds)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	permissionData, errs := a.PermissionRepo.GetByRoleIds(ctx, tx, roleIds)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	result = objects.AdminListWithPagination{
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

func (a adminService) Create(ctx context.Context, data objects.CreateAdmin) *constants.ErrorResponse {
	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	password, err := utils.HashPassword(data.Password)
	if err != nil {
		_ = tx.Rollback()
		return constants.ErrorInternalServer(err.Error())
	}
	createData := models.CreateAdmin{
		Username:  data.Username,
		Name:      data.Name,
		Password:  password,
		RoleId:    data.RoleId,
		CreatedBy: claims.ID,
	}
	errs = a.AdminRepo.Create(ctx, tx, createData)
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

func (a adminService) Update(ctx context.Context, data objects.UpdateAdmin) *constants.ErrorResponse {
	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs := a.AdminRepo.GetDetail(ctx, tx, data.Id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	var password string
	if data.Password != "" {
		password, err = utils.HashPassword(data.Password)
		if err != nil {
			_ = tx.Rollback()
			return constants.ErrorInternalServer(err.Error())
		}
	}
	updateData := models.UpdateAdmin{
		Id:        data.Id,
		Username:  data.Username,
		Name:      data.Name,
		Password:  utils.NewNullString(password),
		RoleId:    data.RoleId,
		UpdatedBy: claims.ID,
	}
	errs = a.AdminRepo.Update(ctx, tx, updateData)
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

func (a adminService) Delete(ctx context.Context, id string) *constants.ErrorResponse {
	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs := a.AdminRepo.GetDetail(ctx, tx, id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	errs = a.AdminRepo.Delete(ctx, tx, id)
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
