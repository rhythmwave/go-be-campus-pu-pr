package subject_grade_component

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
)

type SubjectGradeComponentServiceInterface interface {
	GetList(ctx context.Context, subjectId string) ([]objects.GetSubjectGradeComponent, *constants.ErrorResponse)
	Set(ctx context.Context, subjectId string, data []objects.SetSubjectGradeComponent) *constants.ErrorResponse
}

func NewSubjectGradeComponentService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) SubjectGradeComponentServiceInterface {
	return &subjectGradeComponentService{
		repoCtx,
		infraCtx,
	}
}
