package lecture

import (
	"context"
	"database/sql"
	"fmt"
	"sort"
	"time"

	"github.com/google/uuid"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
	"github.com/sccicitb/pupr-backend/utils"
)

type lectureService struct {
	*repository.RepoCtx
	*infra.InfraCtx
}

func mapGetList(data []models.GetLectureList, examSupervisorData []models.GetExamLectureSupervisor) []objects.GetLecture {
	var result []objects.GetLecture

	examSupervisorMap := make(map[string][]objects.GetLectureExamSupervisor)
	for _, v := range examSupervisorData {
		examSupervisorMap[v.LectureId] = append(examSupervisorMap[v.LectureId], objects.GetLectureExamSupervisor{
			Id:                     v.Id,
			Name:                   v.Name,
			FrontTitle:             v.FrontTitle,
			BackDegree:             v.BackDegree,
			ExamSupervisorRoleId:   v.ExamSupervisorRoleId,
			ExamSupervisorRoleName: v.ExamSupervisorRoleName,
		})
	}

	for _, v := range data {
		result = append(result, objects.GetLecture{
			Id:                               v.Id,
			LecturePlanDate:                  v.LecturePlanDate,
			LecturePlanDayOfWeek:             v.LecturePlanDayOfWeek,
			LecturePlanStartTime:             v.LecturePlanStartTime,
			LecturePlanEndTime:               v.LecturePlanEndTime,
			LectureActualDate:                v.LectureActualDate,
			LectureActualDayOfWeek:           v.LectureActualDayOfWeek,
			LectureActualStartTime:           v.LectureActualStartTime,
			LectureActualEndTime:             v.LectureActualEndTime,
			LecturerId:                       v.LecturerId,
			LecturerName:                     v.LecturerName,
			ForeignLecturerName:              v.ForeignLecturerName,
			ForeignLecturerSourceInstance:    v.ForeignLecturerSourceInstance,
			IsOriginalLecturer:               v.IsOriginalLecturer,
			IsManualParticipation:            v.IsManualParticipation,
			AutonomousParticipationStartTime: v.AutonomousParticipationStartTime,
			AutonomousParticipationEndTime:   v.AutonomousParticipationEndTime,
			AttendingParticipant:             v.AttendingParticipant,
			UpdatedAt:                        v.UpdatedAt,
			ClassId:                          v.ClassId,
			ClassName:                        v.ClassName,
			RoomId:                           v.RoomId,
			RoomName:                         v.RoomName,
			BuildingId:                       v.BuildingId,
			BuildingName:                     v.BuildingName,
			IsMidtermExam:                    v.IsMidtermExam,
			IsEndtermExam:                    v.IsEndtermExam,
			IsTheoryExam:                     v.IsTheoryExam,
			IsPracticumExam:                  v.IsPracticumExam,
			IsFieldPracticumExam:             v.IsFieldPracticumExam,
			SubjectCode:                      v.SubjectCode,
			SubjectName:                      v.SubjectName,
			TotalParticipant:                 v.TotalParticipant,
			ExamSupervisors:                  examSupervisorMap[v.Id],
		})
	}

	return result
}

func mapLectureCalendar(data []models.GetLectureCalendar, lecturerId string) []objects.GetLectureCalendar {
	var result []objects.GetLectureCalendar

	dateMap := make(map[time.Time][]objects.GetLectureCalendarLecture)
	for _, v := range data {
		baseCondition := v.LectureId != nil && v.LecturePlanStartTime != nil && v.LecturePlanEndTime != nil && v.ClassId != nil && v.ClassName != nil && v.RoomId != nil && v.RoomName != nil
		lecturerCondition := (lecturerId == "") || (lecturerId != "" && v.LecturerId != nil)
		if baseCondition && lecturerCondition {
			dateMap[v.Date] = append(dateMap[v.Date], objects.GetLectureCalendarLecture{
				LecturePlanStartTime:          utils.NullUint32Scan(v.LecturePlanStartTime),
				LecturePlanEndTime:            utils.NullUint32Scan(v.LecturePlanEndTime),
				ClassId:                       utils.NullStringScan(v.ClassId),
				ClassName:                     utils.NullStringScan(v.ClassName),
				RoomId:                        utils.NullStringScan(v.RoomId),
				RoomName:                      utils.NullStringScan(v.RoomName),
				LecturerId:                    v.LecturerId,
				LecturerName:                  v.LecturerName,
				LecturerFrontTitle:            v.LecturerFrontTitle,
				LecturerBackDegree:            v.LecturerBackDegree,
				ForeignLecturerName:           v.ForeignLecturerName,
				ForeignLecturerSourceInstance: v.ForeignLecturerSourceInstance,
			})
		} else if dateMap[v.Date] == nil {
			dateMap[v.Date] = []objects.GetLectureCalendarLecture{}
		}
	}

	dates := make([]time.Time, 0, len(dateMap))
	for k := range dateMap {
		dates = append(dates, k)
	}

	sort.Slice(dates, func(i, j int) bool {
		return dates[i].Before(dates[j])
	})

	for _, v := range dates {
		result = append(result, objects.GetLectureCalendar{
			Date:     v,
			Lectures: dateMap[v],
		})
	}

	return result
}

