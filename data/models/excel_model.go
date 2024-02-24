package models

type StudyProgramDistributionDownload struct {
	Id               string  `db:"id"`
	StudentNimNumber string  `db:"student_nim_number"`
	StudentName      string  `db:"student_name"`
	SubjectCode      string  `db:"subject_code"`
	SubjectName      string  `db:"subject_name"`
	ClassName        string  `db:"class_name"`
	GradeCode        *string `db:"grade_code"`
	TotalAttendance  int64   `db:"total_attendance"`
	TotalLectureDone int64   `db:"total_lecture_done"`
	Percentage       float64 `db:"percentage"`
}
