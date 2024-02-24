package models

import (
	"database/sql"
	"time"
)

type GetLecturerList struct {
	Id                              string     `db:"id"`
	Name                            string     `db:"name"`
	PhoneNumber                     *string    `db:"phone_number"`
	MobilePhoneNumber               *string    `db:"mobile_phone_number"`
	OfficePhoneNumber               *string    `db:"office_phone_number"`
	IdNationalLecturer              string     `db:"id_national_lecturer"`
	FrontTitle                      *string    `db:"front_title"`
	BackDegree                      *string    `db:"back_degree"`
	DiktiStudyProgramCode           *string    `db:"dikti_study_program_code"`
	StudyProgramName                *string    `db:"study_program_name"`
	EmploymentStatus                *string    `db:"employment_status"`
	Status                          *string    `db:"status"`
	AuthenticationId                *string    `db:"authentication_id"`
	AuthenticationIsActive          *bool      `db:"authentication_is_active"`
	AuthenticationSuspensionRemarks *string    `db:"authentication_suspension_remarks"`
	AcademicGuidanceId              *string    `db:"academic_guidance_id"`
	AcademicGuidanceTotalStudent    *uint32    `db:"academic_guidance_total_student"`
	AcademicGuidanceDecisionNumber  *string    `db:"academic_guidance_decision_number"`
	AcademicGuidanceDecisionDate    *time.Time `db:"academic_guidance_decision_date"`
	TotalSupervisedThesis           uint32     `db:"total_supervised_thesis"`
}

type GetLecturerSchedule struct {
	Id                 string    `db:"id"`
	IdNationalLecturer string    `db:"id_national_lecturer"`
	Name               string    `db:"name"`
	FrontTitle         *string   `db:"front_title"`
	BackDegree         *string   `db:"back_degree"`
	StudyProgramName   string    `db:"study_program_name"`
	SubjectName        string    `db:"subject_name"`
	ClassName          string    `db:"class_name"`
	TotalSubjectCredit uint32    `db:"total_subject_credit"`
	LecturePlanDate    time.Time `db:"lecture_plan_date"`
	StartTime          uint32    `db:"start_time"`
	EndTime            uint32    `db:"end_time"`
	RoomName           *string   `db:"room_name"`
	TotalParticipant   uint32    `db:"total_participant"`
}

type GetLecturerDetail struct {
	Id                        string     `db:"id"`
	IdNationalLecturer        string     `db:"id_national_lecturer"`
	Name                      string     `db:"name"`
	FrontTitle                *string    `db:"front_title"`
	BackDegree                *string    `db:"back_degree"`
	StudyProgramId            *string    `db:"study_program_id"`
	StudyProgramName          *string    `db:"study_program_name"`
	IdNumber                  *string    `db:"id_number"`
	BirthDate                 *time.Time `db:"birth_date"`
	BirthRegencyId            *uint32    `db:"birth_regency_id"`
	BirthRegencyName          *string    `db:"birth_regency_name"`
	BirthCountryId            *uint32    `db:"birth_country_id"`
	BirthCountryName          *string    `db:"birth_country_name"`
	IdEmployee                *string    `db:"id_employee"`
	Stambuk                   *string    `db:"stambuk"`
	Sex                       *string    `db:"sex"`
	BloodType                 *string    `db:"blood_type"`
	Religion                  *string    `db:"religion"`
	MaritalStatus             *string    `db:"marital_status"`
	Address                   *string    `db:"address"`
	RegencyId                 *uint32    `db:"regency_id"`
	RegencyName               *string    `db:"regency_name"`
	CountryId                 *uint32    `db:"country_id"`
	CountryName               *string    `db:"country_name"`
	PostalCode                *string    `db:"postal_code"`
	PhoneNumber               *string    `db:"phone_number"`
	Fax                       *string    `db:"fax"`
	MobilePhoneNumber         *string    `db:"mobile_phone_number"`
	OfficePhoneNumber         *string    `db:"office_phone_number"`
	EmployeeType              *string    `db:"employee_type"`
	EmployeeStatus            *string    `db:"employee_status"`
	SkCpnsNumber              *string    `db:"sk_cpns_number"`
	SkCpnsDate                *time.Time `db:"sk_cpns_date"`
	TmtCpnsDate               *time.Time `db:"tmt_cpns_date"`
	CpnsCategory              *string    `db:"cpns_category"`
	CpnsDurationMonth         *uint32    `db:"cpns_duration_month"`
	PrePositionDate           *time.Time `db:"pre_position_date"`
	SkPnsNumber               *string    `db:"sk_pns_number"`
	SkPnsDate                 *time.Time `db:"sk_pns_date"`
	TmtPnsDate                *time.Time `db:"tmt_pns_date"`
	PnsCategory               *string    `db:"pns_category"`
	PnsOathDate               *time.Time `db:"pns_oath_date"`
	JoinDate                  *time.Time `db:"join_date"`
	EndDate                   *time.Time `db:"end_date"`
	TaspenNumber              *string    `db:"taspen_number"`
	FormerInstance            *string    `db:"former_instance"`
	Remarks                   *string    `db:"remarks"`
	LecturerNumber            *string    `db:"lecturer_number"`
	AcademicPosition          *string    `db:"academic_position"`
	EmploymentStatus          *string    `db:"employment_status"`
	Expertise                 *string    `db:"expertise"`
	HighestDegree             *string    `db:"highest_degree"`
	InstanceCode              *string    `db:"instance_code"`
	TeachingCertificateNumber *string    `db:"teaching_certificate_number"`
	TeachingPermitNumber      *string    `db:"teaching_permit_number"`
	Status                    *string    `db:"status"`
	ResignSemester            *string    `db:"resign_semester"`
	ExpertiseGroupId          *string    `db:"expertise_group_id"`
	ExpertiseGroupName        *string    `db:"expertise_group_name"`
}

