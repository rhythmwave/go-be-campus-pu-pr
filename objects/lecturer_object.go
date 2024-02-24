package objects

import (
	"time"

	"github.com/sccicitb/pupr-backend/objects/common"
)

type GetLecturerRequest struct {
	StudyProgramId             string
	IdNationalLecturer         string
	EmploymentStatus           string
	AcademicGuidanceSemesterId string
	Status                     string
	HasAuthentication          *bool
	ClassId                    string
	ExcludeLectureDate         time.Time
	ExcludeStartTime           uint32
	ExcludeEndTime             uint32
	ForceIncludeLectureId      string
}

type GetLecturer struct {
	Id                              string
	Name                            string
	PhoneNumber                     *string
	MobilePhoneNumber               *string
	OfficePhoneNumber               *string
	IdNationalLecturer              string
	FrontTitle                      *string
	BackDegree                      *string
	DiktiStudyProgramCode           *string
	StudyProgramName                *string
	EmploymentStatus                *string
	Status                          *string
	AuthenticationId                *string
	AuthenticationIsActive          *bool
	AuthenticationSuspensionRemarks *string
	AcademicGuidanceId              *string
	AcademicGuidanceTotalStudent    *uint32
	AcademicGuidanceDecisionNumber  *string
	AcademicGuidanceDecisionDate    *time.Time
}

type LecturerListWithPagination struct {
	Pagination common.Pagination
	Data       []GetLecturer
}

type GetLecturerSchedule struct {
	Id                  string
	IdNationalLecturer  string
	Name                string
	FrontTitle          *string
	BackDegree          *string
	StudyProgramName    string
	SubjectName         string
	ClassName           string
	TotalScheduleCredit uint32
	TotalSubjectCredit  uint32
	LecturePlanDate     time.Time
	StartTime           uint32
	EndTime             uint32
	RoomName            *string
	TotalParticipant    uint32
}

type LecturerScheduleWithPagination struct {
	Pagination common.Pagination
	Data       []GetLecturerSchedule
}

type GetLecturerDetail struct {
	Id                        string
	IdNationalLecturer        string
	Name                      string
	FrontTitle                *string
	BackDegree                *string
	StudyProgramId            *string
	StudyProgramName          *string
	IdNumber                  *string
	BirthDate                 *time.Time
	BirthRegencyId            *uint32
	BirthRegencyName          *string
	BirthCountryId            *uint32
	BirthCountryName          *string
	IdEmployee                *string
	Stambuk                   *string
	Sex                       *string
	BloodType                 *string
	Religion                  *string
	MaritalStatus             *string
	Address                   *string
	RegencyId                 *uint32
	RegencyName               *string
	CountryId                 *uint32
	CountryName               *string
	PostalCode                *string
	PhoneNumber               *string
	Fax                       *string
	MobilePhoneNumber         *string
	OfficePhoneNumber         *string
	EmployeeType              *string
	EmployeeStatus            *string
	SkCpnsNumber              *string
	SkCpnsDate                *time.Time
	TmtCpnsDate               *time.Time
	CpnsCategory              *string
	CpnsDurationMonth         *uint32
	PrePositionDate           *time.Time
	SkPnsNumber               *string
	SkPnsDate                 *time.Time
	TmtPnsDate                *time.Time
	PnsCategory               *string
	PnsOathDate               *time.Time
	JoinDate                  *time.Time
	EndDate                   *time.Time
	TaspenNumber              *string
	FormerInstance            *string
	Remarks                   *string
	LecturerNumber            *string
	AcademicPosition          *string
	EmploymentStatus          *string
	Expertise                 *string
	HighestDegree             *string
	InstanceCode              *string
	TeachingCertificateNumber *string
	TeachingPermitNumber      *string
	Status                    *string
	ResignSemester            *string
	ExpertiseGroupId          *string
	ExpertiseGroupName        *string
}

