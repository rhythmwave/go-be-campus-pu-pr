package objects

import (
	"time"

	"github.com/sccicitb/pupr-backend/objects/common"
)

type GetStudentRequest struct {
	StudyProgramId       string
	StudentForceFrom     uint32
	StudentForceTo       uint32
	NimNumberFrom        int64
	NimNumberTo          int64
	Name                 string
	Address              string
	RegencyId            string
	Status               []string
	GetAcademicGuidance  bool
	HasAuthentication    *bool
	StudyPlanSemesterId  string
	StudyPlanIsSubmitted *bool
	StudyPlanIsApproved  *bool
	HasStudyPlan         *bool
	StatusSemesterId     string
	IsRegistered         *bool
	HasPaid              *bool
	IsGraduateEligible   *bool
	IsThesisEligible     *bool
	YudiciumSessionId    string
}

type GetStudentPreHighSchoolHistory struct {
	Id             string
	Level          string
	Name           string
	GraduationYear string
}

type GetStudent struct {
	Id                                 string
	Name                               string
	Sex                                *string
	MaritalStatus                      *string
	BirthRegencyId                     *uint32
	BirthRegencyName                   *string
	BirthDate                          *time.Time
	Religion                           *string
	Address                            *string
	Rt                                 *string
	Rw                                 *string
	VillageId                          *uint32
	VillageName                        *string
	DistrictId                         *uint32
	DistrictName                       *string
	RegencyId                          *uint32
	RegencyName                        *string
	ProvinceId                         *uint32
	ProvinceName                       *string
	CountryId                          *uint32
	CountryName                        *string
	PostalCode                         *string
	PreviousAddress                    *string
	IdNumber                           *string
	NpwpNumber                         *string
	NisnNumber                         *string
	ResidenceType                      *string
	TransportationMean                 *string
	KpsReceiver                        *string
	PhoneNumber                        *string
	MobilePhoneNumber                  *string
	Email                              *string
	Homepage                           *string
	WorkType                           *string
	WorkPlace                          *string
	Nationality                        *string
	AskesNumber                        *string
	TotalBrother                       *uint32
	TotalSister                        *uint32
	Hobby                              *string
	Experience                         *string
	TotalDependent                     *uint32
	NimNumber                          int64
	DiktiStudyProgramType              *string
	StudyLevelShortName                *string
	StudentForce                       *uint32
	AdmittanceSemester                 *string
	StudyProgramId                     *string
	StudyProgramName                   *string
	CurriculumId                       *string
	CurriculumName                     *string
	AdmittanceTestNumber               *string
	AdmittanceDate                     *time.Time
	AdmittanceStatus                   *string
	TotalTransferCredit                *uint32
	PreviousCollege                    *string
	PreviousStudyProgram               *string
	PreviousNimNumber                  *int64
	PreviousNimAdmittanceYear          *string
	Status                             *string
	IsForeignStudent                   *bool
	CollegeEntranceType                *string
	ClassTime                          *string
	FundSource                         *string
	IsScholarshipGrantee               *bool
	HasCompleteRequirement             *bool
	SchoolCertificateType              *string
	SchoolGraduationYear               *string
	SchoolName                         *string
	SchoolAccreditation                *string
	SchoolAddress                      *string
	SchoolMajor                        *string
	SchoolCertificateNumber            *string
	SchoolCertificateDate              *time.Time
	TotalSchoolFinalExamSubject        *uint32
	SchoolFinalExamScore               *float64
	GuardianName                       *string
	GuardianBirthDate                  *time.Time
	GuardianDeathDate                  *time.Time
	GuardianAddress                    *string
	GuardianRegencyId                  *uint32
	GuardianRegencyName                *string
	GuardianPostalCode                 *string
	GuardianPhoneNumber                *string
	GuardianEmail                      *string
	GuardianFinalAcademicBackground    *string
	GuardianOccupation                 *string
	FatherIdNumber                     *string
	FatherName                         *string
	FatherBirthDate                    *time.Time
	FatherDeathDate                    *time.Time
	MotherIdNumber                     *string
	MotherName                         *string
	MotherBirthDate                    *time.Time
	MotherDeathDate                    *time.Time
	ParentAddress                      *string
	ParentRegencyId                    *uint32
	ParentRegencyName                  *string
	ParentPostalCode                   *string
	ParentPhoneNumber                  *string
	ParentEmail                        *string
	FatherFinalAcademicBackground      *string
	FatherOccupation                   *string
	MotherFinalAcademicBackground      *string
	MotherOccupation                   *string
	ParentIncome                       *float64
	IsFinanciallyCapable               *bool
	AuthenticationId                   *string
	AuthenticationIsActive             *bool
	AuthenticationSuspensionRemarks    *string
	DiktiStudyProgramCode              *string
	AcademicGuidanceLecturerId         *string
	AcademicGuidanceLecturerName       *string
	AcademicGuidanceSemesterId         *string
	AcademicGuidanceSemesterSchoolYear *string
	StudyPlanId                        *string
	StudyPlanTotalMandatoryCredit      *uint32
	StudyPlanTotalOptionalCredit       *uint32
	StudyPlanMaximumCredit             *uint32
	StudyPlanIsApproved                *bool
	CurrentSemesterPackage             uint32
	TotalStudyPlan                     uint32
	StatusSemesterId                   *string
	StatusSemesterSchoolYear           string
	StatusSemesterType                 *string
	StatusDate                         *time.Time
	StatusReferenceNumber              *string
	StatusPurpose                      *string
	StatusRemarks                      *string
	BloodType                          *string
	ProfilePhotoPath                   *string
	ProfilePhotoPathType               *string
	ProfilePhotoUrl                    string
	BirthProvinceId                    *uint32
	BirthProvinceName                  *string
	Height                             *float64
	Weight                             *float64
	IsColorBlind                       *bool
	UseGlasses                         *bool
	HasCompleteTeeth                   *bool
	IsKpsRecipient                     *bool
	WorkAddress                        *string
	AssuranceNumber                    *string
	ParentReligion                     *string
	ParentNationality                  *string
	FatherWorkAddress                  *string
	ParentProvinceId                   *uint32
	ParentProvinceName                 *string
	GuardianProvinceId                 *uint32
	GuardianProvinceName               *string
	SchoolEnrollmentYear               *string
	SchoolEnrollmentClass              *string
	SchoolType                         *string
	SchoolProvinceId                   *uint32
	SchoolProvinceName                 *string
	SchoolStatus                       *string
	SchoolMathematicsFinalExamScore    *float64
	SchoolIndonesianFinalExamScore     *float64
	SchoolEnglishFinalExamScore        *float64
	SchoolMathematicsReportScore       *float64
	SchoolIndonesianReportScore        *float64
	SchoolEnglishReportScore           *float64
	Gpa                                *float64
	TotalCredit                        *uint32
	GraduationPredicate                *string
	TranscriptIsArchived               bool
	HasPaid                            bool
	StudyDurationMonth                 uint32
	ThesisDurationMonth                uint32
	ThesisDurationSemester             uint32
	PreHighSchoolHistories             []GetStudentPreHighSchoolHistory
}

