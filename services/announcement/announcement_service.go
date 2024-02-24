package announcement

import (
	"context"

	"github.com/google/uuid"
	"github.com/sccicitb/pupr-backend/constants"
	appConstants "github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
	"github.com/sccicitb/pupr-backend/utils"
)

type announcementService struct {
	*repository.RepoCtx
	*infra.InfraCtx
}

func (a announcementService) mapGetList(data []models.GetAnnouncement, announcementStudyProgramData []models.GetAnnouncementStudyProgram) ([]objects.GetAnnouncement, *constants.ErrorResponse) {
	results := []objects.GetAnnouncement{}
	var errs *constants.ErrorResponse

	announcementStudyProgramMap := make(map[string][]objects.GetAnnouncementStudyProgram)
	for _, v := range announcementStudyProgramData {
		announcementStudyProgramMap[v.AnnouncementId] = append(announcementStudyProgramMap[v.AnnouncementId], objects.GetAnnouncementStudyProgram{
			StudyProgramId:   v.StudyProgramId,
			StudyProgramName: v.StudyProgramName,
		})
	}

	for _, v := range data {
		var fileUrl string
		if v.FilePath != nil && v.FilePathType != nil {
			fileUrl, errs = a.Storage.GetURL(*v.FilePath, *v.FilePathType, nil)
			if errs != nil {
				return results, errs
			}
		}

		results = append(results, objects.GetAnnouncement{
			Id:               v.Id,
			Type:             v.Type,
			Title:            v.Title,
			AnnouncementDate: v.AnnouncementDate,
			FileUrl:          fileUrl,
			FilePath:         v.FilePath,
			FilePathType:     v.FilePathType,
			FileTitle:        v.FileTitle,
			Content:          v.Content,
			ForLecturer:      v.ForLecturer,
			ForStudent:       v.ForStudent,
			StudyPrograms:    announcementStudyProgramMap[v.Id],
		})
	}

	return results, nil
}

func (a announcementService) GetList(ctx context.Context, paginationData common.PaginationRequest, appType string, announcementType string) (objects.AnnouncementListWithPagination, *constants.ErrorResponse) {
	var result objects.AnnouncementListWithPagination

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return result, errs
	}

	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	var studyProgramIds []string
	switch appType {
	case appConstants.AppTypeAdmin:
		studyProgramData, _, errs := a.StudyProgramRepo.GetList(ctx, tx, common.PaginationRequest{Page: constants.DefaultPage, Limit: constants.DefaultUnlimited}, "", claims.Role, claims.ID)
		if errs != nil {
			_ = tx.Rollback()
			return result, errs
		}
		for _, v := range studyProgramData {
			studyProgramIds = append(studyProgramIds, v.Id)
		}
	case appConstants.AppTypeLecturer:
		lecturerData, errs := a.LecturerRepo.GetDetail(ctx, tx, claims.ID)
		if errs != nil {
			_ = tx.Rollback()
			return result, errs
		}
		if lecturerData.StudyProgramId != nil {
			studyProgramIds = append(studyProgramIds, *lecturerData.StudyProgramId)
		}
	case appConstants.AppTypeStudent:
		studentData, errs := a.StudentRepo.GetDetail(ctx, tx, claims.ID, "")
		if errs != nil {
			_ = tx.Rollback()
			return result, errs
		}
		if studentData.StudyProgramId != nil {
			studyProgramIds = append(studyProgramIds, *studentData.StudyProgramId)
		}
	}
	var refIds []string
	if len(studyProgramIds) != 0 {
		refData, errs := a.AnnouncementRepo.GetAnnouncementStudyProgramByStudyProgramIds(ctx, tx, studyProgramIds)
		if errs != nil {
			_ = tx.Rollback()
			return result, errs
		}
		for _, v := range refData {
			refIds = append(refIds, v.AnnouncementId)
		}
	}

	modelResult, paginationResult, errs := a.AnnouncementRepo.GetList(ctx, tx, paginationData, appType, announcementType, refIds)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	var announcementIds []string
	for _, v := range modelResult {
		announcementIds = append(announcementIds, v.Id)
	}

	var announcementStudyProgramData []models.GetAnnouncementStudyProgram
	if len(announcementIds) != 0 {
		announcementStudyProgramData, errs = a.AnnouncementRepo.GetAnnouncementStudyProgramByAnnouncementIds(ctx, tx, announcementIds)
		if errs != nil {
			_ = tx.Rollback()
			return result, errs
		}
	}

	resultData, errs := a.mapGetList(modelResult, announcementStudyProgramData)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	result = objects.AnnouncementListWithPagination{
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

func (a announcementService) Create(ctx context.Context, data objects.CreateAnnouncement) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	if !utils.InArrayExist(data.Type, appConstants.ValidAnnouncementType()) {
		return appConstants.ErrInvalidAnnouncementType
	}

	announcementId := uuid.New().String()
	createData := models.CreateAnnouncement{
		Id:               announcementId,
		Type:             data.Type,
		Title:            data.Title,
		AnnouncementDate: utils.NewNullTime(data.AnnouncementDate),
		FilePath:         utils.NewNullString(data.FilePath),
		FilePathType:     utils.NewNullString(data.FilePathType),
		FileTitle:        utils.NewNullString(data.FileTitle),
		Content:          utils.NewNullString(data.Content),
		ForLecturer:      data.ForLecturer,
		ForStudent:       data.ForStudent,
		CreatedBy:        claims.ID,
	}
	errs = a.AnnouncementRepo.Create(ctx, tx, createData)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	announcementStudyProgramData := []models.UpsertAnnouncementStudyProgram{}
	for _, v := range data.StudyProgramIds {
		announcementStudyProgramData = append(announcementStudyProgramData, models.UpsertAnnouncementStudyProgram{
			AnnouncementId: announcementId,
			StudyProgramId: v,
		})
	}

	if len(announcementStudyProgramData) != 0 {
		errs = a.AnnouncementRepo.UpsertAnnouncementStudyProgram(ctx, tx, announcementStudyProgramData)
		if errs != nil {
			_ = tx.Rollback()
			return errs
		}
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return constants.ErrUnknown
	}

	return nil
}

