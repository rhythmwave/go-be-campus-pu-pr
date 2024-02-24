package class_grade_component

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
)

type ClassGradeComponentServiceInterface interface {
	GetList(ctx context.Context, classId string) ([]objects.GetClassGradeComponent, *constants.ErrorResponse)
	Set(ctx context.Context, classId string, data []objects.SetClassGradeComponent) *constants.ErrorResponse
}

func NewClassGradeComponentService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) ClassGradeComponentServiceInterface {
	return &classGradeComponentService{
		repoCtx,
		infraCtx,
	}
}