type StudentListWithPagination struct {
	Pagination common.Pagination
	Data       []GetStudent
}

type CreateStudent struct {
	Name                            string
	Sex                             string
	MaritalStatus                   string
	BirthRegencyId                  uint32
	BirthDate                       time.Time
	Religion                        string
	Address                         string
	Rt                              string
	Rw                              string
	VillageId                       uint32
	PostalCode                      string
	PreviousAddress                 string
	IdNumber                        string
	NpwpNumber                      string
	NisnNumber                      string
	ResidenceType                   string
	TransportationMean              string
	KpsReceiver                     string
	PhoneNumber                     string
	MobilePhoneNumber               string
	Email                           string
	Homepage                        string
	WorkType                        string
	WorkPlace                       string
	Nationality                     string
	AskesNumber                     string
	TotalBrother                    uint32
	TotalSister                     uint32
	Hobby                           string
	Experience                      string
	TotalDependent                  uint32
	NimNumber                       int64
	StudentForce                    uint32
	AdmittanceSemester              string
	StudyProgramId                  string
	CurriculumId                    string
	AdmittanceTestNumber            string
	AdmittanceDate                  time.Time
	AdmittanceStatus                string
	TotalTransferCredit             uint32
	PreviousCollege                 string
	PreviousStudyProgram            string
	PreviousNimNumber               int64
	PreviousNimAdmittanceYear       string
	Status                          string
	IsForeignStudent                bool
	CollegeEntranceType             string
	ClassTime                       string
	FundSource                      string
	IsScholarshipGrantee            bool
	HasCompleteRequirement          bool
	SchoolCertificateType           string
	SchoolGraduationYear            string
	SchoolName                      string
	SchoolAccreditation             string
	SchoolAddress                   string
	SchoolMajor                     string
	SchoolCertificateNumber         string
	SchoolCertificateDate           time.Time
	TotalSchoolFinalExamSubject     uint32
	SchoolFinalExamScore            float64
	GuardianName                    string
	GuardianBirthDate               time.Time
	GuardianDeathDate               time.Time
	GuardianAddress                 string
	GuardianRegencyId               uint32
	GuardianPostalCode              string
	GuardianPhoneNumber             string
	GuardianEmail                   string
	GuardianFinalAcademicBackground string
	GuardianOccupation              string
	FatherIdNumber                  string
	FatherName                      string
	FatherBirthDate                 time.Time
	FatherDeathDate                 time.Time
	MotherIdNumber                  string
	MotherName                      string
	MotherBirthDate                 time.Time
	MotherDeathDate                 time.Time
	ParentAddress                   string
	ParentRegencyId                 uint32
	ParentPostalCode                string
	ParentPhoneNumber               string
	ParentEmail                     string
	FatherFinalAcademicBackground   string
	FatherOccupation                string
	MotherFinalAcademicBackground   string
	MotherOccupation                string
	ParentIncome                    float64
	IsFinanciallyCapable            bool
}

