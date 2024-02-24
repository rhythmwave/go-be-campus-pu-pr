package objects

import (
	"github.com/sccicitb/pupr-backend/objects/common"
)

type GetLessonPlan struct {
	Id            string
	MeetingOrder  uint32
	Lesson        string
	EnglishLesson string
}

type LessonPlanListWithPagination struct {
	Pagination common.Pagination
	Data       []GetLessonPlan
}

type CreateLessonPlan struct {
	SubjectId     string
	MeetingOrder  uint32
	Lesson        string
	EnglishLesson string
}

type UpdateLessonPlan struct {
	Id            string
	MeetingOrder  uint32
	Lesson        string
	EnglishLesson string
}
