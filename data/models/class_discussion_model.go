package models

import (
	"database/sql"
)

type GetClassDiscussion struct {
	Id                 string  `db:"id"`
	ClassId            string  `db:"class_id"`
	Title              string  `db:"title"`
	Abstraction        string  `db:"abstraction"`
	LecturerId         string  `db:"lecturer_id"`
	LecturerName       string  `db:"lecturer_name"`
	LecturerFrontTitle *string `db:"lecturer_front_title"`
	LecturerBackDegree *string `db:"lecturer_back_degree"`
	TotalComment       uint32  `db:"total_comment"`
	LastComment        *string `db:"last_comment"`
}

type GetClassDiscussionComment struct {
	Id                 string  `db:"id"`
	StudentId          *string `db:"student_id"`
	StudentNimNumber   *int64  `db:"student_nim_number"`
	StudentName        *string `db:"student_name"`
	LecturerId         *string `db:"lecturer_id"`
	LecturerName       *string `db:"lecturer_name"`
	LecturerFrontTitle *string `db:"lecturer_front_title"`
	LecturerBackDegree *string `db:"lecturer_back_degree"`
	Comment            string  `db:"comment"`
}

type CreateClassDiscussion struct {
	LecturerId  string `db:"lecturer_id"`
	ClassId     string `db:"class_id"`
	Title       string `db:"title"`
	Abstraction string `db:"abstraction"`
}

type UpdateClassDiscussion struct {
	Id          string `db:"id"`
	Title       string `db:"title"`
	Abstraction string `db:"abstraction"`
}

type CreateClassDiscussionComment struct {
	ClassDiscussionId string         `db:"class_discussion_id"`
	LecturerId        sql.NullString `db:"lecturer_id"`
	StudentId         sql.NullString `db:"student_id"`
	Comment           string         `db:"comment"`
}
