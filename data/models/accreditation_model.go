package models

import (
	"time"
)

type GetAccreditation struct {
	Id             string    `db:"id"`
	StudyProgramId string    `db:"study_program_id"`
	DecreeNumber   string    `db:"decree_number"`
	DecreeDate     time.Time `db:"decree_date"`
	DecreeDueDate  time.Time `db:"decree_due_date"`
	Accreditation  string    `db:"accreditation"`
	IsActive       bool      `db:"is_active"`
}

type CreateAccreditation struct {
	StudyProgramId string    `db:"study_program_id"`
	DecreeNumber   string    `db:"decree_number"`
	DecreeDate     time.Time `db:"decree_date"`
	DecreeDueDate  time.Time `db:"decree_due_date"`
	Accreditation  string    `db:"accreditation"`
	IsActive       bool      `db:"is_active"`
	CreatedBy      string    `db:"created_by"`
}

type UpdateAccreditation struct {
	Id            string    `db:"id"`
	DecreeNumber  string    `db:"decree_number"`
	DecreeDate    time.Time `db:"decree_date"`
	DecreeDueDate time.Time `db:"decree_due_date"`
	Accreditation string    `db:"accreditation"`
	IsActive      bool      `db:"is_active"`
	UpdatedBy     string    `db:"updated_by"`
}
