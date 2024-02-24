package models

import "database/sql"

type GetMajorList struct {
	Id          string `db:"id"`
	FacultyName string `db:"faculty_name"`
	Name        string `db:"name"`
}

type GetMajorDetail struct {
	Id                        string   `db:"id"`
	FacultyId                 string   `db:"faculty_id"`
	FacultyName               string   `db:"faculty_name"`
	Name                      string   `db:"name"`
	ShortName                 *string  `db:"short_name"`
	EnglishName               *string  `db:"english_name"`
	EnglishShortName          *string  `db:"english_short_name"`
	Address                   string   `db:"address"`
	PhoneNumber               *string  `db:"phone_number"`
	Fax                       *string  `db:"fax"`
	Email                     *string  `db:"email"`
	ContactPerson             *string  `db:"contact_person"`
	ExperimentBuildingArea    *float64 `db:"experiment_building_area"`
	LectureHallArea           *float64 `db:"lecture_hall_area"`
	LectureHallCount          *uint32  `db:"lecture_hall_count"`
	LaboratoriumArea          *float64 `db:"laboratorium_area"`
	LaboratoriumCount         *uint32  `db:"laboratorium_count"`
	PermanentLecturerRoomArea *float64 `db:"permanent_lecturer_room_area"`
	AdministrationRoomArea    *float64 `db:"administration_room_area"`
	BookCount                 *uint32  `db:"book_count"`
	BookCopyCount             *uint32  `db:"book_copy_count"`
}

type CreateMajor struct {
	FacultyId                 string          `db:"faculty_id"`
	Name                      string          `db:"name"`
	ShortName                 sql.NullString  `db:"short_name"`
	EnglishName               sql.NullString  `db:"english_name"`
	EnglishShortName          sql.NullString  `db:"english_short_name"`
	Address                   string          `db:"address"`
	PhoneNumber               sql.NullString  `db:"phone_number"`
	Fax                       sql.NullString  `db:"fax"`
	Email                     sql.NullString  `db:"email"`
	ContactPerson             sql.NullString  `db:"contact_person"`
	ExperimentBuildingArea    sql.NullFloat64 `db:"experiment_building_area"`
	LectureHallArea           sql.NullFloat64 `db:"lecture_hall_area"`
	LectureHallCount          sql.NullInt32   `db:"lecture_hall_count"`
	LaboratoriumArea          sql.NullFloat64 `db:"laboratorium_area"`
	LaboratoriumCount         sql.NullInt32   `db:"laboratorium_count"`
	PermanentLecturerRoomArea sql.NullFloat64 `db:"permanent_lecturer_room_area"`
	AdministrationRoomArea    sql.NullFloat64 `db:"administration_room_area"`
	BookCount                 sql.NullInt32   `db:"book_count"`
	BookCopyCount             sql.NullInt32   `db:"book_copy_count"`
	CreatedBy                 string          `db:"created_by"`
}

type UpdateMajor struct {
	Id                        string          `db:"id"`
	FacultyId                 string          `db:"faculty_id"`
	Name                      string          `db:"name"`
	ShortName                 sql.NullString  `db:"short_name"`
	EnglishName               sql.NullString  `db:"english_name"`
	EnglishShortName          sql.NullString  `db:"english_short_name"`
	Address                   string          `db:"address"`
	PhoneNumber               sql.NullString  `db:"phone_number"`
	Fax                       sql.NullString  `db:"fax"`
	Email                     sql.NullString  `db:"email"`
	ContactPerson             sql.NullString  `db:"contact_person"`
	ExperimentBuildingArea    sql.NullFloat64 `db:"experiment_building_area"`
	LectureHallArea           sql.NullFloat64 `db:"lecture_hall_area"`
	LectureHallCount          sql.NullInt32   `db:"lecture_hall_count"`
	LaboratoriumArea          sql.NullFloat64 `db:"laboratorium_area"`
	LaboratoriumCount         sql.NullInt32   `db:"laboratorium_count"`
	PermanentLecturerRoomArea sql.NullFloat64 `db:"permanent_lecturer_room_area"`
	AdministrationRoomArea    sql.NullFloat64 `db:"administration_room_area"`
	BookCount                 sql.NullInt32   `db:"book_count"`
	BookCopyCount             sql.NullInt32   `db:"book_copy_count"`
	UpdatedBy                 string          `db:"updated_by"`
}
