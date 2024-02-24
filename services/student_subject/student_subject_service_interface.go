package student_subject

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
)

type StudentSubjectServiceInterface interface {
	GetDetail(ctx context.Context, studentId string) (objects.GetDetailStudentSubject, *constants.ErrorResponse)
	GetPdfDetail(ctx context.Context, studentId string) (objects.GetTranscriptDetail, *constants.ErrorResponse)
	ConvertMbkmGrade(ctx context.Context, data objects.ConvertMbkmGrade) *constants.ErrorResponse
}

func NewStudentSubjectService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) StudentSubjectServiceInterface {
	return &studentSubjectService{
		repoCtx,
		infraCtx,
	}
}
