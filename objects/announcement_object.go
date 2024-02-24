package objects

import (
	"time"

	"github.com/sccicitb/pupr-backend/objects/common"
)

type GetAnnouncementStudyProgram struct {
	StudyProgramId   string
	StudyProgramName string
}

type GetAnnouncement struct {
	Id               string
	Type             string
	Title            string
	AnnouncementDate *time.Time
	FileUrl          string
	FilePath         *string
	FilePathType     *string
	FileTitle        *string
	Content          *string
	ForLecturer      bool
	ForStudent       bool
	StudyPrograms    []GetAnnouncementStudyProgram
}

type AnnouncementListWithPagination struct {
	Pagination common.Pagination
	Data       []GetAnnouncement
}

type CreateAnnouncement struct {
	Type             string
	Title            string
	AnnouncementDate time.Time
	FilePath         string
	FilePathType     string
	FileTitle        string
	Content          string
	ForLecturer      bool
	ForStudent       bool
	StudyProgramIds  []string
}

type UpdateAnnouncement struct {
	Id               string
	Type             string
	Title            string
	AnnouncementDate time.Time
	FilePath         string
	FilePathType     string
	FileTitle        string
	Content          string
	ForLecturer      bool
	ForStudent       bool
	StudyProgramIds  []string
}
