package lecturer_student_activity_log

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

type LecturerStudentActivityLogServiceInterface interface {
	GetList(ctx context.Context, paginationData common.PaginationRequest, year, month uint32) (objects.LecturerStudentActivityLogWithPagination, *constants.ErrorResponse)
	Create(ctx context.Context, r *http.Request, startTime time.Time, module, action string, body interface{}) *constants.ErrorResponse
}

func NewLecturerStudentActivityLogService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx, mw middleware.MiddlewareInterface) LecturerStudentActivityLogServiceInterface {
	return &lecturerStudentActivityLogService{
		repoCtx,
		infraCtx,
		mw,
	}
}
