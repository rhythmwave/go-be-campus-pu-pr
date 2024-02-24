package models

type GetDiktiStudyProgramList struct {
	Id                  string `db:"id"`
	Code                string `db:"code"`
	Name                string `db:"name"`
	StudyLevelShortName string `db:"study_level_short_name"`
	StudyLevelName      string `db:"study_level_name"`
}