type UpdateStudent struct {
	Id                              string
	Name                            string
	Sex                             string
	MaritalStatus                   string
	BirthRegencyId                  uint32
	BirthDate                       time.Time
	Religion                        string
	Address                         string
	Rt                              string
	Rw                              string
	VillageId                       uint32
	PostalCode                      string
	PreviousAddress                 string
	IdNumber                        string
	NpwpNumber                      string
	NisnNumber                      string
	ResidenceType                   string
	TransportationMean              string
	KpsReceiver                     string
	PhoneNumber                     string
	MobilePhoneNumber               string
	Email                           string
	Homepage                        string
	WorkType                        string
	WorkPlace                       string
	Nationality                     string
	AskesNumber                     string
	TotalBrother                    uint32
	TotalSister                     uint32
	Hobby                           string
	Experience                      string
	TotalDependent                  uint32
	NimNumber                       int64
	StudentForce                    uint32
	AdmittanceSemester              string
	StudyProgramId                  string
	CurriculumId                    string
	AdmittanceTestNumber            string
	AdmittanceDate                  time.Time
	AdmittanceStatus                string
	TotalTransferCredit             uint32
	PreviousCollege                 string
	PreviousStudyProgram            string
	PreviousNimNumber               int64
	PreviousNimAdmittanceYear       string
	Status                          string
	IsForeignStudent                bool
	CollegeEntranceType             string
	ClassTime                       string
	FundSource                      string
	IsScholarshipGrantee            bool
	HasCompleteRequirement          bool
	SchoolCertificateType           string
	SchoolGraduationYear            string
	SchoolName                      string
	SchoolAccreditation             string
	SchoolAddress                   string
	SchoolMajor                     string
	SchoolCertificateNumber         string
	SchoolCertificateDate           time.Time
	TotalSchoolFinalExamSubject     uint32
	SchoolFinalExamScore            float64
	GuardianName                    string
	GuardianBirthDate               time.Time
	GuardianDeathDate               time.Time
	GuardianAddress                 string
	GuardianRegencyId               uint32
	GuardianPostalCode              string
	GuardianPhoneNumber             string
	GuardianEmail                   string
	GuardianFinalAcademicBackground string
	GuardianOccupation              string
	FatherIdNumber                  string
	FatherName                      string
	FatherBirthDate                 time.Time
	FatherDeathDate                 time.Time
	MotherIdNumber                  string
	MotherName                      string
	MotherBirthDate                 time.Time
	MotherDeathDate                 time.Time
	ParentAddress                   string
	ParentRegencyId                 uint32
	ParentPostalCode                string
	ParentPhoneNumber               string
	ParentEmail                     string
	FatherFinalAcademicBackground   string
	FatherOccupation                string
	MotherFinalAcademicBackground   string
	MotherOccupation                string
	ParentIncome                    float64
	IsFinanciallyCapable            bool
}

type BulkUpdateStatusStudent struct {
	StudentIds            []string
	Status                string
	StatusDate            time.Time
	StatusReferenceNumber string
	StatusPurpose         string
	StatusRemarks         string
}

type GetStatusSummaryStudentStatus struct {
	Status string
	Total  uint32
}

type GetStatusSummaryStudent struct {
	StudyProgramId        string
	StudyProgramName      string
	DiktiStudyProgramType string
	StudyLevelShortName   string
	Statuses              []GetStatusSummaryStudentStatus
}

type GetStudentSemesterSummary struct {
	SemesterId                         string
	SemesterSchoolYear                 string
	SemesterType                       string
	Status                             *string
	StudyProgramId                     *string
	StudyProgramName                   *string
	HasPaid                            bool
	AcademicGuidanceLecturerId         *string
	AcademicGuidanceLecturerName       *string
	AcademicGuidanceLecturerFrontTitle *string
	AcademicGuidanceLecturerBackDegree *string
	MaximumCredit                      uint32
	StudyPlanInputStartDate            time.Time
	StudyPlanInputEndDate              time.Time
	StudyPlanApprovalStartDate         time.Time
	StudyPlanApprovalEndDate           time.Time
	TotalMandatoryCreditTaken          uint32
	TotalOptionalCreditTaken           uint32
	Gpa                                float64
}

