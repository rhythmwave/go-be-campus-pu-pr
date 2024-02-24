package class_exam

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type ClassExamServiceInterface interface {
	GetList(ctx context.Context, paginationData common.PaginationRequest, appType string, classIds []string) (objects.ClassExamListWithPagination, *constants.ErrorResponse)
	GetSubmission(ctx context.Context, paginationData common.PaginationRequest, classExamId string) (objects.ClassExamSubmissionWithPagination, *constants.ErrorResponse)
	Create(ctx context.Context, data objects.CreateClassExam) *constants.ErrorResponse
	Update(ctx context.Context, data objects.UpdateClassExam) *constants.ErrorResponse
	Delete(ctx context.Context, ids []string) *constants.ErrorResponse
	GradeSubmission(ctx context.Context, classExamId string, data []objects.GradeClassExamSubmission) *constants.ErrorResponse
	Submit(ctx context.Context, classExamId, filePath, filePathType string) *constants.ErrorResponse
}

func NewClassExamService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) ClassExamServiceInterface {
	return &classExamService{
		repoCtx,
		infraCtx,
	}
}
