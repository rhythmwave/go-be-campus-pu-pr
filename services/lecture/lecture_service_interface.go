package lecture

import (
	"context"
	"time"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type LectureServiceInterface interface {
	GetList(ctx context.Context, paginationData common.PaginationRequest, classId, semesterId string, hasActualLecture, isExam *bool, examType string) (objects.LectureListWithPagination, *constants.ErrorResponse)
	GetDetail(ctx context.Context, id string) (objects.GetLectureDetail, *constants.ErrorResponse)
	BulkCreate(ctx context.Context, data objects.CreateLecture) *constants.ErrorResponse
	Update(ctx context.Context, data objects.UpdateLecture) *constants.ErrorResponse
	Delete(ctx context.Context, id string) *constants.ErrorResponse
	ResetParticipation(ctx context.Context, id string) *constants.ErrorResponse
	GetStudentParticipation(ctx context.Context, paginationData common.PaginationRequest, classId, studentId string) (objects.LectureParticipationWithPagination, *constants.ErrorResponse)
	AttendAutonomousLecture(ctx context.Context, lectureId string) *constants.ErrorResponse
	GetDetailByClassId(ctx context.Context, classId, appType string) (objects.GetDetailLecture, *constants.ErrorResponse)
	GetLectureCalendar(ctx context.Context, req objects.GetLectureCalendarRequest) ([]objects.GetLectureCalendar, *constants.ErrorResponse)
	GetHistory(ctx context.Context, paginationData common.PaginationRequest, studentId string, startDate, endDate time.Time) (objects.GetLectureHistoryWithPagination, *constants.ErrorResponse)
}

func NewLectureService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) LectureServiceInterface {
	return &lectureService{
		repoCtx,
		infraCtx,
	}
}
