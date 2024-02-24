package objects

import (
	"time"

	"github.com/sccicitb/pupr-backend/objects/common"
)

type GetSharedFile struct {
	Id                 string
	LecturerId         string
	LecturerName       string
	LecturerFrontTitle *string
	LecturerBackDegree *string
	Title              string
	FilePath           string
	FilePathType       string
	FileUrl            string
	Remarks            *string
	CreatedAt          time.Time
	IsApproved         bool
}

type SharedFileListWithPagination struct {
	Pagination common.Pagination
	Data       []GetSharedFile
}

type CreateSharedFile struct {
	Title        string
	FilePath     string
	FilePathType string
	Remarks      string
}

type UpdateSharedFile struct {
	Id      string
	Title   string
	Remarks string
}