func (a announcementService) Update(ctx context.Context, data objects.UpdateAnnouncement) *constants.ErrorResponse {
	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs := a.AnnouncementRepo.GetDetail(ctx, tx, data.Id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	if !utils.InArrayExist(data.Type, appConstants.ValidAnnouncementType()) {
		return appConstants.ErrInvalidAnnouncementType
	}

	updateData := models.UpdateAnnouncement{
		Id:               data.Id,
		Type:             data.Type,
		Title:            data.Title,
		AnnouncementDate: utils.NewNullTime(data.AnnouncementDate),
		FilePath:         utils.NewNullString(data.FilePath),
		FilePathType:     utils.NewNullString(data.FilePathType),
		FileTitle:        utils.NewNullString(data.FileTitle),
		Content:          utils.NewNullString(data.Content),
		ForLecturer:      data.ForLecturer,
		ForStudent:       data.ForStudent,
		UpdatedBy:        claims.ID,
	}
	errs = a.AnnouncementRepo.Update(ctx, tx, updateData)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	errs = a.AnnouncementRepo.DeleteAnnouncementStudyProgramExcludingStudyProgramIds(ctx, tx, data.Id, data.StudyProgramIds)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	announcementStudyProgramData := []models.UpsertAnnouncementStudyProgram{}
	for _, v := range data.StudyProgramIds {
		announcementStudyProgramData = append(announcementStudyProgramData, models.UpsertAnnouncementStudyProgram{
			AnnouncementId: data.Id,
			StudyProgramId: v,
		})
	}

	if len(announcementStudyProgramData) != 0 {
		errs = a.AnnouncementRepo.UpsertAnnouncementStudyProgram(ctx, tx, announcementStudyProgramData)
		if errs != nil {
			_ = tx.Rollback()
			return errs
		}
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return constants.ErrUnknown
	}

	return nil
}

func (a announcementService) Delete(ctx context.Context, id string) *constants.ErrorResponse {
	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs := a.AnnouncementRepo.GetDetail(ctx, tx, id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	errs = a.AnnouncementRepo.Delete(ctx, tx, id)
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
