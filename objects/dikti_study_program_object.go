package objects

import "github.com/sccicitb/pupr-backend/objects/common"

type GetDiktiStudyProgram struct {
	Id                  string
	Code                string
	Name                string
	StudyLevelShortName string
	StudyLevelName      string
}

type DiktiStudyProgramListWithPagination struct {
	Pagination common.Pagination
	Data       []GetDiktiStudyProgram
}
