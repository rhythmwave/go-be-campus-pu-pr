package lecturer_student_activity_log

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/sccicitb/pupr-backend/constants"
	appConstants "github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/infra/middleware"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
	"github.com/sccicitb/pupr-backend/utils"
)

type lecturerStudentActivityLogService struct {
	*repository.RepoCtx
	*infra.InfraCtx
	middleware.MiddlewareInterface
}

func (a lecturerStudentActivityLogService) GetList(ctx context.Context, paginationData common.PaginationRequest, year, month uint32) (objects.LecturerStudentActivityLogWithPagination, *constants.ErrorResponse) {
	var result objects.LecturerStudentActivityLogWithPagination

	tx, err := a.DBLog.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	modelResult, paginationResult, errs := a.LecturerStudentActivityLogRepo.GetList(ctx, tx, paginationData, year, month)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	resultData := []objects.GetLecturerStudentActivityLog{}
	for _, v := range modelResult {
		var userType string
		var userId string
		var userName string
		var userUsername string

		if v.LecturerId != nil {
			userType = appConstants.AppTypeLecturer
			userId = utils.NullStringScan(v.LecturerId)
			userName = utils.NullStringScan(v.LecturerName)
			userUsername = utils.NullStringScan(v.LecturerUsername)
		}
		if v.StudentId != nil {
			userType = appConstants.AppTypeStudent
			userId = utils.NullStringScan(v.StudentId)
			userName = utils.NullStringScan(v.StudentName)
			userUsername = utils.NullStringScan(v.StudentUsername)
		}

		resultData = append(resultData, objects.GetLecturerStudentActivityLog{
			Id:            v.Id,
			UserType:      userType,
			UserId:        userId,
			UserName:      userName,
			UserUsername:  userUsername,
			Module:        v.Module,
			Action:        v.Action,
			IpAddress:     v.IpAddress,
			UserAgent:     v.UserAgent,
			ExecutionTime: v.ExecutionTime,
			MemoryUsage:   v.MemoryUsage,
			CreatedAt:     v.CreatedAt,
		})
	}

	result = objects.LecturerStudentActivityLogWithPagination{
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

func (a lecturerStudentActivityLogService) Create(ctx context.Context, r *http.Request, startTime time.Time, module, action string, body interface{}) *constants.ErrorResponse {
	tx, err := a.DBLog.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	var lecturerId string
	var lecturerName string
	var lecturerUsername string
	var studentId string
	var studentName string
	var studentUsername string
	if claims.Role == appConstants.AppTypeLecturer {
		lecturerId = claims.ID
		lecturerName = claims.Name
		lecturerUsername = claims.Email
	} else if claims.Role == appConstants.AppTypeStudent {
		studentId = claims.ID
		studentName = claims.Name
		studentUsername = claims.Email
	}

	if body != nil {
		action = fmt.Sprintf("%s\n%s", action, utils.StructToJson(body))
	}

	createData := models.CreateLecturerStudentActivityLog{
		LecturerId:       utils.NewNullString(lecturerId),
		LecturerName:     utils.NewNullString(lecturerName),
		LecturerUsername: utils.NewNullString(lecturerUsername),
		StudentId:        utils.NewNullString(studentId),
		StudentName:      utils.NewNullString(studentName),
		StudentUsername:  utils.NewNullString(studentUsername),
		Module:           module,
		Action:           action,
		IpAddress:        utils.GetIP(r),
		UserAgent:        r.UserAgent(),
		ExecutionTime:    time.Since(startTime).Seconds(),
		MemoryUsage:      utils.GetMemoryUsage().Alloc,
	}

	errs = a.LecturerStudentActivityLogRepo.Create(ctx, tx, createData)
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
