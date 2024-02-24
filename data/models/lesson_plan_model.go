package models

type GetLessonPlan struct {
	Id            string `db:"id"`
	MeetingOrder  uint32 `db:"meeting_order"`
	Lesson        string `db:"lesson"`
	EnglishLesson string `db:"english_lesson"`
}

type CreateLessonPlan struct {
	SubjectId     string `db:"subject_id"`
	MeetingOrder  uint32 `db:"meeting_order"`
	Lesson        string `db:"lesson"`
	EnglishLesson string `db:"english_lesson"`
	CreatedBy     string `db:"created_by"`
}

type UpdateLessonPlan struct {
	Id            string `db:"id"`
	MeetingOrder  uint32 `db:"meeting_order"`
	Lesson        string `db:"lesson"`
	EnglishLesson string `db:"english_lesson"`
	UpdatedBy     string `db:"updated_by"`
}
