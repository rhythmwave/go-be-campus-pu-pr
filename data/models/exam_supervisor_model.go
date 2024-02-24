package models

import (
	"database/sql"
	"time"
)

type GetExamSupervisorList struct {
	Id                 string `db:"id"`
	Name               string `db:"name"`
	IdNationalLecturer string `db:"id_national_lecturer"`
}

type GetExamSupervisorDetail struct {
	Id                 string     `db:"id"`
	IdNationalLecturer string     `db:"id_national_lecturer"`
	Name               string     `db:"name"`
	FrontTitle         *string    `db:"front_title"`
	BackDegree         *string    `db:"back_degree"`
	StudyProgramId     *string    `db:"study_program_id"`
	StudyProgramName   *string    `db:"study_program_name"`
	IdNumber           *string    `db:"id_number"`
	BirthDate          *time.Time `db:"birth_date"`
	BirthRegencyId     *uint32    `db:"birth_regency_id"`
	BirthRegencyName   *string    `db:"birth_regency_name"`
	BirthCountryId     *uint32    `db:"birth_country_id"`
	BirthCountryName   *string    `db:"birth_country_name"`
	Sex                *string    `db:"sex"`
	BloodType          *string    `db:"blood_type"`
	Religion           *string    `db:"religion"`
	MaritalStatus      *string    `db:"marital_status"`
	Address            *string    `db:"address"`
	RegencyId          *uint32    `db:"regency_id"`
	RegencyName        *string    `db:"regency_name"`
	CountryId          *uint32    `db:"country_id"`
	CountryName        *string    `db:"country_name"`
	PostalCode         *string    `db:"postal_code"`
	PhoneNumber        *string    `db:"phone_number"`
	Fax                *string    `db:"fax"`
	MobilePhoneNumber  *string    `db:"mobile_phone_number"`
	OfficePhoneNumber  *string    `db:"office_phone_number"`
	EmployeeType       *string    `db:"employee_type"`
	EmployeeStatus     *string    `db:"employee_status"`
	SkCpnsNumber       *string    `db:"sk_cpns_number"`
	SkCpnsDate         *time.Time `db:"sk_cpns_date"`
	TmtCpnsDate        *time.Time `db:"tmt_cpns_date"`
	CpnsCategory       *string    `db:"cpns_category"`
	CpnsDurationMonth  *uint32    `db:"cpns_duration_month"`
	PrePositionDate    *time.Time `db:"pre_position_date"`
	SkPnsNumber        *string    `db:"sk_pns_number"`
	SkPnsDate          *time.Time `db:"sk_pns_date"`
	TmtPnsDate         *time.Time `db:"tmt_pns_date"`
	PnsCategory        *string    `db:"pns_category"`
	PnsOathDate        *time.Time `db:"pns_oath_date"`
	JoinDate           *time.Time `db:"join_date"`
	EndDate            *time.Time `db:"end_date"`
	TaspenNumber       *string    `db:"taspen_number"`
	FormerInstance     *string    `db:"former_instance"`
	Remarks            *string    `db:"remarks"`
}

type CreateExamSupervisor struct {
	IdNationalLecturer string         `db:"id_national_lecturer"`
	Name               string         `db:"name"`
	FrontTitle         sql.NullString `db:"front_title"`
	BackDegree         sql.NullString `db:"back_degree"`
	StudyProgramId     sql.NullString `db:"study_program_id"`
	IdNumber           sql.NullString `db:"id_number"`
	BirthDate          sql.NullTime   `db:"birth_date"`
	BirthRegencyId     sql.NullInt32  `db:"birth_regency_id"`
	Sex                sql.NullString `db:"sex"`
	BloodType          sql.NullString `db:"blood_type"`
	Religion           sql.NullString `db:"religion"`
	MaritalStatus      sql.NullString `db:"marital_status"`
	Address            sql.NullString `db:"address"`
	RegencyId          sql.NullInt32  `db:"regency_id"`
	PostalCode         sql.NullString `db:"postal_code"`
	PhoneNumber        sql.NullString `db:"phone_number"`
	Fax                sql.NullString `db:"fax"`
	MobilePhoneNumber  sql.NullString `db:"mobile_phone_number"`
	OfficePhoneNumber  sql.NullString `db:"office_phone_number"`
	EmployeeType       sql.NullString `db:"employee_type"`
	EmployeeStatus     sql.NullString `db:"employee_status"`
	SkCpnsNumber       sql.NullString `db:"sk_cpns_number"`
	SkCpnsDate         sql.NullTime   `db:"sk_cpns_date"`
	TmtCpnsDate        sql.NullTime   `db:"tmt_cpns_date"`
	CpnsCategory       sql.NullString `db:"cpns_category"`
	CpnsDurationMonth  sql.NullInt32  `db:"cpns_duration_month"`
	PrePositionDate    sql.NullTime   `db:"pre_position_date"`
	SkPnsNumber        sql.NullString `db:"sk_pns_number"`
	SkPnsDate          sql.NullTime   `db:"sk_pns_date"`
	TmtPnsDate         sql.NullTime   `db:"tmt_pns_date"`
	PnsCategory        sql.NullString `db:"pns_category"`
	PnsOathDate        sql.NullTime   `db:"pns_oath_date"`
	JoinDate           sql.NullTime   `db:"join_date"`
	EndDate            sql.NullTime   `db:"end_date"`
	TaspenNumber       sql.NullString `db:"taspen_number"`
	FormerInstance     sql.NullString `db:"former_instance"`
	Remarks            sql.NullString `db:"remarks"`
	CreatedBy          string         `db:"created_by"`
}

