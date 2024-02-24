package constants

import "time"

type contextKeyType int

const (
	// Bearer variable get from header auth
	Bearer = "Bearer"
	// ClaimsContextKey context key for jwt claims
	ClaimsContextKey contextKeyType = iota
	// FlutterNotificationClick default click action for mobile device
	FlutterNotificationClick = "FLUTTER_NOTIFICATION_CLICK"
	// RoleUser
	RoleUser = "user"
)

const (
	// Authorization represents authorization header key
	Authorization = "Authorization"
	// DateFormatStd represents standard date format: YYYY-MM-DD
	DateFormatStd = "2006-01-02"
	// TimeFormatStd represents standard time format: HH:mm:ss
	TimeFormatStd = "15:04:05"
	// TimeFormatShort represents standard time format without second: HH:mm
	TimeFormatShort = "15:04"
	// DateFromStd represents standard full datetime format: YYYY-MM-DD HH:mm:ss
	DateFromStd = "2006-01-02 15:04:05"
	// DateRFC represents standard full datetime format: 2006-01-02T15:04:05-0700
	DateRFC = "2006-01-02T15:04:05-0700"
	// DefaultPassword represents default passwords
	DefaultPassword = "12345678"
	// DefaultPhoneAreaCode represents Indonesia phone area code
	DefaultPhoneAreaCode = "62"
	// HttpRequestFormFileKey represents default upload file key
	HttpRequestFormFileKey = "file"
	// HttpRequestFormFolderName represents default folder name in S3
	HttpRequestFormFolderName = "type"
	// HttpClientDefaultTimeout represents http client default timeout
	HttpClientDefaultTimeout = 60
	// TempDownloadedFileDir represents default temporary file directory
	TempDownloadedFileDir = "./temp/"
	// LogDir represents default log file directory
	LogDir = "./log/"
	// MultipartFormMaxMemory represents max memeory for multipart/form
	MultipartFormMaxMemory = 2000
	// IndonesianRupiahCode represents indonesian rupiah code
	IndonesianRupiahCode = "IDR"
	// WIBTimezone represents WIB timezone
	WIBTimezone = "WIB"
	// DefaultAttendancePolicyTemplate represent default name for attendance policy template
	DefaultAttendancePolicyTemplate = "Default Policy"
	// DefaultSingleSendNotification represent default total token for single send notification
	DefaultSingleSendNotification = 500
	// DefaultTimezone Asia/Jakarta
	DefaultTimezone = "Asia/Jakarta"
)

// blood types
const (
	// APlusBloodType represent blood type A+
	APlusBloodType = "A+"
	// BPlusBloodType represent blood type B+
	BPlusBloodType = "B+"
	// ABPlusBloodType represent blood type AB+
	ABPlusBloodType = "AB+"
	// OPlusBloodType represent blood type O+
	OPlusBloodType = "O+"
	// AMinusBloodType represent blood type A-
	AMinusBloodType = "A-"
	// BMinusBloodType represent blood type B-
	BMinusBloodType = "B-"
	// ABMinusBloodType represent blood type AB-
	ABMinusBloodType = "AB-"
	// OMinusBloodType represent blood type O-
	OMinusBloodType = "O-"
)

// important location
const (
	// IndonesiaCountry Indonesia Country Name
	IndonesiaCountry = "Indonesia"
	// IndonesiaAlphaCode Indonesia Alpha-2 Code
	IndonesiaAlphaCode = "ID"
	// JakartaTimezone timezone for asia/jakarta
	JakartaTimezone = "Asia/Jakarta"
	// oneday in seconds
	Oneday = time.Second * 3600 * 24
)

// day of week int code
const (
	// SundayCode code of Sunday
	SundayCode = 0
	// MondayCode code of Monday
	MondayCode = 1
	// TuesdayCode code of Tuesday
	TuesdayCode = 2
	// WednesdayCode code of Wednesday
	WednesdayCode = 3
	// ThursdayCode code of Thursday
	ThursdayCode = 4
	// FridayCode code of Friday
	FridayCode = 5
	// SaturdayCode code of Saturday
	SaturdayCode = 6
	// SundayCodeIso ISO code of Sunday
	SundayCodeIso = 7
)

const (
	Fcm = "fcm"
)

const (
	Instagram    = "instagram"
	Facebook     = "facebook"
	Twitter      = "twitter"
	Youtube      = "youtube"
	Tiktok       = "tiktok"
	Threads      = "threads"
	Detik        = "detik"
	ShortYoutube = "youtu"
	X            = "x.com"
)

const (
	Previous = "previous"
	Today    = "today"
	Upcoming = "upcoming"
)

const (
	UserSexMale      = "M"
	UserSexFemale    = "F"
	UserSexUndefined = "0"
)

const (
	LeaveTypeLeave  = "leave"
	LeaveTypePermit = "permit"
)

const (
	MaritalStatusMarried = "married"
	MaritalStatusSingle  = "single"
	MaritalStatusBoth    = "both"
)

const (
	LeaveRequestStatusPendingApproval = "pending_approval"
	LeaveRequestStatusApproved        = "approved"
	LeaveRequestStatusOngoing         = "ongoing"
	LeaveRequestStatusCompleted       = "completed"
	LeaveRequestStatusDeclined        = "declined"
	LeaveRequestStatusCancelled       = "cancelled"
)
