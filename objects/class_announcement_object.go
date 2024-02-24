package objects

import (
	"time"

	"github.com/sccicitb/pupr-backend/objects/common"
)

type GetClassAnnouncement struct {
	Id                 string
	Title              string
	Content            string
	LecturerId         string
	LecturerName       string
	LecturerFrontTitle *string
	LecturerBackDegree *string
	FileUrl            string
	FilePath           *string
	FilePathType       *string
	StartTime          *time.Time
	EndTime            *time.Time
}

type ClassAnnouncementListWithPagination struct {
	Pagination common.Pagination
	Data       []GetClassAnnouncement
}

type CreateClassAnnouncement struct {
	ClassId      string
	Title        string
	Content      string
	FilePath     string
	FilePathType string
	StartTime    time.Time
	EndTime      time.Time
}

type UpdateClassAnnouncement struct {
	Id           string
	Title        string
	Content      string
	FilePath     string
	FilePathType string
	StartTime    time.Time
	EndTime      time.Time
}