type UpdateStudentProfile struct {
	ProfilePhotoPath     string
	ProfilePhotoPathType string
	Sex                  string
	BirthRegencyId       uint32
	BloodType            string
	Height               float64
	Weight               float64
	IsColorBlind         bool
	UseGlasses           bool
	HasCompleteTeeth     bool
	IdNumber             string
	NpwpNumber           string
	NisnNumber           string
	Religion             string
	MaritalStatus        string
	Nationality          string
	VillageId            uint32
	Rt                   string
	Rw                   string
	PostalCode           string
	Address              string
	PhoneNumber          string
	MobilePhoneNumber    string
	Email                string
	TransportationMean   string
	IsKpsRecipient       bool
	FundSource           string
	IsScholarshipGrantee bool
	TotalBrother         uint32
	TotalSister          uint32
	WorkType             string
	WorkPlace            string
	WorkAddress          string
	AssuranceNumber      string
	Hobby                string
}

type UpdateStudentParentProfile struct {
	FatherIdNumber                  string
	FatherName                      string
	FatherBirthDate                 string
	FatherDeathDate                 string
	FatherFinalAcademicBackground   string
	FatherOccupation                string
	MotherIdNumber                  string
	MotherName                      string
	MotherBirthDate                 string
	MotherDeathDate                 string
	MotherFinalAcademicBackground   string
	MotherOccupation                string
	ParentReligion                  string
	ParentNationality               string
	ParentAddress                   string
	FatherWorkAddress               string
	ParentRegencyId                 uint32
	ParentPostalCode                string
	ParentPhoneNumber               string
	ParentEmail                     string
	IsFinanciallyCapable            bool
	ParentIncome                    float64
	TotalDependent                  uint32
	GuardianName                    string
	GuardianBirthDate               string
	GuardianDeathDate               string
	GuardianAddress                 string
	GuardianRegencyId               uint32
	GuardianPostalCode              string
	GuardianPhoneNumber             string
	GuardianEmail                   string
	GuardianFinalAcademicBackground string
	GuardianOccupation              string
}

type UpdateStudentSchoolProfilePreHighSchoolHistory struct {
	Level          string
	Name           string
	GraduationYear string
}

type UpdateStudentSchoolProfile struct {
	SchoolEnrollmentYear            string
	SchoolGraduationYear            string
	SchoolEnrollmentClass           string
	SchoolMajor                     string
	SchoolType                      string
	SchoolName                      string
	SchoolProvinceId                uint32
	SchoolAddress                   string
	SchoolCertificateNumber         string
	SchoolCertificateDate           string
	SchoolStatus                    string
	SchoolAccreditation             string
	SchoolFinalExamScore            float64
	SchoolMathematicsFinalExamScore float64
	SchoolIndonesianFinalExamScore  float64
	SchoolEnglishFinalExamScore     float64
	SchoolMathematicsReportScore    float64
	SchoolIndonesianReportScore     float64
	SchoolEnglishReportScore        float64
	PreHighSchoolHistories          []UpdateStudentSchoolProfilePreHighSchoolHistory
}

type GetStudentSubject struct {
	SubjectId               string
	SubjectCode             string
	SubjectName             string
	GradeSemesterId         string
	GradeSemesterSchoolYear string
	GradeSemesterType       string
	GradePoint              float64
	GradeCode               *string
	SubjectIsMandatory      bool
	SemesterPackage         uint32
	SubjectTotalCredit      uint32
	SubjectType             string
}

type StudentSubjectWithPagination struct {
	Pagination common.Pagination
	Data       []GetStudentSubject
}

type GetStudentPaymentLog struct {
	SemesterId         string
	SemesterType       string
	SemesterStartYear  uint32
	SemesterSchoolYear string
	CreatedAt          time.Time
}

type BulkCreateStudent struct {
	NimNumber                       int64
	Name                            string
	Sex                             string
	MaritalStatus                   string
	BirthRegencyId                  uint32
	BirthDate                       time.Time
	Religion                        string
	Address                         string
	Rt                              string
	Rw                              string
	VillageId                       uint32
	PostalCode                      string
	IdNumber                        string
	NisnNumber                      string
	MobilePhoneNumber               string
	Nationality                     string
	DiktiStudyProgramCode           string
	SchoolName                      string
	SchoolAddress                   string
	SchoolProvinceId                uint32
	SchoolMajor                     string
	SchoolType                      string
	SchoolGraduationYear            string
	FatherName                      string
	FatherIdNumber                  string
	FatherBirthDate                 time.Time
	FatherFinalAcademicBackground   string
	FatherOccupation                string
	MotherName                      string
	MotherIdNumber                  string
	MotherBirthDate                 time.Time
	MotherFinalAcademicBackground   string
	MotherOccupation                string
	GuardianName                    string
	GuardianIdNumber                string
	GuardianBirthDate               time.Time
	GuardianFinalAcademicBackground string
	GuardianOccupation              string
}