type GetLecturerProfile struct {
	Id                 string
	IdNationalLecturer string
	Name               string
	FrontTitle         *string
	BackDegree         *string
	StudyProgramId     *string
	StudyProgramName   *string
	BirthDate          *time.Time
	BirthRegencyId     *uint32
	BirthRegencyName   *string
	BirthCountryId     *uint32
	BirthCountryName   *string
	Sex                *string
	Religion           *string
	Address            *string
	RegencyId          *uint32
	RegencyName        *string
	CountryId          *uint32
	CountryName        *string
	PostalCode         *string
	PhoneNumber        *string
	Fax                *string
	MobilePhoneNumber  *string
	OfficePhoneNumber  *string
	AcademicPosition   *string
	EmploymentStatus   *string
	Status             *string
}

type CreateLecturer struct {
	IdNationalLecturer        string
	Name                      string
	FrontTitle                string
	BackDegree                string
	StudyProgramId            string
	IdNumber                  string
	BirthDate                 time.Time
	BirthRegencyId            uint32
	IdEmployee                string
	Stambuk                   string
	Sex                       string
	BloodType                 string
	Religion                  string
	MaritalStatus             string
	Address                   string
	RegencyId                 uint32
	PostalCode                string
	PhoneNumber               string
	Fax                       string
	MobilePhoneNumber         string
	OfficePhoneNumber         string
	EmployeeType              string
	EmployeeStatus            string
	SkCpnsNumber              string
	SkCpnsDate                time.Time
	TmtCpnsDate               time.Time
	CpnsCategory              string
	CpnsDurationMonth         uint32
	PrePositionDate           time.Time
	SkPnsNumber               string
	SkPnsDate                 time.Time
	TmtPnsDate                time.Time
	PnsCategory               string
	PnsOathDate               time.Time
	JoinDate                  time.Time
	EndDate                   time.Time
	TaspenNumber              string
	FormerInstance            string
	Remarks                   string
	LecturerNumber            string
	AcademicPosition          string
	EmploymentStatus          string
	Expertise                 string
	HighestDegree             string
	InstanceCode              string
	TeachingCertificateNumber string
	TeachingPermitNumber      string
	Status                    string
	ResignSemester            string
	ExpertiseGroupId          string
	ExpertiseGroupName        string
}

type UpdateLecturer struct {
	Id                        string
	IdNationalLecturer        string
	Name                      string
	FrontTitle                string
	BackDegree                string
	StudyProgramId            string
	IdNumber                  string
	BirthDate                 time.Time
	BirthRegencyId            uint32
	IdEmployee                string
	Stambuk                   string
	Sex                       string
	BloodType                 string
	Religion                  string
	MaritalStatus             string
	Address                   string
	RegencyId                 uint32
	PostalCode                string
	PhoneNumber               string
	Fax                       string
	MobilePhoneNumber         string
	OfficePhoneNumber         string
	EmployeeType              string
	EmployeeStatus            string
	SkCpnsNumber              string
	SkCpnsDate                time.Time
	TmtCpnsDate               time.Time
	CpnsCategory              string
	CpnsDurationMonth         uint32
	PrePositionDate           time.Time
	SkPnsNumber               string
	SkPnsDate                 time.Time
	TmtPnsDate                time.Time
	PnsCategory               string
	PnsOathDate               time.Time
	JoinDate                  time.Time
	EndDate                   time.Time
	TaspenNumber              string
	FormerInstance            string
	Remarks                   string
	LecturerNumber            string
	AcademicPosition          string
	EmploymentStatus          string
	Expertise                 string
	HighestDegree             string
	InstanceCode              string
	TeachingCertificateNumber string
	TeachingPermitNumber      string
	Status                    string
	ResignSemester            string
	ExpertiseGroupId          string
}

type GetLecturerSemesterSummary struct {
	SemesterId                   string
	StudyPlanApprovalStartDate   time.Time
	StudyPlanApprovalEndDate     time.Time
	AcademicGuidanceTotalStudent uint32
	TotalClass                   uint32
	SchoolYear                   string
	SemesterType                 string
	GradingStartDate             *time.Time
	GradingEndDate               *time.Time
}

type GetLecturerAssignedClass struct {
	Id                   string
	Name                 string
	SubjectCode          string
	SubjectName          string
	TheoryCredit         uint32
	PracticumCredit      uint32
	FieldPracticumCredit uint32
	IsGradingResponsible bool
	StudyProgramId       string
	StudyProgramName     string
	TotalAttendance      uint32
	TotalLectureDone     uint32
	AttendancePercentage float64
}
