package constants

type permissionKeyType int

const (
	AppName = "pupr-backend"
)

const (
	SsoUrlConfig    = "sso_url"
	SsoApiUrlConfig = "sso_api_url"
	SsoAppIdConfig  = "sso_app_id"
	SsoRedirectUrl  = "sso_redirect_url"
)

const (
	PmbApiKeyConfig    = "pmb_api_key"
	CareerApiKeyConfig = "career_api_key"
)

const (
	AppTypeRoot     = "root"
	AppTypeAdmin    = "admin"
	AppTypeLecturer = "lecturer"
	AppTypeStudent  = "student"
)

const (
	PermissionContextKey permissionKeyType = iota
)

const (
	FacultyBuilding = "faculty"
	MajorBuilding   = "major"
)

const (
	OddSemester  = "Ganjil"
	EvenSemester = "Genap"
)

const (
	DefaultMaximumCredit = 24
)

const (
	ScheduleStartHour uint32 = 600
	ScheduleEndHour   uint32 = 2400
	ScheduleStartDay  uint32 = 1
	ScheduleEndDay    uint32 = 7
)

const (
	MidtermExam = "midterm"
	EndtermExam = "endterm"
)

const (
	VocationalSchool = "SMK"
)