type CreateLecturer struct {
	IdNationalLecturer        string         `db:"id_national_lecturer"`
	Name                      string         `db:"name"`
	FrontTitle                sql.NullString `db:"front_title"`
	BackDegree                sql.NullString `db:"back_degree"`
	StudyProgramId            sql.NullString `db:"study_program_id"`
	IdNumber                  sql.NullString `db:"id_number"`
	BirthDate                 sql.NullTime   `db:"birth_date"`
	BirthRegencyId            sql.NullInt32  `db:"birth_regency_id"`
	IdEmployee                sql.NullString `db:"id_employee"`
	Stambuk                   sql.NullString `db:"stambuk"`
	Sex                       sql.NullString `db:"sex"`
	BloodType                 sql.NullString `db:"blood_type"`
	Religion                  sql.NullString `db:"religion"`
	MaritalStatus             sql.NullString `db:"marital_status"`
	Address                   sql.NullString `db:"address"`
	RegencyId                 sql.NullInt32  `db:"regency_id"`
	PostalCode                sql.NullString `db:"postal_code"`
	PhoneNumber               sql.NullString `db:"phone_number"`
	Fax                       sql.NullString `db:"fax"`
	MobilePhoneNumber         sql.NullString `db:"mobile_phone_number"`
	OfficePhoneNumber         sql.NullString `db:"office_phone_number"`
	EmployeeType              sql.NullString `db:"employee_type"`
	EmployeeStatus            sql.NullString `db:"employee_status"`
	SkCpnsNumber              sql.NullString `db:"sk_cpns_number"`
	SkCpnsDate                sql.NullTime   `db:"sk_cpns_date"`
	TmtCpnsDate               sql.NullTime   `db:"tmt_cpns_date"`
	CpnsCategory              sql.NullString `db:"cpns_category"`
	CpnsDurationMonth         sql.NullInt32  `db:"cpns_duration_month"`
	PrePositionDate           sql.NullTime   `db:"pre_position_date"`
	SkPnsNumber               sql.NullString `db:"sk_pns_number"`
	SkPnsDate                 sql.NullTime   `db:"sk_pns_date"`
	TmtPnsDate                sql.NullTime   `db:"tmt_pns_date"`
	PnsCategory               sql.NullString `db:"pns_category"`
	PnsOathDate               sql.NullTime   `db:"pns_oath_date"`
	JoinDate                  sql.NullTime   `db:"join_date"`
	EndDate                   sql.NullTime   `db:"end_date"`
	TaspenNumber              sql.NullString `db:"taspen_number"`
	FormerInstance            sql.NullString `db:"former_instance"`
	Remarks                   sql.NullString `db:"remarks"`
	LecturerNumber            sql.NullString `db:"lecturer_number"`
	AcademicPosition          sql.NullString `db:"academic_position"`
	EmploymentStatus          sql.NullString `db:"employment_status"`
	Expertise                 sql.NullString `db:"expertise"`
	HighestDegree             sql.NullString `db:"highest_degree"`
	InstanceCode              sql.NullString `db:"instance_code"`
	TeachingCertificateNumber sql.NullString `db:"teaching_certificate_number"`
	TeachingPermitNumber      sql.NullString `db:"teaching_permit_number"`
	Status                    sql.NullString `db:"status"`
	ResignSemester            sql.NullString `db:"resign_semester"`
	ExpertiseGroupId          sql.NullString `db:"expertise_group_id"`
	CreatedBy                 string         `db:"created_by"`
}

