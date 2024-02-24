package models

import (
	"database/sql"
	"time"
)

type GetAnnouncement struct {
	Id               string     `db:"id"`
	Type             string     `db:"type"`
	Title            string     `db:"title"`
	AnnouncementDate *time.Time `db:"announcement_date"`
	FilePath         *string    `db:"file_path"`
	FilePathType     *string    `db:"file_path_type"`
	FileTitle        *string    `db:"file_title"`
	Content          *string    `db:"content"`
	ForLecturer      bool       `db:"for_lecturer"`
	ForStudent       bool       `db:"for_student"`
}

type GetAnnouncementStudyProgram struct {
	AnnouncementId    string `db:"announcement_id"`
	AnnouncementTitle string `db:"announcement_title"`
	StudyProgramId    string `db:"study_program_id"`
	StudyProgramName  string `db:"study_program_name"`
}

type CreateAnnouncement struct {
	Id               string         `db:"id"`
	Type             string         `db:"type"`
	Title            string         `db:"title"`
	AnnouncementDate sql.NullTime   `db:"announcement_date"`
	FilePath         sql.NullString `db:"file_path"`
	FilePathType     sql.NullString `db:"file_path_type"`
	FileTitle        sql.NullString `db:"file_title"`
	Content          sql.NullString `db:"content"`
	ForLecturer      bool           `db:"for_lecturer"`
	ForStudent       bool           `db:"for_student"`
	CreatedBy        string         `db:"created_by"`
}

type UpdateAnnouncement struct {
	Id               string         `db:"id"`
	Type             string         `db:"type"`
	Title            string         `db:"title"`
	AnnouncementDate sql.NullTime   `db:"announcement_date"`
	FilePath         sql.NullString `db:"file_path"`
	FilePathType     sql.NullString `db:"file_path_type"`
	FileTitle        sql.NullString `db:"file_title"`
	Content          sql.NullString `db:"content"`
	ForLecturer      bool           `db:"for_lecturer"`
	ForStudent       bool           `db:"for_student"`
	UpdatedBy        string         `db:"updated_by"`
}

type UpsertAnnouncementStudyProgram struct {
	AnnouncementId string `db:"announcement_id"`
	StudyProgramId string `db:"study_program_id"`
}
