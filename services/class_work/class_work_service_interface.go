package class_work

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type ClassWorkServiceInterface interface {
	GetList(ctx context.Context, paginationData common.PaginationRequest, appType string, classIds []string) (objects.ClassWorkListWithPagination, *constants.ErrorResponse)
	GetSubmission(ctx context.Context, paginationData common.PaginationRequest, classWorkId string) (objects.ClassWorkSubmissionWithPagination, *constants.ErrorResponse)
	Create(ctx context.Context, data objects.CreateClassWork) *constants.ErrorResponse
	Update(ctx context.Context, data objects.UpdateClassWork) *constants.ErrorResponse
	Delete(ctx context.Context, ids []string) *constants.ErrorResponse
	GradeSubmission(ctx context.Context, classWorkId string, data []objects.GradeClassWorkSubmission) *constants.ErrorResponse
	Submit(ctx context.Context, classWorkId, filePath, filePathType string) *constants.ErrorResponse
}

func NewClassWorkService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) ClassWorkServiceInterface {
	return &classWorkService{
		repoCtx,
		infraCtx,
	}
}
