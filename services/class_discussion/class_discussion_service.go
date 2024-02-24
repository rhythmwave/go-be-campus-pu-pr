package class_discussion

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	appConstants "github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
	"github.com/sccicitb/pupr-backend/utils"
)

type classDiscussionService struct {
	*repository.RepoCtx
	*infra.InfraCtx
}

func (c classDiscussionService) GetList(ctx context.Context, paginationData common.PaginationRequest, classId string) (objects.ClassDiscussionListWithPagination, *constants.ErrorResponse) {
	var result objects.ClassDiscussionListWithPagination

	tx, err := c.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	modelResult, paginationResult, errs := c.ClassDiscussionRepo.GetList(ctx, tx, paginationData, classId)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	resultData := []objects.GetClassDiscussion{}
	for _, v := range modelResult {
		resultData = append(resultData, objects.GetClassDiscussion{
			Id:                 v.Id,
			Title:              v.Title,
			Abstraction:        v.Abstraction,
			LecturerId:         v.LecturerId,
			LecturerName:       v.LecturerName,
			LecturerFrontTitle: v.LecturerFrontTitle,
			LecturerBackDegree: v.LecturerBackDegree,
			TotalComment:       v.TotalComment,
			LastComment:        v.LastComment,
		})
	}

	result = objects.ClassDiscussionListWithPagination{
		Pagination: paginationResult,
		Data:       resultData,
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, constants.ErrUnknown
	}

	return result, nil
}

func (c classDiscussionService) GetComment(ctx context.Context, paginationData common.PaginationRequest, classDiscussionId string) (objects.ClassDiscussionCommentWithPagination, *constants.ErrorResponse) {
	var result objects.ClassDiscussionCommentWithPagination

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return result, errs
	}

	tx, err := c.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	modelResult, paginationResult, errs := c.ClassDiscussionRepo.GetComment(ctx, tx, paginationData, classDiscussionId)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	resultData := []objects.GetClassDiscussionComment{}
	for _, v := range modelResult {
		var selfComment bool
		if claims.Role == appConstants.AppTypeLecturer && claims.ID == utils.NullStringScan(v.LecturerId) {
			selfComment = true
		}
		if claims.Role == appConstants.AppTypeStudent && claims.ID == utils.NullStringScan(v.StudentId) {
			selfComment = true
		}

		resultData = append(resultData, objects.GetClassDiscussionComment{
			Id:                 v.Id,
			StudentId:          v.StudentId,
			StudentNimNumber:   v.StudentNimNumber,
			StudentName:        v.StudentName,
			LecturerId:         v.LecturerId,
			LecturerName:       v.LecturerName,
			LecturerFrontTitle: v.LecturerFrontTitle,
			LecturerBackDegree: v.LecturerBackDegree,
			Comment:            v.Comment,
			SelfComment:        selfComment,
		})
	}

	result = objects.ClassDiscussionCommentWithPagination{
		Pagination: paginationResult,
		Data:       resultData,
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, constants.ErrUnknown
	}

	return result, nil
}

func (c classDiscussionService) Create(ctx context.Context, data objects.CreateClassDiscussion) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := c.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	classLecturerData, errs := c.ClassLecturerRepo.GetByClassIdLecturerId(ctx, tx, data.ClassId, claims.ID)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}
	if classLecturerData.Id == "" {
		_ = tx.Rollback()
		return appConstants.ErrUneditableClassDiscussion
	}

	createData := models.CreateClassDiscussion{
		LecturerId:  claims.ID,
		ClassId:     data.ClassId,
		Title:       data.Title,
		Abstraction: data.Abstraction,
	}
	errs = c.ClassDiscussionRepo.Create(ctx, tx, createData)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return constants.ErrUnknown
	}

	return nil
}

func (c classDiscussionService) Update(ctx context.Context, data objects.UpdateClassDiscussion) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := c.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	classDiscussionData, errs := c.ClassDiscussionRepo.GetDetail(ctx, tx, data.Id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}
	if classDiscussionData.LecturerId != claims.ID {
		_ = tx.Rollback()
		return appConstants.ErrUneditableClassDiscussion
	}

	updateData := models.UpdateClassDiscussion{
		Id:          data.Id,
		Title:       data.Title,
		Abstraction: data.Abstraction,
	}
	errs = c.ClassDiscussionRepo.Update(ctx, tx, updateData)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return constants.ErrUnknown
	}

	return nil
}

func (c classDiscussionService) Delete(ctx context.Context, id string) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := c.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	classDiscussionData, errs := c.ClassDiscussionRepo.GetDetail(ctx, tx, id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}
	if classDiscussionData.LecturerId != claims.ID {
		_ = tx.Rollback()
		return appConstants.ErrUneditableClassDiscussion
	}

	errs = c.ClassDiscussionRepo.Delete(ctx, tx, id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return constants.ErrUnknown
	}

	return nil
}

func (c classDiscussionService) CreateComment(ctx context.Context, data objects.CreateClassDiscussionComment) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := c.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	createData := models.CreateClassDiscussionComment{
		ClassDiscussionId: data.ClassDiscussionId,
		Comment:           data.Comment,
	}
	if claims.Role == appConstants.AppTypeLecturer {
		createData.LecturerId = utils.NewNullString(claims.ID)
	} else {
		createData.StudentId = utils.NewNullString(claims.ID)
	}

	errs = c.ClassDiscussionRepo.CreateComment(ctx, tx, createData)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return constants.ErrUnknown
	}

	return nil
}

func (c classDiscussionService) DeleteComment(ctx context.Context, id string) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := c.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	commentData, errs := c.ClassDiscussionRepo.GetDetailComment(ctx, tx, id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}
	isLecturer := claims.Role == appConstants.AppTypeLecturer
	if (isLecturer && utils.NullStringScan(commentData.LecturerId) != claims.ID) || (!isLecturer && utils.NullStringScan(commentData.StudentId) != claims.ID) {
		return appConstants.ErrUneditableClassDiscussionComment
	}

	errs = c.ClassDiscussionRepo.DeleteComment(ctx, tx, id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return constants.ErrUnknown
	}

	return nil
}
