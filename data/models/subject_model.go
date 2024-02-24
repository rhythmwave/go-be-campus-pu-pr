package models

import (
	"database/sql"
	"time"
)

type GetSubject struct {
	Id                            string   `db:"id"`
	StudyProgramId                string   `db:"study_program_id"`
	StudyProgramName              string   `db:"study_program_name"`
	CurriculumId                  string   `db:"curriculum_id"`
	CurriculumName                string   `db:"curriculum_name"`
	Code                          string   `db:"code"`
	Name                          string   `db:"name"`
	IsMandatory                   bool     `db:"is_mandatory"`
	SemesterPackage               uint32   `db:"semester_package"`
	TheoryCredit                  uint32   `db:"theory_credit"`
	PracticumCredit               uint32   `db:"practicum_credit"`
	FieldPracticumCredit          uint32   `db:"field_practicum_credit"`
	SubjectPrerequisiteId         *string  `db:"subject_prerequisite_id"`
	PrerequisiteType              *string  `db:"prerequisite_type"`
	PrerequisiteMinimumGradePoint *float64 `db:"prerequisite_minimum_grade_point"`
	EquivalentStudyProgramId      *string  `db:"equivalent_study_program_id"`
	EquivalentStudyProgramName    *string  `db:"equivalent_study_program_name"`
	EquivalentCurriculumId        *string  `db:"equivalent_curriculum_id"`
	EquivalentCurriculumName      *string  `db:"equivalent_curriculum_name"`
	EquivalentSubjectId           *string  `db:"equivalent_subject_id"`
	EquivalentSubjectCode         *string  `db:"equivalent_subject_code"`
	EquivalentSubjectName         *string  `db:"equivalent_subject_name"`
	SubjectCategoryId             string   `db:"subject_category_id"`
	SubjectCategoryName           string   `db:"subject_category_name"`
	IsThesis                      *bool    `db:"is_thesis"`
	IsMbkm                        bool     `db:"is_mbkm"`
	TotalLessonPlan               uint32   `db:"total_lesson_plan"`
}

type GetSubjectDetail struct {
	Id                           string     `db:"id"`
	StudyProgramId               string     `db:"study_program_id"`
	StudyProgramName             string     `db:"study_program_name"`
	CurriculumId                 string     `db:"curriculum_id"`
	CurriculumName               string     `db:"curriculum_name"`
	Code                         string     `db:"code"`
	Name                         string     `db:"name"`
	ShortName                    *string    `db:"short_name"`
	EnglishName                  *string    `db:"english_name"`
	EnglishShortName             *string    `db:"english_short_name"`
	IsMandatory                  bool       `db:"is_mandatory"`
	Trait                        string     `db:"trait"`
	Type                         *string    `db:"type"`
	SubjectCategoryId            string     `db:"subject_category_id"`
	SubjectCategoryName          string     `db:"subject_category_name"`
	CurriculumType               string     `db:"curriculum_type"`
	TheoryCredit                 uint32     `db:"theory_credit"`
	PracticumCredit              uint32     `db:"practicum_credit"`
	FieldPracticumCredit         uint32     `db:"field_practicum_credit"`
	SemesterPackage              uint32     `db:"semester_package"`
	RepeatCourseLimit            uint32     `db:"repeat_course_limit"`
	IsActive                     bool       `db:"is_active"`
	HasLectureUnit               bool       `db:"has_lecture_unit"`
	HasTeachingMaterial          bool       `db:"has_teaching_material"`
	HasLectureSummary            bool       `db:"has_lecture_summary"`
	SupportingLecturerId         *string    `db:"supporting_lecturer_id"`
	SupportingLecturerName       *string    `db:"supporting_lecturer_name"`
	StartDate                    *time.Time `db:"start_date"`
	EndDate                      *time.Time `db:"end_date"`
	MinimumPassingGradePoint     float64    `db:"minimum_passing_grade_point"`
	MinimumMandatoryCreditTaken  *uint32    `db:"minimum_mandatory_credit_taken"`
	MinimumOptionalCreditTaken   *uint32    `db:"minimum_optional_credit_taken"`
	MinimumTotalCreditTaken      *uint32    `db:"minimum_total_credit_taken"`
	MinimumMandatoryCreditPassed *uint32    `db:"minimum_mandatory_credit_passed"`
	MinimumOptionalCreditPassed  *uint32    `db:"minimum_optional_credit_passed"`
	MinimumTotalCreditPassed     *uint32    `db:"minimum_total_credit_passed"`
	MinimumGpa                   *float64   `db:"minimum_gpa"`
	Abstraction                  *string    `db:"abstraction"`
	SyllabusPath                 *string    `db:"syllabus_path"`
	SyllabusPathType             *string    `db:"syllabus_path_type"`
	IsThesis                     *bool      `db:"is_thesis"`
	IsMbkm                       bool       `db:"is_mbkm"`
}

