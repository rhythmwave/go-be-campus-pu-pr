package objects

import (
	"time"

	"github.com/sccicitb/pupr-backend/objects/common"
)

type GetExamSupervisor struct {
	Id                 string
	Name               string
	IdNationalLecturer string
}

type ExamSupervisorListWithPagination struct {
	Pagination common.Pagination
	Data       []GetExamSupervisor
}

type GetExamSupervisorDetail struct {
	Id                 string
	IdNationalLecturer string
	Name               string
	FrontTitle         *string
	BackDegree         *string
	StudyProgramId     *string
	StudyProgramName   *string
	IdNumber           *string
	BirthDate          *time.Time
	BirthRegencyId     *uint32
	BirthRegencyName   *string
	BirthCountryId     *uint32
	BirthCountryName   *string
	Sex                *string
	BloodType          *string
	Religion           *string
	MaritalStatus      *string
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
	EmployeeType       *string
	EmployeeStatus     *string
	SkCpnsNumber       *string
	SkCpnsDate         *time.Time
	TmtCpnsDate        *time.Time
	CpnsCategory       *string
	CpnsDurationMonth  *uint32
	PrePositionDate    *time.Time
	SkPnsNumber        *string
	SkPnsDate          *time.Time
	TmtPnsDate         *time.Time
	PnsCategory        *string
	PnsOathDate        *time.Time
	JoinDate           *time.Time
	EndDate            *time.Time
	TaspenNumber       *string
	FormerInstance     *string
	Remarks            *string
}

type GetExamSupervisorProfile struct {
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
}

type CreateExamSupervisor struct {
	IdNationalLecturer string
	Name               string
	FrontTitle         string
	BackDegree         string
	StudyProgramId     string
	IdNumber           string
	BirthDate          time.Time
	BirthRegencyId     uint32
	Sex                string
	BloodType          string
	Religion           string
	MaritalStatus      string
	Address            string
	RegencyId          uint32
	PostalCode         string
	PhoneNumber        string
	Fax                string
	MobilePhoneNumber  string
	OfficePhoneNumber  string
	EmployeeType       string
	EmployeeStatus     string
	SkCpnsNumber       string
	SkCpnsDate         time.Time
	TmtCpnsDate        time.Time
	CpnsCategory       string
	CpnsDurationMonth  uint32
	PrePositionDate    time.Time
	SkPnsNumber        string
	SkPnsDate          time.Time
	TmtPnsDate         time.Time
	PnsCategory        string
	PnsOathDate        time.Time
	JoinDate           time.Time
	EndDate            time.Time
	TaspenNumber       string
	FormerInstance     string
	Remarks            string
}

type UpdateExamSupervisor struct {
	Id                 string
	IdNationalLecturer string
	Name               string
	FrontTitle         string
	BackDegree         string
	StudyProgramId     string
	IdNumber           string
	BirthDate          time.Time
	BirthRegencyId     uint32
	Sex                string
	BloodType          string
	Religion           string
	MaritalStatus      string
	Address            string
	RegencyId          uint32
	PostalCode         string
	PhoneNumber        string
	Fax                string
	MobilePhoneNumber  string
	OfficePhoneNumber  string
	EmployeeType       string
	EmployeeStatus     string
	SkCpnsNumber       string
	SkCpnsDate         time.Time
	TmtCpnsDate        time.Time
	CpnsCategory       string
	CpnsDurationMonth  uint32
	PrePositionDate    time.Time
	SkPnsNumber        string
	SkPnsDate          time.Time
	TmtPnsDate         time.Time
	PnsCategory        string
	PnsOathDate        time.Time
	JoinDate           time.Time
	EndDate            time.Time
	TaspenNumber       string
	FormerInstance     string
	Remarks            string
}
