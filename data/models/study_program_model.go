package models

import (
	"database/sql"
	"time"
)

type GetStudyProgramList struct {
	Id                    string  `db:"id"`
	Name                  string  `db:"name"`
	StudyLevelShortName   string  `db:"study_level_short_name"`
	StudyLevelName        string  `db:"study_level_name"`
	DiktiStudyProgramType string  `db:"dikti_study_program_type"`
	DiktiStudyProgramCode string  `db:"dikti_study_program_code"`
	Accreditation         *string `db:"accreditation"`
	ActiveCurriculumYear  *string `db:"active_curriculum_year"`
	Degree                *string `db:"degree"`
	ShortDegree           *string `db:"short_degree"`
	EnglishDegree         *string `db:"english_degree"`
}

type GetStudyProgramByRoleIds struct {
	GetStudyProgramList
	RoleId string `db:"role_id"`
}

type GetStudyProgramDetail struct {
	Id                            string     `db:"id"`
	DiktiStudyProgramId           string     `db:"dikti_study_program_id"`
	DiktiStudyProgramCode         string     `db:"dikti_study_program_code"`
	DiktiStudyProgramName         string     `db:"dikti_study_program_name"`
	DiktiStudyProgramType         string     `db:"dikti_study_program_type"`
	StudyLevelShortName           string     `db:"study_level_short_name"`
	StudyLevelName                string     `db:"study_level_name"`
	Name                          string     `db:"name"`
	EnglishName                   *string    `db:"english_name"`
	ShortName                     *string    `db:"short_name"`
	EnglishShortName              *string    `db:"english_short_name"`
	AdministrativeUnit            *string    `db:"administrative_unit"`
	FacultyId                     string     `db:"faculty_id"`
	FacultyName                   string     `db:"faculty_name"`
	MajorId                       string     `db:"major_id"`
	MajorName                     string     `db:"major_name"`
	Address                       *string    `db:"address"`
	PhoneNumber                   *string    `db:"phone_number"`
	Fax                           *string    `db:"fax"`
	Email                         *string    `db:"email"`
	Website                       *string    `db:"website"`
	ContactPerson                 *string    `db:"contact_person"`
	CuriculumReviewFrequency      *string    `db:"curiculum_review_frequency"`
	CuriculumReviewMethod         *string    `db:"curiculum_review_method"`
	EstablishmentDate             *time.Time `db:"establishment_date"`
	IsActive                      bool       `db:"is_active"`
	StartSemester                 *string    `db:"start_semester"`
	OperationalPermitNumber       *string    `db:"operational_permit_number"`
	OperationalPermitDate         *time.Time `db:"operational_permit_date"`
	OperationalPermitDueDate      *time.Time `db:"operational_permit_due_date"`
	HeadLecturerId                *string    `db:"head_lecturer_id"`
	HeadLecturerName              *string    `db:"head_lecturer_name"`
	HeadLecturerMobilePhoneNumber *string    `db:"head_lecturer_mobile_phone_number"`
	OperatorName                  *string    `db:"operator_name"`
	OperatorPhoneNumber           *string    `db:"operator_phone_number"`
	MinimumGraduationCredit       uint32     `db:"minimum_graduation_credit"`
	MinimumThesisCredit           uint32     `db:"minimum_thesis_credit"`
}

type CreateStudyProgram struct {
	DiktiStudyProgramId      string         `db:"dikti_study_program_id"`
	Name                     string         `db:"name"`
	EnglishName              sql.NullString `db:"english_name"`
	ShortName                sql.NullString `db:"short_name"`
	EnglishShortName         sql.NullString `db:"english_short_name"`
	AdministrativeUnit       sql.NullString `db:"administrative_unit"`
	MajorId                  string         `db:"major_id"`
	Address                  sql.NullString `db:"address"`
	PhoneNumber              sql.NullString `db:"phone_number"`
	Fax                      sql.NullString `db:"fax"`
	Email                    sql.NullString `db:"email"`
	Website                  sql.NullString `db:"website"`
	ContactPerson            sql.NullString `db:"contact_person"`
	CuriculumReviewFrequency sql.NullString `db:"curiculum_review_frequency"`
	CuriculumReviewMethod    sql.NullString `db:"curiculum_review_method"`
	EstablishmentDate        sql.NullTime   `db:"establishment_date"`
	IsActive                 bool           `db:"is_active"`
	StartSemester            sql.NullString `db:"start_semester"`
	OperationalPermitNumber  sql.NullString `db:"operational_permit_number"`
	OperationalPermitDate    sql.NullTime   `db:"operational_permit_date"`
	OperationalPermitDueDate sql.NullTime   `db:"operational_permit_due_date"`
	HeadLecturerId           sql.NullString `db:"head_lecturer_id"`
	OperatorName             sql.NullString `db:"operator_name"`
	OperatorPhoneNumber      sql.NullString `db:"operator_phone_number"`
	CreatedBy                string         `db:"created_by"`
}

type UpdateStudyProgram struct {
	Id                       string         `db:"id"`
	DiktiStudyProgramId      string         `db:"dikti_study_program_id"`
	Name                     string         `db:"name"`
	EnglishName              sql.NullString `db:"english_name"`
	ShortName                sql.NullString `db:"short_name"`
	EnglishShortName         sql.NullString `db:"english_short_name"`
	AdministrativeUnit       sql.NullString `db:"administrative_unit"`
	MajorId                  string         `db:"major_id"`
	Address                  sql.NullString `db:"address"`
	PhoneNumber              sql.NullString `db:"phone_number"`
	Fax                      sql.NullString `db:"fax"`
	Email                    sql.NullString `db:"email"`
	Website                  sql.NullString `db:"website"`
	ContactPerson            sql.NullString `db:"contact_person"`
	CuriculumReviewFrequency sql.NullString `db:"curiculum_review_frequency"`
	CuriculumReviewMethod    sql.NullString `db:"curiculum_review_method"`
	EstablishmentDate        sql.NullTime   `db:"establishment_date"`
	IsActive                 bool           `db:"is_active"`
	StartSemester            sql.NullString `db:"start_semester"`
	OperationalPermitNumber  sql.NullString `db:"operational_permit_number"`
	OperationalPermitDate    sql.NullTime   `db:"operational_permit_date"`
	OperationalPermitDueDate sql.NullTime   `db:"operational_permit_due_date"`
	HeadLecturerId           sql.NullString `db:"head_lecturer_id"`
	OperatorName             sql.NullString `db:"operator_name"`
	OperatorPhoneNumber      sql.NullString `db:"operator_phone_number"`
	MinimumGraduationCredit  sql.NullInt32  `db:"minimum_graduation_credit"`
	MinimumThesisCredit      sql.NullInt32  `db:"minimum_thesis_credit"`
	UpdatedBy                string         `db:"updated_by"`
}

type UpdateDegreeStudyProgram struct {
	Id            string         `db:"id"`
	Degree        sql.NullString `db:"degree"`
	ShortDegree   sql.NullString `db:"short_degree"`
	EnglishDegree sql.NullString `db:"english_degree"`
	UpdatedBy     string         `db:"updated_by"`
}
