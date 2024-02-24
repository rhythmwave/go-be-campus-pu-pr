package models

type GetYudiciumTerm struct {
	Id               string `db:"id"`
	CurriculumId     string `db:"curriculum_id"`
	CurriculumName   string `db:"curriculum_name"`
	StudyProgramId   string `db:"study_program_id"`
	StudyProgramName string `db:"study_program_name"`
	Term             string `db:"term"`
	Remarks          string `db:"remarks"`
}

type CreateYudiciumTerm struct {
	CurriculumId string `db:"curriculum_id"`
	Term         string `db:"term"`
	Remarks      string `db:"remarks"`
	CreatedBy    string `db:"created_by"`
}

type UpdateYudiciumTerm struct {
	Id        string `db:"id"`
	Term      string `db:"term"`
	Remarks   string `db:"remarks"`
	UpdatedBy string `db:"updated_by"`
}
