package models

import (
	"database/sql"
	"time"
)

type GetLectureList struct {
	Id                               string     `db:"id"`
	LecturePlanDate                  time.Time  `db:"lecture_plan_date"`
	LecturePlanDayOfWeek             uint32     `db:"lecture_plan_day_of_week"`
	LecturePlanStartTime             uint32     `db:"lecture_plan_start_time"`
	LecturePlanEndTime               uint32     `db:"lecture_plan_end_time"`
	LectureActualDate                *time.Time `db:"lecture_actual_date"`
	LectureActualDayOfWeek           *uint32    `db:"lecture_actual_day_of_week"`
	LectureActualStartTime           *uint32    `db:"lecture_actual_start_time"`
	LectureActualEndTime             *uint32    `db:"lecture_actual_end_time"`
	LecturerId                       *string    `db:"lecturer_id"`
	LecturerName                     *string    `db:"lecturer_name"`
	ForeignLecturerName              *string    `db:"foreign_lecturer_name"`
	ForeignLecturerSourceInstance    *string    `db:"foreign_lecturer_source_instance"`
	IsOriginalLecturer               *bool      `db:"is_original_lecturer"`
	IsManualParticipation            *bool      `db:"is_manual_participation"`
	AutonomousParticipationStartTime *time.Time `db:"autonomous_participation_start_time"`
	AutonomousParticipationEndTime   *time.Time `db:"autonomous_participation_end_time"`
	AttendingParticipant             uint32     `db:"attending_participant"`
	ClassId                          *string    `db:"class_id"`
	ClassName                        *string    `db:"class_name"`
	RoomId                           string     `db:"room_id"`
	RoomName                         *string    `db:"room_name"`
	BuildingId                       string     `db:"building_id"`
	BuildingName                     string     `db:"building_name"`
	IsMidtermExam                    *bool      `db:"is_midterm_exam"`
	IsEndtermExam                    *bool      `db:"is_endterm_exam"`
	IsTheoryExam                     *bool      `db:"is_theory_exam"`
	IsPracticumExam                  *bool      `db:"is_practicum_exam"`
	IsFieldPracticumExam             *bool      `db:"is_field_practicum_exam"`
	SubjectCode                      string     `db:"subject_code"`
	SubjectName                      string     `db:"subject_name"`
	TotalParticipant                 uint32     `db:"total_participant"`
	UpdatedAt                        *time.Time `db:"updated_at"`
}

type GetLectureDetail struct {
	Id                               string     `db:"id"`
	ClassId                          *string    `db:"class_id"`
	LecturePlanDate                  time.Time  `db:"lecture_plan_date"`
	LecturePlanDayOfWeek             uint32     `db:"lecture_plan_day_of_week"`
	LecturePlanStartTime             uint32     `db:"lecture_plan_start_time"`
	LecturePlanEndTime               uint32     `db:"lecture_plan_end_time"`
	LectureActualDate                *time.Time `db:"lecture_actual_date"`
	LectureActualDayOfWeek           *uint32    `db:"lecture_actual_day_of_week"`
	LectureActualStartTime           *uint32    `db:"lecture_actual_start_time"`
	LectureActualEndTime             *uint32    `db:"lecture_actual_end_time"`
	LecturerId                       *string    `db:"lecturer_id"`
	ForeignLecturerName              *string    `db:"foreign_lecturer_name"`
	ForeignLecturerSourceInstance    *string    `db:"foreign_lecturer_source_instance"`
	LectureTheme                     *string    `db:"lecture_theme"`
	LectureSubject                   *string    `db:"lecture_subject"`
	Remarks                          *string    `db:"remarks"`
	ClassName                        *string    `db:"class_name"`
	SubjectId                        string     `db:"subject_id"`
	SubjectName                      string     `db:"subject_name"`
	StudyProgramId                   string     `db:"study_program_id"`
	StudyProgramName                 string     `db:"study_program_name"`
	SemesterId                       string     `db:"semester_id"`
	SemesterStartYear                uint32     `db:"semester_start_year"`
	SemesterType                     string     `db:"semester_type"`
	RoomId                           string     `db:"room_id"`
	IsExam                           bool       `db:"is_exam"`
	RoomName                         *string    `db:"room_name"`
	IsMidtermExam                    *bool      `db:"is_midterm_exam"`
	IsEndtermExam                    *bool      `db:"is_endterm_exam"`
	IsTheoryExam                     *bool      `db:"is_theory_exam"`
	IsPracticumExam                  *bool      `db:"is_practicum_exam"`
	IsFieldPracticumExam             *bool      `db:"is_field_practicum_exam"`
	AutonomousParticipationStartTime *time.Time `db:"autonomous_participation_start_time"`
	AutonomousParticipationEndTime   *time.Time `db:"autonomous_participation_end_time"`
}

