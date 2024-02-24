package models

import (
	"database/sql"
	"time"
)

type GetStudentSkpi struct {
	Id                           string `db:"id"`
	StudentId                    string `db:"student_id"`
	StudentNimNumber             int64  `db:"student_nim_number"`
	StudentName                  string `db:"student_name"`
	StudentStudyProgramId        string `db:"student_study_program_id"`
	StudentStudyProgramName      string `db:"student_study_program_name"`
	StudentDiktiStudyProgramCode string `db:"student_dikti_study_program_code"`
	IsApproved                   bool   `db:"is_approved"`
}

type GetStudentSkpiAchievement struct {
	Id   string `db:"id"`
	Name string `db:"name"`
	Year uint32 `db:"year"`
}

type GetStudentSkpiOrganization struct {
	Id            string `db:"id"`
	Name          string `db:"name"`
	Position      string `db:"position"`
	ServiceLength uint32 `db:"service_length"`
}

type GetStudentSkpiCertificate struct {
	Id   string `db:"id"`
	Name string `db:"name"`
}

type GetStudentSkpiCharacterBuilding struct {
	Id   string `db:"id"`
	Name string `db:"name"`
}

type GetStudentSkpiInternship struct {
	Id   string `db:"id"`
	Name string `db:"name"`
}

type GetStudentSkpiLanguage struct {
	Id    string    `db:"id"`
	Name  string    `db:"name"`
	Score string    `db:"score"`
	Date  time.Time `db:"date"`
}

type GetStudentSkpiDetail struct {
	Id                           string  `db:"id"`
	StudentId                    string  `db:"student_id"`
	StudentNimNumber             int64   `db:"student_nim_number"`
	StudentName                  string  `db:"student_name"`
	StudentStudyProgramId        string  `db:"student_study_program_id"`
	StudentStudyProgramName      string  `db:"student_study_program_name"`
	StudentDiktiStudyProgramCode string  `db:"student_dikti_study_program_code"`
	SkpiNumber                   *string `db:"skpi_number"`
	IsApproved                   bool    `db:"is_approved"`
	AchievementPath              *string `db:"achievement_path"`
	AchievementPathType          *string `db:"achievement_path_type"`
	OrganizationPath             *string `db:"organization_path"`
	OrganizationPathType         *string `db:"organization_path_type"`
	CertificatePath              *string `db:"certificate_path"`
	CertificatePathType          *string `db:"certificate_path_type"`
	LanguagePath                 *string `db:"language_path"`
	LanguagePathType             *string `db:"language_path_type"`
}

type UpsertStudentSkpiAchievement struct {
	StudentSkpiId string `db:"student_skpi_id"`
	Name          string `db:"name"`
	Year          uint32 `db:"year"`
}

type UpsertStudentSkpiOrganization struct {
	StudentSkpiId string `db:"student_skpi_id"`
	Name          string `db:"name"`
	Position      string `db:"position"`
	ServiceLength uint32 `db:"service_length"`
}

type UpsertStudentSkpiCertificate struct {
	StudentSkpiId string `db:"student_skpi_id"`
	Name          string `db:"name"`
}

type UpsertStudentSkpiCharacterBuilding struct {
	StudentSkpiId string `db:"student_skpi_id"`
	Name          string `db:"name"`
}

type UpsertStudentSkpiInternship struct {
	StudentSkpiId string `db:"student_skpi_id"`
	Name          string `db:"name"`
}

type UpsertStudentSkpiLanguage struct {
	StudentSkpiId string    `db:"student_skpi_id"`
	Name          string    `db:"name"`
	Score         string    `db:"score"`
	Date          time.Time `db:"date"`
}

type UpsertStudentSkpi struct {
	StudentId            string         `db:"student_id"`
	AchievementPath      sql.NullString `db:"achievement_path"`
	AchievementPathType  sql.NullString `db:"achievement_path_type"`
	OrganizationPath     sql.NullString `db:"organization_path"`
	OrganizationPathType sql.NullString `db:"organization_path_type"`
	CertificatePath      sql.NullString `db:"certificate_path"`
	CertificatePathType  sql.NullString `db:"certificate_path_type"`
	LanguagePath         sql.NullString `db:"language_path"`
	LanguagePathType     sql.NullString `db:"language_path_type"`
}

type ApproveStudentSkpi struct {
	Id         string `db:"id"`
	SkpiNumber string `db:"skpi_number"`
}