func (f lectureService) GetList(ctx context.Context, paginationData common.PaginationRequest, classId, semesterId string, hasActualLecture, isExam *bool, examType string) (objects.LectureListWithPagination, *constants.ErrorResponse) {
	var result objects.LectureListWithPagination

	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	modelResult, paginationResult, errs := f.LectureRepo.GetList(ctx, tx, paginationData, classId, semesterId, hasActualLecture, isExam, examType)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	var lectureIds []string
	for _, v := range modelResult {
		lectureIds = append(lectureIds, v.Id)
	}

	var examSupervisorData []models.GetExamLectureSupervisor
	if len(lectureIds) != 0 {
		examSupervisorData, errs = f.ExamSupervisorRepo.GetExamLectureSupervisorByLectureIds(ctx, tx, lectureIds)
		if errs != nil {
			_ = tx.Rollback()
			return result, errs
		}

	}

	result = objects.LectureListWithPagination{
		Pagination: paginationResult,
		Data:       mapGetList(modelResult, examSupervisorData),
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, constants.ErrUnknown
	}

	return result, nil
}

func (f lectureService) GetDetail(ctx context.Context, id string) (objects.GetLectureDetail, *constants.ErrorResponse) {
	var result objects.GetLectureDetail

	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	resultData, errs := f.LectureRepo.GetDetail(ctx, tx, id)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}
	if resultData.ClassId == nil {
		_ = tx.Rollback()
		return result, constants.ErrInvalidLecture
	}
	classId := utils.NullStringScan(resultData.ClassId)
	studentData, _, errs := f.StudentClassRepo.GetClassParticipant(ctx, tx, common.PaginationRequest{Page: constants.DefaultPage, Limit: constants.DefaultUnlimited}, []string{classId}, id, nil, "")
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	students := []objects.GetLectureDetailStudent{}
	for _, v := range studentData {
		students = append(students, objects.GetLectureDetailStudent{
			Id:        v.StudentId,
			NimNumber: v.StudentNimNumber,
			Name:      v.StudentName,
			IsAttend:  v.IsAttend,
			IsSick:    v.IsSick,
			IsLeave:   v.IsLeave,
			IsAwol:    v.IsAwol,
		})
	}

	result = objects.GetLectureDetail{
		Id:                     resultData.Id,
		LecturePlanDate:        resultData.LecturePlanDate,
		LecturePlanDayOfWeek:   resultData.LecturePlanDayOfWeek,
		LecturePlanStartTime:   resultData.LecturePlanStartTime,
		LecturePlanEndTime:     resultData.LecturePlanEndTime,
		LectureActualDate:      resultData.LectureActualDate,
		LectureActualDayOfWeek: resultData.LectureActualDayOfWeek,
		LectureActualStartTime: resultData.LectureActualStartTime,
		LectureActualEndTime:   resultData.LectureActualEndTime,
		ClassId:                resultData.ClassId,
		ClassName:              resultData.ClassName,
		SubjectId:              resultData.SubjectId,
		SubjectName:            resultData.SubjectName,
		StudyProgramId:         resultData.StudyProgramId,
		StudyProgramName:       resultData.StudyProgramName,
		SemesterId:             resultData.SemesterId,
		SemesterSchoolYear:     utils.GenerateSchoolYear(resultData.SemesterStartYear),
		SemesterType:           resultData.SemesterType,
		LectureTheme:           resultData.LectureTheme,
		LectureSubject:         resultData.LectureSubject,
		Remarks:                resultData.Remarks,
		Students:               students,
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, constants.ErrUnknown
	}

	return result, nil
}

