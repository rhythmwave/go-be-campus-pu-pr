package models

import (
	"database/sql"
	"time"
)

type GetRoom struct {
	Id           string  `db:"id"`
	Code         string  `db:"code"`
	Name         *string `db:"name"`
	Capacity     *uint32 `db:"capacity"`
	IsLaboratory bool    `db:"is_laboratory"`
}

type GetRoomDetail struct {
	Id               string   `db:"id"`
	BuildingId       string   `db:"building_id"`
	BuildingCode     string   `db:"building_code"`
	BuildingName     string   `db:"building_name"`
	Code             string   `db:"code"`
	Name             *string  `db:"name"`
	Capacity         *uint32  `db:"capacity"`
	ExamCapacity     *uint32  `db:"exam_capacity"`
	Purpose          string   `db:"purpose"`
	IsUsable         bool     `db:"is_usable"`
	Area             *float64 `db:"area"`
	PhoneNumber      *string  `db:"phone_number"`
	Facility         *string  `db:"facility"`
	Remarks          *string  `db:"remarks"`
	Owner            *string  `db:"owner"`
	Location         *string  `db:"location"`
	StudyProgramId   *string  `db:"study_program_id"`
	StudyProgramName *string  `db:"study_program_name"`
	IsLaboratory     bool     `db:"is_laboratory"`
}

type CreateRoom struct {
	BuildingId     string          `db:"building_id"`
	Code           string          `db:"code"`
	Name           sql.NullString  `db:"name"`
	Capacity       sql.NullInt32   `db:"capacity"`
	ExamCapacity   sql.NullInt32   `db:"exam_capacity"`
	IsUsable       bool            `db:"is_usable"`
	Area           sql.NullFloat64 `db:"area"`
	PhoneNumber    sql.NullString  `db:"phone_number"`
	Facility       sql.NullString  `db:"facility"`
	Remarks        sql.NullString  `db:"remarks"`
	Purpose        string          `db:"purpose"`
	Owner          sql.NullString  `db:"owner"`
	Location       sql.NullString  `db:"location"`
	StudyProgramId sql.NullString  `db:"study_program_id"`
	IsLaboratory   bool            `db:"is_laboratory"`
	CreatedBy      string          `db:"created_by"`
}

type UpdateRoom struct {
	Id             string          `db:"id"`
	Code           string          `db:"code"`
	Name           sql.NullString  `db:"name"`
	Capacity       sql.NullInt32   `db:"capacity"`
	ExamCapacity   sql.NullInt32   `db:"exam_capacity"`
	IsUsable       bool            `db:"is_usable"`
	Area           sql.NullFloat64 `db:"area"`
	PhoneNumber    sql.NullString  `db:"phone_number"`
	Facility       sql.NullString  `db:"facility"`
	Remarks        sql.NullString  `db:"remarks"`
	Purpose        string          `db:"purpose"`
	Owner          sql.NullString  `db:"owner"`
	Location       sql.NullString  `db:"location"`
	StudyProgramId sql.NullString  `db:"study_program_id"`
	UpdatedBy      string          `db:"updated_by"`
}

type GetRoomSchedule struct {
	RoomId   string  `db:"room_id"`
	RoomName *string `db:"room_name"`
}

type GetRoomScheduleDetail struct {
	RoomId           string    `db:"room_id"`
	RoomName         *string   `db:"room_name"`
	LecturePlanDate  time.Time `db:"lecture_plan_date"`
	StartTime        uint32    `db:"start_time"`
	EndTime          uint32    `db:"end_time"`
	SubjectName      string    `db:"subject_name"`
	ClassName        string    `db:"class_name"`
	StudyProgramName string    `db:"study_program_name"`
}
