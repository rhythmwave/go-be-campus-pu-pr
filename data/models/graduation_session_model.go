package models

import "time"

type GetGraduationSession struct {
	Id            string    `db:"id"`
	SessionYear   uint32    `db:"session_year"`
	SessionNumber uint32    `db:"session_number"`
	SessionDate   time.Time `db:"session_date"`
}

type CreateGraduationSession struct {
	SessionYear   uint32    `db:"session_year"`
	SessionNumber uint32    `db:"session_number"`
	SessionDate   time.Time `db:"session_date"`
	CreatedBy     string    `db:"created_by"`
}

type UpdateGraduationSession struct {
	Id            string    `db:"id"`
	SessionYear   uint32    `db:"session_year"`
	SessionNumber uint32    `db:"session_number"`
	SessionDate   time.Time `db:"session_date"`
	UpdatedBy     string    `db:"updated_by"`
}
