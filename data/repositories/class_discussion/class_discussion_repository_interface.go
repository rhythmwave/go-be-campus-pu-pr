package class_discussion

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type ClassDiscussionRepositoryInterface interface {
	GetList(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, classId string) ([]models.GetClassDiscussion, common.Pagination, *constants.ErrorResponse)
	GetDetail(ctx context.Context, tx *sqlx.Tx, id string) (models.GetClassDiscussion, *constants.ErrorResponse)
	GetComment(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, classDiscussionId string) ([]models.GetClassDiscussionComment, common.Pagination, *constants.ErrorResponse)
	GetDetailComment(ctx context.Context, tx *sqlx.Tx, id string) (models.GetClassDiscussionComment, *constants.ErrorResponse)
	Create(ctx context.Context, tx *sqlx.Tx, data models.CreateClassDiscussion) *constants.ErrorResponse
	Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateClassDiscussion) *constants.ErrorResponse
	Delete(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse
	CreateComment(ctx context.Context, tx *sqlx.Tx, data models.CreateClassDiscussionComment) *constants.ErrorResponse
	DeleteComment(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse
}

func NewClassDiscussionRepository(db *db.DB) ClassDiscussionRepositoryInterface {
	return &classDiscussionRepository{
		db,
	}
}
