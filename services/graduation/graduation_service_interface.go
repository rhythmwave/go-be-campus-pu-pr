package graduation

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type GraduationServiceInterface interface {
	Apply(ctx context.Context, data objects.ApplyGraduation) *constants.ErrorResponse
	GetListStudent(ctx context.Context, pagination common.PaginationRequest, studyProgramId, graduationSessionId string) (objects.GetListStudentGraduationWithPagination, *constants.ErrorResponse)
}

func NewGraduationService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) GraduationServiceInterface {
	return &graduationService{
		repoCtx,
		infraCtx,
	}
}
