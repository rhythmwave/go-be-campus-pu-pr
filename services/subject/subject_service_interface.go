package subject

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type SubjectServiceInterface interface {
	GetList(ctx context.Context, paginationData common.PaginationRequest, req objects.GetSubjectRequest) (objects.SubjectListWithPagination, *constants.ErrorResponse)
	GetDetail(ctx context.Context, id string) (objects.GetSubjectDetail, *constants.ErrorResponse)
	Create(ctx context.Context, data objects.CreateSubject) *constants.ErrorResponse
	Update(ctx context.Context, data objects.UpdateSubject) *constants.ErrorResponse
	Delete(ctx context.Context, id string) *constants.ErrorResponse
	SetPrerequisiteSubject(ctx context.Context, subjectId string, prerequisites []objects.SetPrerequisiteSubject) *constants.ErrorResponse
	SetEquivalentSubject(ctx context.Context, data objects.SetEquivalentSubject) *constants.ErrorResponse
	DeleteEquivalentSubject(ctx context.Context, subjectId, equivalentSubjectId string) *constants.ErrorResponse
}

func NewSubjectService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) SubjectServiceInterface {
	return &subjectService{
		repoCtx,
		infraCtx,
	}
}
