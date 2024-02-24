package room

import (
	"fmt"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/objects"
)

func mapGetList(roomData []models.GetRoom) []objects.GetRoom {
	results := []objects.GetRoom{}

	for _, v := range roomData {
		results = append(results, objects.GetRoom{
			Id:           v.Id,
			Code:         v.Code,
			Name:         v.Name,
			Capacity:     v.Capacity,
			IsLaboratory: v.IsLaboratory,
		})
	}

	return results
}

func mapGetSchedule(data []models.GetRoomSchedule, details []models.GetRoomScheduleDetail) []objects.GetRoomSchedule {
	results := []objects.GetRoomSchedule{}

	scheduleData := make(map[string][]objects.GetRoomScheduleDateSchedule)
	dateData := make(map[string][]objects.GetRoomScheduleDate)

	for _, v := range details {
		scheduleKey := fmt.Sprintf("%s-%s", v.RoomId, v.LecturePlanDate.Format(constants.DateRFC))

		scheduleData[scheduleKey] = append(scheduleData[scheduleKey], objects.GetRoomScheduleDateSchedule{
			RoomId:           v.RoomId,
			Date:             v.LecturePlanDate,
			StartTime:        v.StartTime,
			EndTime:          v.EndTime,
			SubjectName:      v.SubjectName,
			ClassName:        v.ClassName,
			StudyProgramName: v.StudyProgramName,
		})
	}

	checkDate := make(map[string]bool)
	for _, v := range details {
		scheduleKey := fmt.Sprintf("%s-%s", v.RoomId, v.LecturePlanDate.Format(constants.DateRFC))
		if scheduleData[scheduleKey] == nil {
			scheduleData[scheduleKey] = []objects.GetRoomScheduleDateSchedule{}
		}

		if !checkDate[scheduleKey] {
			checkDate[scheduleKey] = true
			dateData[v.RoomId] = append(dateData[v.RoomId], objects.GetRoomScheduleDate{
				RoomId:    v.RoomId,
				Date:      v.LecturePlanDate,
				Schedules: scheduleData[scheduleKey],
			})
		}
	}

	for _, v := range data {
		if dateData[v.RoomId] == nil {
			dateData[v.RoomId] = []objects.GetRoomScheduleDate{}
		}

		results = append(results, objects.GetRoomSchedule{
			RoomId:   v.RoomId,
			RoomName: v.RoomName,
			Dates:    dateData[v.RoomId],
		})
	}

	return results
}
