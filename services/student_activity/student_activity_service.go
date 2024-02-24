package student_activity

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
	appUtils "github.com/sccicitb/pupr-backend/utils"
)

type studentActivityService struct {
	*repository.RepoCtx
	*infra.InfraCtx
}

func (s studentActivityService) GetList(ctx context.Context, pagination common.PaginationRequest, activityType, studyProgramId, semesterId string, isMbkm bool) (objects.StudentActivityListWithPagination, *constants.ErrorResponse) {
	var result objects.StudentActivityListWithPagination

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return result, errs
	}

	tx, err := s.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	if studyProgramId != "" {
		_, errs = s.StudyProgramRepo.GetDetail(ctx, tx, studyProgramId, claims.Role, claims.ID)
		if errs != nil {
			_ = tx.Rollback()
			return result, errs
		}
	}

	modelResult, paginationResult, errs := s.StudentActivityRepo.GetList(ctx, tx, pagination, activityType, studyProgramId, semesterId, isMbkm)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	resultData := []objects.GetStudentActivity{}
	for _, v := range modelResult {
		resultData = append(resultData, objects.GetStudentActivity{
			Id:                 v.Id,
			StudyProgramId:     v.StudyProgramId,
			StudyProgramName:   v.StudyProgramName,
			SemesterId:         v.SemesterId,
			SemesterSchoolYear: appUtils.GenerateSchoolYear(v.SemesterStartYear),
			SemesterType:       v.SemesterType,
			ActivityType:       v.ActivityType,
			Title:              v.Title,
		})
	}

	result = objects.StudentActivityListWithPagination{
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

func (s studentActivityService) GetDetail(ctx context.Context, id string) (objects.GetStudentActivityDetail, *constants.ErrorResponse) {
	var result objects.GetStudentActivityDetail

	tx, err := s.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	resultData, errs := s.StudentActivityRepo.GetDetail(ctx, tx, id)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	participantData, errs := s.StudentActivityRepo.GetListParticipantByStudentActivityId(ctx, tx, id)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	participants := []objects.GetStudentActivityDetailParticipant{}
	for _, v := range participantData {
		participants = append(participants, objects.GetStudentActivityDetailParticipant{
			StudentId:        v.StudentId,
			NimNumber:        v.NimNumber,
			Name:             v.Name,
			StudyProgramId:   v.StudyProgramId,
			StudyProgramName: v.StudyProgramName,
			Role:             v.Role,
		})
	}

	lecturerData, errs := s.StudentActivityRepo.GetListLecturerByStudentActivityId(ctx, tx, id)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	mentors := []objects.GetStudentActivityDetailLecturer{}
	examiners := []objects.GetStudentActivityDetailLecturer{}
	for _, v := range lecturerData {
		switch v.Role {
		case appConstants.StudentActivityLecturerMentor:
			mentors = append(mentors, objects.GetStudentActivityDetailLecturer{
				LecturerId:         v.LecturerId,
				IdNationalLecturer: v.IdNationalLecturer,
				Name:               v.Name,
				FrontTitle:         v.FrontTitle,
				BackDegree:         v.BackDegree,
				ActivityCategory:   v.ActivityCategory,
				Sort:               v.Sort,
			})
		case appConstants.StudentActivityLecturerExaminer:
			examiners = append(examiners, objects.GetStudentActivityDetailLecturer{
				LecturerId:         v.LecturerId,
				IdNationalLecturer: v.IdNationalLecturer,
				Name:               v.Name,
				FrontTitle:         v.FrontTitle,
				BackDegree:         v.BackDegree,
				ActivityCategory:   v.ActivityCategory,
				Sort:               v.Sort,
			})
		}
	}

	result = objects.GetStudentActivityDetail{
		Id:                 resultData.Id,
		StudyProgramId:     resultData.StudyProgramId,
		StudyProgramName:   resultData.StudyProgramName,
		SemesterId:         resultData.SemesterId,
		SemesterSchoolYear: appUtils.GenerateSchoolYear(resultData.SemesterStartYear),
		SemesterType:       resultData.SemesterType,
		ActivityType:       resultData.ActivityType,
		Title:              resultData.Title,
		Location:           resultData.Location,
		DecisionNumber:     resultData.DecisionNumber,
		DecisionDate:       resultData.DecisionDate,
		IsGroupActivity:    resultData.IsGroupActivity,
		Remarks:            resultData.Remarks,
		Participants:       participants,
		Mentors:            mentors,
		Examiners:          examiners,
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, constants.ErrUnknown
	}

	return result, nil
}

func (s studentActivityService) Create(ctx context.Context, data objects.CreateStudentActivity) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := s.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	createData := models.CreateStudentActivity{
		StudyProgramId:  data.StudyProgramId,
		SemesterId:      data.SemesterId,
		ActivityType:    data.ActivityType,
		Title:           data.Title,
		Location:        utils.NewNullString(data.Location),
		DecisionNumber:  utils.NewNullString(data.DecisionNumber),
		DecisionDate:    utils.NewNullString(data.DecisionDate),
		IsGroupActivity: data.IsGroupActivity,
		Remarks:         utils.NewNullString(data.Remarks),
		IsMbkm:          data.IsMbkm,
		CreatedBy:       claims.ID,
	}
	studentActivityId, errs := s.StudentActivityRepo.Create(ctx, tx, createData)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	if len(data.Participants) != 0 {
		participantData := []models.UpsertStudentActivityParticipant{}
		for _, v := range data.Participants {
			participantData = append(participantData, models.UpsertStudentActivityParticipant{
				StudentActivityId: studentActivityId,
				StudentId:         v.StudentId,
				Role:              v.Role,
				CreatedBy:         claims.ID,
			})
		}
		errs = s.StudentActivityRepo.UpsertParticipant(ctx, tx, participantData)
		if errs != nil {
			_ = tx.Rollback()
			return errs
		}
	}

	lecturerData := []models.UpsertStudentActivityLecturer{}
	for _, v := range data.Mentors {
		lecturerData = append(lecturerData, models.UpsertStudentActivityLecturer{
			StudentActivityId: studentActivityId,
			LecturerId:        v.LecturerId,
			ActivityCategory:  v.ActivityCategory,
			Role:              appConstants.StudentActivityLecturerMentor,
			Sort:              v.Sort,
			CreatedBy:         claims.ID,
		})
	}
	for _, v := range data.Examiners {
		lecturerData = append(lecturerData, models.UpsertStudentActivityLecturer{
			StudentActivityId: studentActivityId,
			LecturerId:        v.LecturerId,
			ActivityCategory:  v.ActivityCategory,
			Role:              appConstants.StudentActivityLecturerExaminer,
			Sort:              v.Sort,
			CreatedBy:         claims.ID,
		})
	}

	if len(lecturerData) != 0 {
		errs = s.StudentActivityRepo.UpsertLecturer(ctx, tx, lecturerData)
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

func (s studentActivityService) Update(ctx context.Context, data objects.UpdateStudentActivity) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := s.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs = s.StudentActivityRepo.GetDetail(ctx, tx, data.Id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	updateData := models.UpdateStudentActivity{
		Id:              data.Id,
		StudyProgramId:  data.StudyProgramId,
		SemesterId:      data.SemesterId,
		ActivityType:    data.ActivityType,
		Title:           data.Title,
		Location:        utils.NewNullString(data.Location),
		DecisionNumber:  utils.NewNullString(data.DecisionNumber),
		DecisionDate:    utils.NewNullString(data.DecisionDate),
		IsGroupActivity: data.IsGroupActivity,
		Remarks:         utils.NewNullString(data.Remarks),
		IsMbkm:          data.IsMbkm,
		UpdatedBy:       claims.ID,
	}
	errs = s.StudentActivityRepo.Update(ctx, tx, updateData)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	studentIds := []string{}
	participantData := []models.UpsertStudentActivityParticipant{}
	for _, v := range data.Participants {
		studentIds = append(studentIds, v.StudentId)
		participantData = append(participantData, models.UpsertStudentActivityParticipant{
			StudentActivityId: data.Id,
			StudentId:         v.StudentId,
			Role:              v.Role,
			CreatedBy:         claims.ID,
		})
	}
	errs = s.StudentActivityRepo.DeleteParticipantExcludingStudentIds(ctx, tx, data.Id, studentIds)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}
	if len(participantData) != 0 {
		errs = s.StudentActivityRepo.UpsertParticipant(ctx, tx, participantData)
		if errs != nil {
			_ = tx.Rollback()
			return errs
		}
	}

	mentorIds := []string{}
	examinerIds := []string{}
	lecturerData := []models.UpsertStudentActivityLecturer{}
	for _, v := range data.Mentors {
		mentorIds = append(mentorIds, v.LecturerId)
		lecturerData = append(lecturerData, models.UpsertStudentActivityLecturer{
			StudentActivityId: data.Id,
			LecturerId:        v.LecturerId,
			ActivityCategory:  v.ActivityCategory,
			Role:              appConstants.StudentActivityLecturerMentor,
			Sort:              v.Sort,
			CreatedBy:         claims.ID,
		})
	}
	for _, v := range data.Examiners {
		examinerIds = append(examinerIds, v.LecturerId)
		lecturerData = append(lecturerData, models.UpsertStudentActivityLecturer{
			StudentActivityId: data.Id,
			LecturerId:        v.LecturerId,
			ActivityCategory:  v.ActivityCategory,
			Role:              appConstants.StudentActivityLecturerExaminer,
			Sort:              v.Sort,
			CreatedBy:         claims.ID,
		})
	}

	errs = s.StudentActivityRepo.DeleteLecturerExcludingLecturerIds(ctx, tx, data.Id, appConstants.StudentActivityLecturerMentor, mentorIds)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}
	errs = s.StudentActivityRepo.DeleteLecturerExcludingLecturerIds(ctx, tx, data.Id, appConstants.StudentActivityLecturerExaminer, examinerIds)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	if len(lecturerData) != 0 {
		errs = s.StudentActivityRepo.UpsertLecturer(ctx, tx, lecturerData)
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

func (s studentActivityService) Delete(ctx context.Context, id string) *constants.ErrorResponse {
	tx, err := s.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs := s.StudentActivityRepo.GetDetail(ctx, tx, id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	errs = s.StudentActivityRepo.Delete(ctx, tx, id)
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
