package objects

import (
	"github.com/sccicitb/pupr-backend/objects/common"
)

type GetClassDiscussion struct {
	Id                 string
	Title              string
	Abstraction        string
	LecturerId         string
	LecturerName       string
	LecturerFrontTitle *string
	LecturerBackDegree *string
	TotalComment       uint32
	LastComment        *string
}

type ClassDiscussionListWithPagination struct {
	Pagination common.Pagination
	Data       []GetClassDiscussion
}

type GetClassDiscussionComment struct {
	Id                 string
	StudentId          *string
	StudentNimNumber   *int64
	StudentName        *string
	LecturerId         *string
	LecturerName       *string
	LecturerFrontTitle *string
	LecturerBackDegree *string
	Comment            string
	SelfComment        bool
}

type ClassDiscussionCommentWithPagination struct {
	Pagination common.Pagination
	Data       []GetClassDiscussionComment
}

type CreateClassDiscussion struct {
	ClassId     string
	Title       string
	Abstraction string
}

type UpdateClassDiscussion struct {
	Id          string
	Title       string
	Abstraction string
}

type CreateClassDiscussionComment struct {
	ClassDiscussionId string
	Comment           string
}
