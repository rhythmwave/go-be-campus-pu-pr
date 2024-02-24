package student_class

import (
	"math"

	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/objects"
)

func mapGetList(data []models.GetStudentClass, scheduleData []models.GetLectureDetail) []objects.GetStudentClass {
	resultData := []objects.GetStudentClass{}

	schedules := []objects.GetStudentClassSchedule{}
	for _, v := range scheduleData {
		schedules = append(schedules, objects.GetStudentClassSchedule{
			Date:      v.LecturePlanDate,
			StartTime: v.LecturePlanStartTime,
			EndTime:   v.LecturePlanEndTime,
			RoomId:    v.RoomId,
			RoomName:  v.RoomName,
		})
	}

	for _, v := range data {
		var attendancePercentage float64
		if v.TotalLectureDone != 0 {
			attendancePercentage = math.Round((float64(v.TotalAttendance)/float64(v.TotalLectureDone))*10000) / 100
		}

		resultData = append(resultData, objects.GetStudentClass{
			Id:                          v.Id,
			ClassId:                     v.ClassId,
			ClassName:                   v.ClassName,
			SubjectId:                   v.SubjectId,
			SubjectCode:                 v.SubjectCode,
			SubjectName:                 v.SubjectName,
			SubjectTheoryCredit:         v.SubjectTheoryCredit,
			SubjectPracticumCredit:      v.SubjectPracticumCredit,
			SubjectFieldPracticumCredit: v.SubjectFieldPracticumCredit,
			SubjectRepetition:           v.SubjectRepetition,
			SubjectIsMandatory:          v.SubjectIsMandatory,
			TotalAttendance:             v.TotalAttendance,
			TotalSick:                   v.TotalSick,
			TotalLeave:                  v.TotalLeave,
			TotalAwol:                   v.TotalAwol,
			GradePoint:                  v.GradePoint,
			GradeCode:                   v.GradeCode,
			GradedByAdminId:             v.GradedByAdminId,
			GradedByAdminName:           v.GradedByAdminName,
			GradedByLecturerId:          v.GradedByLecturerId,
			GradedByLecturerName:        v.GradedByLecturerName,
			GradedAt:                    v.GradedAt,
			AttendancePercentage:        attendancePercentage,
			TotalLecture:                v.TotalLectureDone,
			Schedules:                   schedules,
		})
	}

	return resultData
}