type UpdateExamSupervisor struct {
	Id                 string         `db:"id"`
	IdNationalLecturer string         `db:"id_national_lecturer"`
	Name               string         `db:"name"`
	FrontTitle         sql.NullString `db:"front_title"`
	BackDegree         sql.NullString `db:"back_degree"`
	StudyProgramId     sql.NullString `db:"study_program_id"`
	IdNumber           sql.NullString `db:"id_number"`
	BirthDate          sql.NullTime   `db:"birth_date"`
	BirthRegencyId     sql.NullInt32  `db:"birth_regency_id"`
	Sex                sql.NullString `db:"sex"`
	BloodType          sql.NullString `db:"blood_type"`
	Religion           sql.NullString `db:"religion"`
	MaritalStatus      sql.NullString `db:"marital_status"`
	Address            sql.NullString `db:"address"`
	RegencyId          sql.NullInt32  `db:"regency_id"`
	PostalCode         sql.NullString `db:"postal_code"`
	PhoneNumber        sql.NullString `db:"phone_number"`
	Fax                sql.NullString `db:"fax"`
	MobilePhoneNumber  sql.NullString `db:"mobile_phone_number"`
	OfficePhoneNumber  sql.NullString `db:"office_phone_number"`
	EmployeeType       sql.NullString `db:"employee_type"`
	EmployeeStatus     sql.NullString `db:"employee_status"`
	SkCpnsNumber       sql.NullString `db:"sk_cpns_number"`
	SkCpnsDate         sql.NullTime   `db:"sk_cpns_date"`
	TmtCpnsDate        sql.NullTime   `db:"tmt_cpns_date"`
	CpnsCategory       sql.NullString `db:"cpns_category"`
	CpnsDurationMonth  sql.NullInt32  `db:"cpns_duration_month"`
	PrePositionDate    sql.NullTime   `db:"pre_position_date"`
	SkPnsNumber        sql.NullString `db:"sk_pns_number"`
	SkPnsDate          sql.NullTime   `db:"sk_pns_date"`
	TmtPnsDate         sql.NullTime   `db:"tmt_pns_date"`
	PnsCategory        sql.NullString `db:"pns_category"`
	PnsOathDate        sql.NullTime   `db:"pns_oath_date"`
	JoinDate           sql.NullTime   `db:"join_date"`
	EndDate            sql.NullTime   `db:"end_date"`
	TaspenNumber       sql.NullString `db:"taspen_number"`
	FormerInstance     sql.NullString `db:"former_instance"`
	Remarks            sql.NullString `db:"remarks"`
	UpdatedBy          string         `db:"updated_by"`
}

type GetExamLectureSupervisor struct {
	Id                     string  `db:"id"`
	LectureId              string  `db:"lecture_id"`
	IdNationalLecturer     string  `db:"id_national_lecturer"`
	Name                   string  `db:"name"`
	FrontTitle             *string `db:"front_title"`
	BackDegree             *string `db:"back_degree"`
	ExamSupervisorRoleId   string  `db:"exam_supervisor_role_id"`
	ExamSupervisorRoleName string  `db:"exam_supervisor_role_name"`
}

type UpsertExamLectureSupervisor struct {
	LectureId            string `db:"lecture_id"`
	ExamSupervisorId     string `db:"exam_supervisor_id"`
	ExamSupervisorRoleId string `db:"exam_supervisor_role_id"`
	CreatedBy            string `db:"created_by"`
}
