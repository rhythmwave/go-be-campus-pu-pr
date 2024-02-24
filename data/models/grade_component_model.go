package models

type GetGradeComponent struct {
	Id                  string  `db:"id"`
	StudyProgramId      string  `db:"study_program_id"`
	StudyProgramName    string  `db:"study_program_name"`
	SubjectCategoryId   string  `db:"subject_category_id"`
	SubjectCategoryName string  `db:"subject_category_name"`
	Name                string  `db:"name"`
	IsActive            bool    `db:"is_active"`
	DefaultPercentage   float64 `db:"default_percentage"`
}

type CreateGradeComponent struct {
	StudyProgramId    string `db:"study_program_id"`
	SubjectCategoryId string `db:"subject_category_id"`
	Name              string `db:"name"`
	IsActive          bool   `db:"is_active"`
	CreatedBy         string `db:"created_by"`
}

type UpdateGradeComponent struct {
	Id                string `db:"id"`
	SubjectCategoryId string `db:"subject_category_id"`
	Name              string `db:"name"`
	IsActive          bool   `db:"is_active"`
	UpdatedBy         string `db:"updated_by"`
}

type GetGradeComponentDistinctSubjectCategory struct {
	StudyProgramId      string `db:"study_program_id"`
	StudyProgramName    string `db:"study_program_name"`
	SubjectCategoryId   string `db:"subject_category_id"`
	SubjectCategoryName string `db:"subject_category_name"`
}

type GetPercentageBySubjectCategories struct {
	Id                string  `db:"id"`
	SubjectCategoryId string  `db:"subject_category_id"`
	Name              string  `db:"name"`
	DefaultPercentage float64 `db:"default_percentage"`
	IsActive          bool    `db:"is_active"`
}

type BulkUpdateGradeComponentPercentage struct {
	Id                string  `db:"id"`
	StudyProgramId    string  `db:"study_program_id"`
	SubjectCategoryId string  `db:"subject_category_id"`
	Name              string  `db:"name"`
	DefaultPercentage float64 `db:"default_percentage"`
	IsActive          bool    `db:"is_active"`
	UpdatedBy         string  `db:"updated_by"`
}
