package models

import (
	"database/sql"
	"time"
)

type GetGradeType struct {
	Id                  string    `db:"id"`
	StudyLevelId        string    `db:"study_level_id"`
	StudyLevelShortName string    `db:"study_level_short_name"`
	Code                string    `db:"code"`
	GradePoint          float64   `db:"grade_point"`
	MinimumGrade        float64   `db:"minimum_grade"`
	MaximumGrade        float64   `db:"maximum_grade"`
	GradeCategory       string    `db:"grade_category"`
	GradePointCategory  float64   `db:"grade_point_category"`
	Label               *string   `db:"label"`
	EnglishLabel        *string   `db:"english_label"`
	StartDate           time.Time `db:"start_date"`
	EndDate             time.Time `db:"end_date"`
}

type CreateGradeType struct {
	StudyLevelId       string         `db:"study_level_id"`
	Code               string         `db:"code"`
	GradePoint         float64        `db:"grade_point"`
	MinimumGrade       float64        `db:"minimum_grade"`
	MaximumGrade       float64        `db:"maximum_grade"`
	GradeCategory      string         `db:"grade_category"`
	GradePointCategory float64        `db:"grade_point_category"`
	Label              sql.NullString `db:"label"`
	EnglishLabel       sql.NullString `db:"english_label"`
	StartDate          time.Time      `db:"start_date"`
	EndDate            time.Time      `db:"end_date"`
	CreatedBy          string         `db:"created_by"`
}

type UpdateGradeType struct {
	Id                 string         `db:"id"`
	Code               string         `db:"code"`
	GradePoint         float64        `db:"grade_point"`
	MinimumGrade       float64        `db:"minimum_grade"`
	MaximumGrade       float64        `db:"maximum_grade"`
	GradeCategory      string         `db:"grade_category"`
	GradePointCategory float64        `db:"grade_point_category"`
	Label              sql.NullString `db:"label"`
	EnglishLabel       sql.NullString `db:"english_label"`
	StartDate          time.Time      `db:"start_date"`
	EndDate            time.Time      `db:"end_date"`
	UpdatedBy          string         `db:"updated_by"`
}
