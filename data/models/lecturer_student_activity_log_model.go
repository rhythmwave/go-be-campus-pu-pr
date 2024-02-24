package models

import (
	"database/sql"
	"time"
)

type GetLecturerStudentActivityLog struct {
	Id               string    `db:"id"`
	StudentId        *string   `db:"student_id"`
	StudentName      *string   `db:"student_name"`
	StudentUsername  *string   `db:"student_username"`
	LecturerId       *string   `db:"lecturer_id"`
	LecturerName     *string   `db:"lecturer_name"`
	LecturerUsername *string   `db:"lecturer_username"`
	Module           string    `db:"module"`
	Action           string    `db:"action"`
	IpAddress        string    `db:"ip_address"`
	UserAgent        string    `db:"user_agent"`
	ExecutionTime    float64   `db:"execution_time"`
	MemoryUsage      float64   `db:"memory_usage"`
	CreatedAt        time.Time `db:"created_at"`
}

type CreateLecturerStudentActivityLog struct {
	LecturerId       sql.NullString `db:"lecturer_id"`
	LecturerName     sql.NullString `db:"lecturer_name"`
	LecturerUsername sql.NullString `db:"lecturer_username"`
	StudentId        sql.NullString `db:"student_id"`
	StudentName      sql.NullString `db:"student_name"`
	StudentUsername  sql.NullString `db:"student_username"`
	Module           string         `db:"module"`
	Action           string         `db:"action"`
	IpAddress        string         `db:"ip_address"`
	UserAgent        string         `db:"user_agent"`
	ExecutionTime    float64        `db:"execution_time"`
	MemoryUsage      float64        `db:"memory_usage"`
}
