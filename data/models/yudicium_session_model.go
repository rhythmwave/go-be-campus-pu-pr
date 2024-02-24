package models

import "time"

type GetYudiciumSession struct {
	Id                string    `db:"id"`
	SemesterId        string    `db:"semester_id"`
	SemesterStartYear uint32    `db:"semester_start_year"`
	SemesterType      string    `db:"semester_type"`
	Name              string    `db:"name"`
	SessionDate       time.Time `db:"session_date"`
}

type CreateYudiciumSession struct {
	SemesterId  string    `db:"semester_id"`
	Name        string    `db:"name"`
	SessionDate time.Time `db:"session_date"`
	CreatedBy   string    `db:"created_by"`
}

type UpdateYudiciumSession struct {
	Id          string    `db:"id"`
	SemesterId  string    `db:"semester_id"`
	Name        string    `db:"name"`
	SessionDate time.Time `db:"session_date"`
	UpdatedBy   string    `db:"updated_by"`
}

type DoYudicium struct {
	Id             string    `db:"id"`
	YudiciumNumber string    `db:"yudicium_number"`
	ActualDate     time.Time `db:"actual_date"`
}