type UpdateLecturer struct {
	Id                        string         `db:"id"`
	IdNationalLecturer        string         `db:"id_national_lecturer"`
	Name                      string         `db:"name"`
	FrontTitle                sql.NullString `db:"front_title"`
	BackDegree                sql.NullString `db:"back_degree"`
	StudyProgramId            sql.NullString `db:"study_program_id"`
	IdNumber                  sql.NullString `db:"id_number"`
	BirthDate                 sql.NullTime   `db:"birth_date"`
	BirthRegencyId            sql.NullInt32  `db:"birth_regency_id"`
	IdEmployee                sql.NullString `db:"id_employee"`
	Stambuk                   sql.NullString `db:"stambuk"`
	Sex                       sql.NullString `db:"sex"`
	BloodType                 sql.NullString `db:"blood_type"`
	Religion                  sql.NullString `db:"religion"`
	MaritalStatus             sql.NullString `db:"marital_status"`
	Address                   sql.NullString `db:"address"`
	RegencyId                 sql.NullInt32  `db:"regency_id"`
	PostalCode                sql.NullString `db:"postal_code"`
	PhoneNumber               sql.NullString `db:"phone_number"`
	Fax                       sql.NullString `db:"fax"`
	MobilePhoneNumber         sql.NullString `db:"mobile_phone_number"`
	OfficePhoneNumber         sql.NullString `db:"office_phone_number"`
	EmployeeType              sql.NullString `db:"employee_type"`
	EmployeeStatus            sql.NullString `db:"employee_status"`
	SkCpnsNumber              sql.NullString `db:"sk_cpns_number"`
	SkCpnsDate                sql.NullTime   `db:"sk_cpns_date"`
	TmtCpnsDate               sql.NullTime   `db:"tmt_cpns_date"`
	CpnsCategory              sql.NullString `db:"cpns_category"`
	CpnsDurationMonth         sql.NullInt32  `db:"cpns_duration_month"`
	PrePositionDate           sql.NullTime   `db:"pre_position_date"`
	SkPnsNumber               sql.NullString `db:"sk_pns_number"`
	SkPnsDate                 sql.NullTime   `db:"sk_pns_date"`
	TmtPnsDate                sql.NullTime   `db:"tmt_pns_date"`
	PnsCategory               sql.NullString `db:"pns_category"`
	PnsOathDate               sql.NullTime   `db:"pns_oath_date"`
	JoinDate                  sql.NullTime   `db:"join_date"`
	EndDate                   sql.NullTime   `db:"end_date"`
	TaspenNumber              sql.NullString `db:"taspen_number"`
	FormerInstance            sql.NullString `db:"former_instance"`
	Remarks                   sql.NullString `db:"remarks"`
	LecturerNumber            sql.NullString `db:"lecturer_number"`
	AcademicPosition          sql.NullString `db:"academic_position"`
	EmploymentStatus          sql.NullString `db:"employment_status"`
	Expertise                 sql.NullString `db:"expertise"`
	HighestDegree             sql.NullString `db:"highest_degree"`
	InstanceCode              sql.NullString `db:"instance_code"`
	TeachingCertificateNumber sql.NullString `db:"teaching_certificate_number"`
	TeachingPermitNumber      sql.NullString `db:"teaching_permit_number"`
	Status                    sql.NullString `db:"status"`
	ResignSemester            sql.NullString `db:"resign_semester"`
	ExpertiseGroupId          sql.NullString `db:"expertise_group_id"`
	UpdatedBy                 string         `db:"updated_by"`
}

type LecturerId struct {
	LecturerId string `db:"lecturer_id"`
}

type GetLecturerAssignedClass struct {
	Id                   string `db:"id"`
	Name                 string `db:"name"`
	SubjectCode          string `db:"subject_code"`
	SubjectName          string `db:"subject_name"`
	TheoryCredit         uint32 `db:"theory_credit"`
	PracticumCredit      uint32 `db:"practicum_credit"`
	FieldPracticumCredit uint32 `db:"field_practicum_credit"`
	IsGradingResponsible bool   `db:"is_grading_responsible"`
	StudyProgramId       string `db:"study_program_id"`
	StudyProgramName     string `db:"study_program_name"`
	TotalAttendance      uint32 `db:"total_attendance"`
	TotalLecturePlan     uint32 `db:"total_lecture_plan"`
	TotalLectureDone     uint32 `db:"total_lecture_done"`
}