type CreateLecture struct {
	Id                   string         `db:"id"`
	ClassId              sql.NullString `db:"class_id"`
	RoomId               string         `db:"room_id"`
	LecturerId           sql.NullString `db:"lecturer_id"`
	LecturePlanDate      time.Time      `db:"lecture_plan_date"`
	LecturePlanStartTime uint32         `db:"lecture_plan_start_time"`
	LecturePlanEndTime   uint32         `db:"lecture_plan_end_time"`
	IsExam               bool           `db:"is_exam"`
	IsTheoryExam         sql.NullBool   `db:"is_theory_exam"`
	IsPracticumExam      sql.NullBool   `db:"is_practicum_exam"`
	IsFieldPracticumExam sql.NullBool   `db:"is_field_practicum_exam"`
	IsMidtermExam        sql.NullBool   `db:"is_midterm_exam"`
	IsEndtermExam        sql.NullBool   `db:"is_endterm_exam"`
}

type UpdateLecture struct {
	Id                            string         `db:"id"`
	RoomId                        sql.NullString `db:"room_id"`
	LecturePlanDate               sql.NullTime   `db:"lecture_plan_date"`
	LecturePlanStartTime          sql.NullInt32  `db:"lecture_plan_start_time"`
	LecturePlanEndTime            sql.NullInt32  `db:"lecture_plan_end_time"`
	LecturerId                    sql.NullString `db:"lecturer_id"`
	ForeignLecturerName           sql.NullString `db:"foreign_lecturer_name"`
	ForeignLecturerSourceInstance sql.NullString `db:"foreign_lecturer_source_instance"`
	IsOriginalLecturer            sql.NullBool   `db:"is_original_lecturer"`
	LectureTheme                  sql.NullString `db:"lecture_theme"`
	LectureSubject                sql.NullString `db:"lecture_subject"`
	Remarks                       sql.NullString `db:"remarks"`
}

type GetLectureParticipant struct {
	Id        string `db:"id"`
	LectureId string `db:"lecture_id"`
	StudentId string `db:"student_id"`
	IsAttend  *bool  `db:"is_attend"`
	IsSick    *bool  `db:"is_sick"`
	IsLeave   *bool  `db:"is_leave"`
	IsAwol    *bool  `db:"is_awol"`
}

type UpdateLectureParticipant struct {
	LectureId string       `db:"lecture_id"`
	StudentId string       `db:"student_id"`
	IsAttend  sql.NullBool `db:"is_attend"`
	IsSick    sql.NullBool `db:"is_sick"`
	IsLeave   sql.NullBool `db:"is_leave"`
	IsAwol    sql.NullBool `db:"is_awol"`
}

type GetLectureParticipation struct {
	Id                     string     `db:"id"`
	LecturePlanDate        time.Time  `db:"lecture_plan_date"`
	LecturePlanDayOfWeek   uint32     `db:"lecture_plan_day_of_week"`
	LecturePlanStartTime   uint32     `db:"lecture_plan_start_time"`
	LecturePlanEndTime     uint32     `db:"lecture_plan_end_time"`
	LectureActualDate      *time.Time `db:"lecture_actual_date"`
	LectureActualDayOfWeek *uint32    `db:"lecture_actual_day_of_week"`
	LectureActualStartTime *uint32    `db:"lecture_actual_start_time"`
	LectureActualEndTime   *uint32    `db:"lecture_actual_end_time"`
	IsAttend               *bool      `db:"is_attend"`
	IsSick                 *bool      `db:"is_sick"`
	IsLeave                *bool      `db:"is_leave"`
	IsAwol                 *bool      `db:"is_awol"`
}

type GetLectureCalendar struct {
	Date                          time.Time `db:"date"`
	LectureId                     *string   `db:"lecture_id"`
	LecturePlanStartTime          *uint32   `db:"lecture_plan_start_time"`
	LecturePlanEndTime            *uint32   `db:"lecture_plan_end_time"`
	ClassId                       *string   `db:"class_id"`
	ClassName                     *string   `db:"class_name"`
	RoomId                        *string   `db:"room_id"`
	RoomName                      *string   `db:"room_name"`
	LecturerId                    *string   `db:"lecturer_id"`
	LecturerName                  *string   `db:"lecturer_name"`
	LecturerFrontTitle            *string   `db:"lecturer_front_title"`
	LecturerBackDegree            *string   `db:"lecturer_back_degree"`
	ForeignLecturerName           *string   `db:"foreign_lecturer_name"`
	ForeignLecturerSourceInstance *string   `db:"foreign_lecturer_source_instance"`
}

type GetLectureHistory struct {
	Id          string    `db:"id"`
	LectureDate time.Time `db:"lecture_date"`
	SubjectName string    `db:"subject_name"`
	AttendTime  time.Time `db:"attend_time"`
}
