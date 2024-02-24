package objects

import (
	"time"

	"github.com/sccicitb/pupr-backend/objects/common"
)

type GetRoomRequest struct {
	BuildingId            string
	IsLaboratory          *bool
	ExcludeLectureDate    time.Time
	ExcludeStartTime      uint32
	ExcludeEndTime        uint32
	MaximumParticipant    uint32
	ForExam               bool
	ForceIncludeLectureId string
}

type GetRoom struct {
	Id           string
	Code         string
	Name         *string
	Capacity     *uint32
	IsLaboratory bool
}

type RoomListWithPagination struct {
	Pagination common.Pagination
	Data       []GetRoom
}

type GetRoomDetail struct {
	Id               string
	BuildingId       string
	BuildingCode     string
	BuildingName     string
	Code             string
	Name             *string
	Capacity         *uint32
	ExamCapacity     *uint32
	Purpose          string
	IsUsable         bool
	Area             *float64
	PhoneNumber      *string
	Facility         *string
	Remarks          *string
	Owner            *string
	Location         *string
	StudyProgramId   *string
	StudyProgramName *string
	IsLaboratory     bool
}

type CreateRoom struct {
	BuildingId     string
	Code           string
	Name           string
	Capacity       uint32
	ExamCapacity   uint32
	IsUsable       bool
	Area           float64
	PhoneNumber    string
	Facility       string
	Remarks        string
	Purpose        string
	Owner          string
	Location       string
	StudyProgramId string
	IsLaboratory   bool
}

type UpdateRoom struct {
	Id             string
	Code           string
	Name           string
	Capacity       uint32
	ExamCapacity   uint32
	IsUsable       bool
	Area           float64
	PhoneNumber    string
	Facility       string
	Remarks        string
	Purpose        string
	Owner          string
	Location       string
	StudyProgramId string
}

type GetRoomScheduleDateSchedule struct {
	RoomId           string
	Date             time.Time
	StartTime        uint32
	EndTime          uint32
	SubjectName      string
	ClassName        string
	StudyProgramName string
}

type GetRoomScheduleDate struct {
	RoomId    string
	Date      time.Time
	Schedules []GetRoomScheduleDateSchedule
}

type GetRoomSchedule struct {
	RoomId   string
	RoomName *string
	Dates    []GetRoomScheduleDate
}

type RoomScheduleWithPagination struct {
	Pagination common.Pagination
	Data       []GetRoomSchedule
}
