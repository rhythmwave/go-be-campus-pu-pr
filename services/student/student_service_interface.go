package student

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type StudentServiceInterface interface {
	GetList(ctx context.Context, paginationData common.PaginationRequest, requestData objects.GetStudentRequest) (objects.StudentListWithPagination, *constants.ErrorResponse)
	GetDetail(ctx context.Context, id string) (objects.GetStudent, *constants.ErrorResponse)
	Create(ctx context.Context, data objects.CreateStudent) *constants.ErrorResponse
	Update(ctx context.Context, data objects.UpdateStudent) *constants.ErrorResponse
	Delete(ctx context.Context, id string) *constants.ErrorResponse
	BulkUpdateStatus(ctx context.Context, data objects.BulkUpdateStatusStudent) *constants.ErrorResponse
	GetStatusSummary(ctx context.Context, semesterId string) ([]objects.GetStatusSummaryStudent, *constants.ErrorResponse)
	GetSemesterSummary(ctx context.Context) (objects.GetStudentSemesterSummary, *constants.ErrorResponse)
	GetProfile(ctx context.Context) (objects.GetStudent, *constants.ErrorResponse)
	UpdateProfile(ctx context.Context, data objects.UpdateStudentProfile) *constants.ErrorResponse
	UpdateParentProfile(ctx context.Context, data objects.UpdateStudentParentProfile) *constants.ErrorResponse
	UpdateSchoolProfile(ctx context.Context, data objects.UpdateStudentSchoolProfile) *constants.ErrorResponse
	GetSubjectGrade(ctx context.Context, paginationData common.PaginationRequest, studentId string) (objects.StudentSubjectWithPagination, *constants.ErrorResponse)
	BulkUpdatePayment(ctx context.Context, studentIds []string) *constants.ErrorResponse
	GetPaymentLog(ctx context.Context, studentId string) ([]objects.GetStudentPaymentLog, *constants.ErrorResponse)
	BulkCreate(ctx context.Context, data []objects.BulkCreateStudent) ([]objects.BulkCreateAuthenticationResponse, *constants.ErrorResponse)
}

func NewStudentService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) StudentServiceInterface {
	return &studentService{
		repoCtx,
		infraCtx,
	}
}