type CreateSubject struct {
	CurriculumId                 string          `db:"curriculum_id"`
	Code                         string          `db:"code"`
	Name                         string          `db:"name"`
	ShortName                    sql.NullString  `db:"short_name"`
	EnglishName                  sql.NullString  `db:"english_name"`
	EnglishShortName             sql.NullString  `db:"english_short_name"`
	IsMandatory                  bool            `db:"is_mandatory"`
	Trait                        string          `db:"trait"`
	Type                         sql.NullString  `db:"type"`
	SubjectCategoryId            string          `db:"subject_category_id"`
	CurriculumType               string          `db:"curriculum_type"`
	TheoryCredit                 uint32          `db:"theory_credit"`
	PracticumCredit              uint32          `db:"practicum_credit"`
	FieldPracticumCredit         uint32          `db:"field_practicum_credit"`
	SemesterPackage              uint32          `db:"semester_package"`
	RepeatCourseLimit            uint32          `db:"repeat_course_limit"`
	IsActive                     bool            `db:"is_active"`
	HasLectureUnit               bool            `db:"has_lecture_unit"`
	HasTeachingMaterial          bool            `db:"has_teaching_material"`
	HasLectureSummary            bool            `db:"has_lecture_summary"`
	SupportingLecturerId         sql.NullString  `db:"supporting_lecturer_id"`
	StartDate                    sql.NullTime    `db:"start_date"`
	EndDate                      sql.NullTime    `db:"end_date"`
	MinimumPassingGradePoint     float64         `db:"minimum_passing_grade_point"`
	MinimumMandatoryCreditTaken  sql.NullInt32   `db:"minimum_mandatory_credit_taken"`
	MinimumOptionalCreditTaken   sql.NullInt32   `db:"minimum_optional_credit_taken"`
	MinimumTotalCreditTaken      sql.NullInt32   `db:"minimum_total_credit_taken"`
	MinimumMandatoryCreditPassed sql.NullInt32   `db:"minimum_mandatory_credit_passed"`
	MinimumOptionalCreditPassed  sql.NullInt32   `db:"minimum_optional_credit_passed"`
	MinimumTotalCreditPassed     sql.NullInt32   `db:"minimum_total_credit_passed"`
	MinimumGpa                   sql.NullFloat64 `db:"minimum_gpa"`
	Abstraction                  sql.NullString  `db:"abstraction"`
	SyllabusPath                 sql.NullString  `db:"syllabus_path"`
	SyllabusPathType             sql.NullString  `db:"syllabus_path_type"`
	IsThesis                     sql.NullBool    `db:"is_thesis"`
	IsMbkm                       bool            `db:"is_mbkm"`
	CreatedBy                    string          `db:"created_by"`
}

type UpdateSubject struct {
	Id                           string          `db:"id"`
	CurriculumId                 string          `db:"curriculum_id"`
	Code                         string          `db:"code"`
	Name                         string          `db:"name"`
	ShortName                    sql.NullString  `db:"short_name"`
	EnglishName                  sql.NullString  `db:"english_name"`
	EnglishShortName             sql.NullString  `db:"english_short_name"`
	IsMandatory                  bool            `db:"is_mandatory"`
	Trait                        string          `db:"trait"`
	Type                         sql.NullString  `db:"type"`
	SubjectCategoryId            string          `db:"subject_category_id"`
	CurriculumType               string          `db:"curriculum_type"`
	TheoryCredit                 uint32          `db:"theory_credit"`
	PracticumCredit              uint32          `db:"practicum_credit"`
	FieldPracticumCredit         uint32          `db:"field_practicum_credit"`
	SemesterPackage              uint32          `db:"semester_package"`
	RepeatCourseLimit            uint32          `db:"repeat_course_limit"`
	IsActive                     bool            `db:"is_active"`
	HasLectureUnit               bool            `db:"has_lecture_unit"`
	HasTeachingMaterial          bool            `db:"has_teaching_material"`
	HasLectureSummary            bool            `db:"has_lecture_summary"`
	SupportingLecturerId         sql.NullString  `db:"supporting_lecturer_id"`
	StartDate                    sql.NullTime    `db:"start_date"`
	EndDate                      sql.NullTime    `db:"end_date"`
	MinimumPassingGradePoint     float64         `db:"minimum_passing_grade_point"`
	MinimumMandatoryCreditTaken  sql.NullInt32   `db:"minimum_mandatory_credit_taken"`
	MinimumOptionalCreditTaken   sql.NullInt32   `db:"minimum_optional_credit_taken"`
	MinimumTotalCreditTaken      sql.NullInt32   `db:"minimum_total_credit_taken"`
	MinimumMandatoryCreditPassed sql.NullInt32   `db:"minimum_mandatory_credit_passed"`
	MinimumOptionalCreditPassed  sql.NullInt32   `db:"minimum_optional_credit_passed"`
	MinimumTotalCreditPassed     sql.NullInt32   `db:"minimum_total_credit_passed"`
	MinimumGpa                   sql.NullFloat64 `db:"minimum_gpa"`
	Abstraction                  sql.NullString  `db:"abstraction"`
	SyllabusPath                 sql.NullString  `db:"syllabus_path"`
	SyllabusPathType             sql.NullString  `db:"syllabus_path_type"`
	IsThesis                     sql.NullBool    `db:"is_thesis"`
	IsMbkm                       bool            `db:"is_mbkm"`
	UpdatedBy                    string          `db:"updated_by"`
}