func (f lectureService) BulkCreate(ctx context.Context, data objects.CreateLecture) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	classData, errs := f.ClassRepo.GetDetail(ctx, tx, data.ClassId, "")
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	var roomIds []string
	for _, v := range data.LecturePlans {
		roomIds = append(roomIds, v.RoomId)
	}
	roomData, errs := f.RoomRepo.GetDetailByRoomIds(ctx, tx, roomIds)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}
	roomMap := make(map[string]uint32)
	for _, v := range roomData {
		normalKey := fmt.Sprintf("%s-%t", v.Id, false)
		examKey := fmt.Sprintf("%s-%t", v.Id, true)
		roomMap[normalKey] = utils.NullUint32Scan(v.Capacity)
		roomMap[examKey] = utils.NullUint32Scan(v.ExamCapacity)
	}

	var createData []models.CreateLecture
	var examSupervisorData []models.UpsertExamLectureSupervisor
	for _, v := range data.LecturePlans {
		lectureId := uuid.New().String()

		roomValidationKey := fmt.Sprintf("%s-%t", v.RoomId, v.IsExam)
		if roomMap[roomValidationKey] != 0 && roomMap[roomValidationKey] < classData.TotalParticipant {
			_ = tx.Rollback()
			return constants.ErrClassExceedRoomCapacity
		}

		createData = append(createData, models.CreateLecture{
			Id:                   lectureId,
			ClassId:              utils.NewNullString(data.ClassId),
			LecturerId:           utils.NewNullString(v.LecturerId),
			RoomId:               v.RoomId,
			LecturePlanDate:      v.LecturePlanDate,
			LecturePlanStartTime: v.LecturePlanStartTime,
			LecturePlanEndTime:   v.LecturePlanEndTime,
			IsExam:               v.IsExam,
			IsTheoryExam:         utils.NewNullBoolean(v.IsTheoryExam),
			IsPracticumExam:      utils.NewNullBoolean(v.IsPracticumExam),
			IsFieldPracticumExam: utils.NewNullBoolean(v.IsFieldPracticumExam),
			IsMidtermExam:        utils.NewNullBoolean(v.IsMidtermExam),
			IsEndtermExam:        utils.NewNullBoolean(v.IsEndtermExam),
		})

		for _, w := range v.ExamSupervisors {
			examSupervisorData = append(examSupervisorData, models.UpsertExamLectureSupervisor{
				LectureId:            lectureId,
				ExamSupervisorId:     w.ExamSupervisorId,
				ExamSupervisorRoleId: w.ExamSupervisorRoleId,
				CreatedBy:            claims.ID,
			})

		}
	}

	if len(createData) != 0 {
		errs := f.LectureRepo.BulkCreate(ctx, tx, createData)
		if errs != nil {
			_ = tx.Rollback()
			return errs
		}
	}
	if len(examSupervisorData) != 0 {
		errs = f.ExamSupervisorRepo.UpsertExamLectureSupervisor(ctx, tx, examSupervisorData)
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

func (f lectureService) Update(ctx context.Context, data objects.UpdateLecture) *constants.ErrorResponse {
	now := time.Now().Local()

	// claims, errs := utils.GetJWTClaimsFromContext(ctx)
	// if errs != nil {
	// 	return errs
	// }

	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	lectureData, errs := f.LectureRepo.GetDetail(ctx, tx, data.Id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	if lectureData.AutonomousParticipationStartTime != nil {
		if now.After(*lectureData.AutonomousParticipationStartTime) {
			_ = tx.Rollback()
			return constants.ErrUneditableLecture
		}
	}

	if data.RoomId != "" {
		if lectureData.ClassId == nil {
			_ = tx.Rollback()
			return constants.ErrInvalidLecture
		}
		classData, errs := f.ClassRepo.GetDetail(ctx, tx, utils.NullStringScan(lectureData.ClassId), "")
		if errs != nil {
			_ = tx.Rollback()
			return errs
		}
		roomData, errs := f.RoomRepo.GetDetail(ctx, tx, data.RoomId)
		if errs != nil {
			_ = tx.Rollback()
			return errs
		}

		if lectureData.IsExam && utils.NullUint32Scan(roomData.ExamCapacity) != 0 {
			if classData.TotalParticipant > utils.NullUint32Scan(roomData.ExamCapacity) {
				_ = tx.Rollback()
				return constants.ErrClassExceedRoomCapacity
			}
		}
		if !lectureData.IsExam && utils.NullUint32Scan(roomData.Capacity) != 0 {
			if classData.TotalParticipant > utils.NullUint32Scan(roomData.Capacity) {
				_ = tx.Rollback()
				return constants.ErrClassExceedRoomCapacity
			}
		}
	}

	// var autonomousParticipationStartTime sql.NullTime
	// var autonomousParticipationEndTime sql.NullTime
	// if claims.Role == constants.AppTypeLecturer {
	// 	data.LecturerId = claims.ID
	// 	timeNum := fmt.Sprintf("%d%02d", now.Hour(), now.Minute())
	// 	lectureActualStartTime, err := strconv.Atoi(timeNum)
	// 	if err != nil {
	// 		_ = tx.Rollback()
	// 		return constants.ErrorInternalServer(err.Error())
	// 	}
	// 	data.LectureActualDate = now
	// 	data.LectureActualStartTime = uint32(lectureActualStartTime)

	// 	autonomousParticipationEndTime = utils.NewNullTime(data.AutonomousParticipationEndTime)
	// 	if !data.IsManualParticipation && autonomousParticipationEndTime.Valid {
	// 		autonomousParticipationStartTime = utils.NewNullTime(now)
	// 	}
	// }

	lecturerId := utils.NewNullString(data.LecturerId)
	foreignLecturerName := utils.NewNullString(data.ForeignLecturerName)
	foreignLecturerSourceInstance := utils.NewNullString(data.ForeignLecturerSourceInstance)
	// lectureActualDate := utils.NewNullTime(data.LectureActualDate)
	// lectureActualStartTime := utils.NewNullInt32(int32(data.LectureActualStartTime))
	// lectureActualEndTime := utils.NewNullInt32(int32(data.LectureActualEndTime))
	lectureTheme := utils.NewNullString(data.LectureTheme)
	lectureSubject := utils.NewNullString(data.LectureSubject)
	remarks := utils.NewNullString(data.Remarks)

	if lectureData.LectureActualDate != nil {
		lecturerId = utils.NewNullString(utils.NullStringScan(lectureData.LecturerId))
		foreignLecturerName = utils.NewNullString(utils.NullStringScan(lectureData.ForeignLecturerName))
		foreignLecturerSourceInstance = utils.NewNullString(utils.NullStringScan(lectureData.ForeignLecturerSourceInstance))
		// lectureActualDate = utils.NewNullTime(*lectureData.LectureActualDate)
		// lectureActualStartTime = utils.NewNullInt32(int32(utils.NullUint32Scan(lectureData.LectureActualStartTime)))
		// lectureActualEndTime = utils.NewNullInt32(int32(utils.NullUint32Scan(lectureData.LectureActualEndTime)))
	}

	var isOriginalLecturer sql.NullBool
	if lecturerId.Valid {
		if lectureData.ClassId == nil {
			_ = tx.Rollback()
			return constants.ErrInvalidLecture
		}
		classLecturerData, errs := f.ClassLecturerRepo.GetByClassIdLecturerId(ctx, tx, utils.NullStringScan(lectureData.ClassId), lecturerId.String)
		if errs != nil {
			_ = tx.Rollback()
			return errs
		}
		isOriginalLecturer = utils.NewNullBoolean(classLecturerData.Id != "")
	}

	// if lectureActualDate.Valid && len(data.Participants) != 0 && data.IsManualParticipation {
	// 	attendanceMap := make(map[string]models.UpdateLectureParticipant)

	// 	for _, v := range data.Participants {
	// 		attendanceMap[v.StudentId] = models.UpdateLectureParticipant{
	// 			StudentId: v.StudentId,
	// 			IsAttend:  utils.NewNullBoolean(v.IsAttend),
	// 			IsSick:    utils.NewNullBoolean(v.IsSick),
	// 			IsLeave:   utils.NewNullBoolean(v.IsLeave),
	// 			IsAwol:    utils.NewNullBoolean(v.IsAwol),
	// 		}
	// 	}

	// 	participantData, errs := f.LectureRepo.GetParticipantByLectureId(ctx, tx, data.Id)
	// 	if err != nil {
	// 		_ = tx.Rollback()
	// 		return errs
	// 	}

	// 	updateParticipantData := []models.UpdateLectureParticipant{}
	// 	for _, v := range participantData {
	// 		if attendanceMap[v.StudentId].StudentId == "" {
	// 			continue
	// 		}
	// 		updateParticipantData = append(updateParticipantData, models.UpdateLectureParticipant{
	// 			LectureId: v.LectureId,
	// 			StudentId: v.StudentId,
	// 			IsAttend:  attendanceMap[v.StudentId].IsAttend,
	// 			IsSick:    attendanceMap[v.StudentId].IsSick,
	// 			IsLeave:   attendanceMap[v.StudentId].IsLeave,
	// 			IsAwol:    attendanceMap[v.StudentId].IsAwol,
	// 		})
	// 	}

	// 	if len(updateParticipantData) != 0 {
	// 		errs = f.LectureRepo.BulkUpdateParticipant(ctx, tx, updateParticipantData)
	// 		if errs != nil {
	// 			_ = tx.Rollback()
	// 			return errs
	// 		}
	// 	}
	// }

	updateData := models.UpdateLecture{
		Id:                            data.Id,
		RoomId:                        utils.NewNullString(data.RoomId),
		LecturePlanDate:               utils.NewNullTime(data.LecturePlanDate),
		LecturePlanStartTime:          utils.NewNullInt32(int32(data.LecturePlanStartTime)),
		LecturePlanEndTime:            utils.NewNullInt32(int32(data.LecturePlanEndTime)),
		LecturerId:                    lecturerId,
		IsOriginalLecturer:            isOriginalLecturer,
		ForeignLecturerName:           foreignLecturerName,
		ForeignLecturerSourceInstance: foreignLecturerSourceInstance,
		// LectureActualDate:                lectureActualDate,
		// LectureActualStartTime:           lectureActualStartTime,
		// LectureActualEndTime:             lectureActualEndTime,
		LectureTheme:   lectureTheme,
		LectureSubject: lectureSubject,
		Remarks:        remarks,
		// IsManualParticipation:            utils.NewNullBoolean(data.IsManualParticipation),
		// AutonomousParticipationStartTime: autonomousParticipationStartTime,
		// AutonomousParticipationEndTime:   autonomousParticipationEndTime,
	}
	errs = f.LectureRepo.Update(ctx, tx, updateData)
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

func (f lectureService) Delete(ctx context.Context, id string) *constants.ErrorResponse {
	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs := f.LectureRepo.GetDetail(ctx, tx, id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	errs = f.LectureRepo.Delete(ctx, tx, id)
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

func (f lectureService) ResetParticipation(ctx context.Context, id string) *constants.ErrorResponse {
	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs := f.LectureRepo.GetDetail(ctx, tx, id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	errs = f.LectureRepo.ResetParticipation(ctx, tx, id)
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

func (f lectureService) GetStudentParticipation(ctx context.Context, paginationData common.PaginationRequest, classId, studentId string) (objects.LectureParticipationWithPagination, *constants.ErrorResponse) {
	var result objects.LectureParticipationWithPagination

	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	modelResult, paginationResult, errs := f.LectureRepo.GetStudentParticipation(ctx, tx, paginationData, classId, studentId)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	resultData := []objects.GetLectureParticipation{}
	for _, v := range modelResult {
		resultData = append(resultData, objects.GetLectureParticipation{
			Id:                     v.Id,
			LecturePlanDate:        v.LecturePlanDate,
			LecturePlanDayOfWeek:   v.LecturePlanDayOfWeek,
			LecturePlanStartTime:   v.LecturePlanStartTime,
			LecturePlanEndTime:     v.LecturePlanEndTime,
			LectureActualDate:      v.LectureActualDate,
			LectureActualDayOfWeek: v.LectureActualDayOfWeek,
			LectureActualStartTime: v.LectureActualStartTime,
			LectureActualEndTime:   v.LectureActualEndTime,
			IsAttend:               v.IsAttend,
			IsSick:                 v.IsSick,
			IsLeave:                v.IsLeave,
			IsAwol:                 v.IsAwol,
		})
	}

	result = objects.LectureParticipationWithPagination{
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

func (f lectureService) AttendAutonomousLecture(ctx context.Context, lectureId string) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs = f.LectureRepo.GetDetail(ctx, tx, lectureId)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	errs = f.LectureRepo.AttendLecture(ctx, tx, lectureId, claims.ID)
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

func (a lectureService) GetDetailByClassId(ctx context.Context, classId, appType string) (objects.GetDetailLecture, *constants.ErrorResponse) {
	var results objects.GetDetailLecture

	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return results, constants.ErrUnknown
	}

	resultData, errs := a.ClassRepo.GetDetail(ctx, tx, classId, "")
	if errs != nil {
		_ = tx.Rollback()
		return results, errs
	}

	prerequisiteData, errs := a.SubjectPrerequisiteRepo.GetBySubjectId(ctx, tx, resultData.SubjectId)
	if errs != nil {
		_ = tx.Rollback()
		return results, errs
	}

	prerequisiteSubjects := []string{}
	for _, v := range prerequisiteData {
		prerequisiteSubjects = append(prerequisiteSubjects, v.PrerequisiteSubjectName)
	}

	lectureData, errs := a.LectureRepo.GetByClassIds(ctx, tx, []string{classId})
	if errs != nil {
		_ = tx.Rollback()
		return results, errs
	}

	lectures := []objects.GetDetailLectureLecture{}
	for _, v := range lectureData {
		lectures = append(lectures, objects.GetDetailLectureLecture{
			LecturePlanDate: v.LecturePlanDate,
			DayOfWeek:       v.LecturePlanDayOfWeek,
			StartTime:       v.LecturePlanStartTime,
			EndTime:         v.LecturePlanEndTime,
			RoomName:        v.RoomName,
			IsMidtermExam:   v.IsMidtermExam,
			IsEndtermExam:   v.IsEndtermExam,
		})
	}

	results = objects.GetDetailLecture{
		StudyProgramName:                resultData.StudyProgramName,
		SubjectCode:                     resultData.SubjectCode,
		SubjectName:                     resultData.SubjectName,
		SemesterPackage:                 resultData.SemesterPackage,
		TheoryCredit:                    resultData.TheoryCredit,
		PracticumCredit:                 resultData.PracticumCredit,
		FieldPracticumCredit:            resultData.FieldPracticumCredit,
		SubjectMinimumPassingGradePoint: resultData.SubjectMinimumPassingGradePoint,
		SubjectIsMandatory:              resultData.SubjectIsMandatory,
		MaximumParticipant:              resultData.MaximumParticipant,
		PrerequisiteSubjects:            prerequisiteSubjects,
		Lectures:                        lectures,
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return results, constants.ErrUnknown
	}

	return results, nil
}

func (a lectureService) GetLectureCalendar(ctx context.Context, req objects.GetLectureCalendarRequest) ([]objects.GetLectureCalendar, *constants.ErrorResponse) {
	var result []objects.GetLectureCalendar

	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	resultData, errs := a.LectureRepo.GetLectureCalendar(ctx, tx, req)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, constants.ErrUnknown
	}

	return mapLectureCalendar(resultData, req.LecturerId), nil
}

func (a lectureService) GetHistory(ctx context.Context, paginationData common.PaginationRequest, studentId string, startDate, endDate time.Time) (objects.GetLectureHistoryWithPagination, *constants.ErrorResponse) {
	var result objects.GetLectureHistoryWithPagination

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return result, errs
	}

	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	if claims.Role == constants.AppTypeStudent {
		studentId = claims.ID
	}

	modelResult, paginationResult, errs := a.LectureRepo.GetHistory(ctx, tx, paginationData, studentId, startDate, endDate)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	var resultData []objects.GetLectureHistory
	for _, v := range modelResult {
		resultData = append(resultData, objects.GetLectureHistory{
			Id:          v.Id,
			LectureDate: v.LectureDate,
			SubjectName: v.SubjectName,
			AttendTime:  v.AttendTime,
		})
	}

	result = objects.GetLectureHistoryWithPagination{
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
