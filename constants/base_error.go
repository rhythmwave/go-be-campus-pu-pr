package constants

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/mitchellh/mapstructure"
	"google.golang.org/grpc/codes"
)

// ErrorResponse struct for standard error returns
// Code: http status code, use http package to define
// Err: error
type ErrorResponse struct {
	HttpCode   int
	GrpcCode   codes.Code
	CustomCode string
	Err        error
	ErrInfo    map[string]string
}

// ErrorResponseInfo struct for additional error info
type ErrorResponseInfo struct {
	Code string `mapstructure:"code"`
	Data string `mapstructure:"data"`
}

const (
	// SuccessCode represents custom ametory code for success response
	SuccessCode = "000"
	// DefaultCustomErrorCode represents custom ametory code for internal server error
	DefaultCustomErrorCode = "999"
)

const (
	// PGDuplicateConstraint define duplicate unique value on postgres
	PGDuplicateConstraint = "pq: duplicate key value violates unique constraint "
	// PGForeignKeyConstraint define invalid foreign key value on postgres
	PGForeignKeyConstraint = " violates foreign key constraint "
	// JWTInvalidType define invalid type on postgres
	JWTInvalidType = "key is of invalid type"
	// RedisNilValue define no data on redis key
	RedisNilValue = "redis: nil"
	// PGTableDoesNotExist define inexistence of table
	PGTableDoesNotExist = "pq: relation \"%s\" does not exist"
	// PGNoRows define no rows in result set when use row.Scan()
	PGNoRows = "sql: no rows in result set"
)

const (
	// ErrDuplicate define standard error message for duplicate data
	ErrDuplicate = "Duplicate data already in database:"
	// ErrForeignKey define standard error message for inexistent foreign key value
	ErrForeignKey = "Data foreign key is not exists:"
	// ErrHttpClient define standard error message when accessing 3rd party
	ErrHttpClient = "Error when trying to access"
	// ErrInvalidDateFormat define standard error message for invalid date format
	ErrInvalidDateFormat = "Invalid Date Format:"
	// ErrDataNotFound define standard error message for data not found
	ErrDataNotFound = "No Data:"
	// ErrInvalidOrderKey define standard error message for order column
	ErrInvalidOrderKey = "Invalid Order Key:"
	// ErrInvalidOrder define standard error message for other than asc or desc
	ErrInvalidOrder = "Invalid Order:"
	// ErrGetTokenFromRedis define standard error message when get token from redis
	ErrGetTokenFromRedis = "Error Get Token from redis:"
	// ErrTableDoesNotExist define standard error message for inexistence of table
	ErrTableDoesNotExist = "Table does not exist:"
	// ErrEmptyValue define standard error message for empty request value
	ErrEmptyValue = "Following request should not be empty:"
	// ErrTooFrequentMail define standard error message for frequent email send request
	ErrTooFrequentMail = "Mail is sent too frequently. Please wait for"
	// ErrUnknownCondition define standard error unknown condition
	ErrUnknownCondition = "Something wrong at server"
	// ErrFileService define standard error message when get error from File Processing
	ErrFileService = "Error File Service:"
)

const (
	// ErrNameIsRequiredCustomCode represents custom ametory code for Name is required
	ErrNameIsRequiredCustomCode = "001"
	// ErrEmailOrPhoneNumberIsNotValidCustomCode represents custom ametory code for Email or phone number is not valid
	ErrEmailOrPhoneNumberIsNotValidCustomCode = "002"
	// ErrEmailIsNotValidCustomCode represents custom ametory code for Email is not valid
	ErrEmailIsNotValidCustomCode = "003"
	// ErrPhoneNumberIsNotValidCustomCode represents custom ametory code for phone number is not valid
	ErrPhoneNumberIsNotValidCustomCode = "004"
	// ErrIDIsRequiredCustomCode represents custom ametory code for ID user is required
	ErrIDIsRequiredCustomCode = "005"
	// ErrTokenIsRequiredCustomCode represents custom ametory code for Token is required
	ErrTokenIsRequiredCustomCode = "006"
	// ErrKeyIsNotInvalidTypeCustomCode represents custom ametory code for JWTInvalidTyp
	ErrKeyIsNotInvalidTypeCustomCode = "007"
	// ErrEligbleAccessCustomCode represents custom ametory code for No right to access the API
	ErrEligbleAccessCustomCode = "008"
	// ErrTokenInvalidCustomCode represents custom ametory code for Token is invalid
	ErrTokenInvalidCustomCode = "009"
	// ErrGetTokenToRedisCustomCode represents custom ametory code for Error Get Token from redis
	ErrGetTokenToRedisCustomCode = "010"
	// ErrTokenAlreadyExpiredCustomCode represents custom ametory code for Token Already Expired
	ErrTokenAlreadyExpiredCustomCode = "011"
	// ErrTokenReplacedCustomCode represents custom ametory code for Please re login for next process
	ErrTokenReplacedCustomCode = "012"
	// ErrEmailExistsCustomCode represents custom ametory code for Email is already exists
	ErrEmailExistsCustomCode = "013"
	// ErrNoDialerCustomCode represents custom ametory code for Dialer not set yet
	ErrNoDialerCustomCode = "014"
	// ErrSaveTokenToRedisCustomCode represents custom ametory code for Error Save Token to redis
	ErrSaveTokenToRedisCustomCode = "015"
	// ErrEmailAndPasswordNotMatchCustomCode represents custom ametory code for Email or Password Not Match
	ErrEmailAndPasswordNotMatchCustomCode = "016"
	// ErrIncorrectOTPCustomCode represents custom ametory code for OTP incorrect
	ErrIncorrectOTPCustomCode = "017"
	// ErrUniqueIDCustomCode represents custom ametory code for Unique ID Not Match
	ErrUniqueIDCustomCode = "018"
	// ErrBuildNumberCustomCode represents custom ametory code for Build Number Not Match
	ErrBuildNumberCustomCode = "019"
	// ErrAppIDCustomCode represents custom ametory code for App ID Not Match
	ErrAppIDCustomCode = "020"
	// ErrCertExpiredCustomCode represents custom ametory code for Certificate expired
	ErrCertExpiredCustomCode = "021"
	// ErrApproverUneligibleCustomCode represents custom ametory code for Approver not eligible to sign
	ErrApproverUneligibleCustomCode = "022"
	// ErrEmailTemplateCustomCode represents custom ametory code for Email should have template and/or html body
	ErrEmailTemplateCustomCode = "023"
	// ErrInvalidBloodTypeCustomCode represents custom ametory code for Invalid Blood Type
	ErrInvalidBloodTypeCustomCode = "024"
	// ErrInvalidGenderCustomCode represents custom ametory code for User has invalid gender value
	ErrInvalidGenderCustomCode = "025"
	// ErrUserNameAbsenceCustomCode represents custom ametory code for User has no name
	ErrUserNameAbsenceCustomCode = "026"
	// ErrMigrationTableNameNonAlphabeticCustomCode represents custom ametory code for Migration server name only allowed alphabetic character
	ErrMigrationTableNameNonAlphabeticCustomCode = "027"
	// ErrUndoDefaultMigrationCustomCode represents custom ametory code for undo default migration
	ErrUndoDefaultMigrationCustomCode = "028"
	// ErrDuplicateCustomCode represents custom ametory code for duplicate data
	ErrDuplicateCustomCode = "029"
	// ErrForeignKeyCustomCode represents custom ametory code for inexistent foreign key value
	ErrForeignKeyCustomCode = "030"
	// ErrHttpClientCustomCode represents custom ametory code for accessing 3rd party
	ErrHttpClientCustomCode = "031"
	// ErrInvalidDateFormatCustomCode represents custom ametory code for invalid date format
	ErrInvalidDateFormatCustomCode = "032"
	// ErrDataNotFoundCustomCode represents custom ametory code for data not found
	ErrDataNotFoundCustomCode = "033"
	// ErrInvalidOrderKeyCustomCode represents custom ametory code for order column
	ErrInvalidOrderKeyCustomCode = "034"
	// ErrInvalidOrderCustomCode represents custom ametory code for other than asc or desc
	ErrInvalidOrderCustomCode = "035"
	// ErrGetTokenFromRedisCustomCode represents custom ametory code for get token from redis
	ErrGetTokenFromRedisCustomCode = "036"
	// ErrTableDoesNotExistCustomCode represents custom ametory code for inexistence of table
	ErrTableDoesNotExistCustomCode = "036"
	// ErrEmptyValueCustomCode represents custom ametory code for empty request value
	ErrEmptyValueCustomCode = "037"
	// ErrSaveOTPToRedisCustomCode represents custom ametory code for Error Save OTP to redis
	ErrSaveOTPToRedisCustomCode = "038"
	// ErrGetOTPFromRedisCustomCode represents custom ametory code for OTP is invalid or has been expired.
	ErrGetOTPFromRedisCustomCode = "039"
	// ErrDeleteOTPFromRedisCustomCode represents custom ametory code for Error Delete OTP from redis
	ErrDeleteOTPFromRedisCustomCode = "040"
	// ErrOTPIsRequiredCustomCode represents custom ametory code for OTP is required
	ErrOTPIsRequiredCustomCode = "041"
	// ErrOTPNotMatchCustomCode represents custom ametory code for OTP doesn't match
	ErrOTPNotMatchCustomCode = "042"
	// ErrTooFrequentMailCustomCode represents custom ametory code for email too frequest
	ErrTooFrequentMailCustomCode = "043"
	// ErrGenerateTokenFromJwtCode represents custom ametory code for generate token from jwt
	ErrGenerateTokenFromJwtCode = "044"
	// ErrGenerateFirestoreCustomTokenFromJwtCode represents custom ametory code for generate custom token from jwt
	ErrGenerateFirestoreCustomTokenFromJwtCode = "045"
	// ErrSetTokenToRedisCode represents custom ametory code for set token to redis code
	ErrSetTokenToRedisCode = "046"
	// ErrPasswordNotNullCode represents custom ametory code for update user password not null
	ErrPasswordNotNullCode = "047"
	// ErrEmptyMigrationListCode represents custom ametory code for empty migration list
	ErrEmptyMigrationListCode = "048"
	// ErrInvalidClaimsCode represents custom ametory code for invalid JWT claims
	ErrInvalidClaimsCode = "049"
	// ErrFileServiceCustomCode represents custom ametory code for get token from redis
	ErrFileServiceCustomCode = "050"
	// ErrClockinOutsideTimeRange represents custom ametory code for clockin outside of allowed time range
	ErrClockinOutsideTimeRangeCode = "051"
	// ErrOvertimeInSchedule represents custom ametory code for overtime clockin inside scheduled work time
	ErrOvertimeInScheduleCode = "052"
	// ErrInvalidCompanyEmployee represents custom ametory code for invalid action for company
	ErrInvalidCompanyEmployeeCode = "053"
	// ErrNonAdminEmployeeCode represents custom ametory code for non admin employee take action
	ErrNonAdminEmployeeCode = "054"
	// ErrActiveAttendanceExistsCode represents custom ametory code for attendance existance as clockin
	ErrActiveAttendanceExistsCode = "055"
	// ErrMaxDailyAttendanceCountCode represents custom ametory code for exceeding daily attendance
	ErrMaxDailyAttendanceCountCode = "056"
	// ErrActiveAttendanceNotExistsCode represents custom ametory code for attendance inexistance as clockout
	ErrActiveAttendanceNotExistsCode = "056"
	// ErrClockoutOutsideTimeRangeCode represents custom ametory code for clockout outside of allowed time range
	ErrClockoutOutsideTimeRangeCode = "057"
	// ErrEmailPhoneNumberAlreadyValidatedCode represents custom ametory code for email or phone number already validated
	ErrEmailPhoneNumberAlreadyValidatedCode = "058"
	// ErrNoEmployeeProfilePictureCode represents custom ametory code when employee has no profile picture
	ErrNoEmployeeProfilePictureCode = "059"
	// ErrFaceNotMatchCode represents custom ametory code for face unmatch
	ErrFaceNotMatchCode = "060"
	// ErrPasswordNotMatchCode represents custom ametory code for Password unmatch
	ErrPasswordNotMatchCode = "061"
	// ErrSendNotificationCode represents custom ametory code for error when sending push notification
	ErrSendNotificationCode = "062"
	// ErrPhoneNumberExistsCustomCode represents custom ametory code for PhoneNumber is already exists
	ErrPhoneNumberExistsCustomCode = "063"
	// ErrPendingUserDeletionRequestCustomCode represents custom ametory code for pending user deletion request is exists
	ErrPendingUserDeletionRequestCustomCode = "064"
	// ErrUserPasswordNotNullCustomCode represents custom ametory code for PhoneNumber is already exists
	ErrUserPasswordNotNullCustomCode = "065"
	// ErrUsernameExistsCustomCode represents custom ametory code for Username is already exists
	ErrUsernameExistsCustomCode = "066"
	// ErrInvalidPlatformCustomCode represents custom ametory code for invalid url for social media platform
	ErrInvalidPlatformCustomCode = "067"
	// ErrUnauthenticatedCustomCode represents custom ametory code for error Email or Password Not Match
	ErrUnauthenticatedCustomCode = "068"
	// ErrInvalidCompanyIdCustomCode represents custom ametory code for error invalid companyID on header
	ErrInvalidCompanyIdCustomCode = "069"
	// ErrAttendanceOutsideRadiusCustomCode represents custom ametory code for error clockin/clockout outside allowed location
	ErrAttendanceOutsideRadiusCustomCode = "070"
	// ErrUnknownCustomCode represents Unknown condition
	ErrUnknownCustomCode = "999"
)

// Error function to generate ErrorResponse
// With parameters as following:
// httpCode: http status code, use http package for defining
// grpcCode: grpc status code
// message: error message
// data: sent data if available, if not, send nil value
// args: arguments for error message
// Returns *ErrorMessage
func Error(httpCode int, grpcCode codes.Code, customCode, message string, args ...string) *ErrorResponse {
	for _, v := range args {
		message = fmt.Sprintf("%s %s", message, v)
	}

	errData := ErrorResponseInfo{
		Code: customCode,
		Data: strings.Join(args, ", "),
	}
	var errInfo map[string]string
	err := mapstructure.Decode(errData, &errInfo)
	if err != nil {
		errs := ErrorResponse{
			HttpCode:   http.StatusInternalServerError,
			GrpcCode:   codes.Internal,
			CustomCode: DefaultCustomErrorCode,
			Err:        err,
			ErrInfo:    errInfo,
		}
		return &errs
	}

	errs := ErrorResponse{
		HttpCode:   httpCode,
		GrpcCode:   grpcCode,
		CustomCode: customCode,
		Err:        errors.New(message),
		ErrInfo:    errInfo,
	}

	return &errs
}

// ErrorInternalServer This is wrapper for error internal server only
func ErrorInternalServer(message string, args ...string) *ErrorResponse {
	return Error(http.StatusInternalServerError, codes.Internal, ErrUnknownCustomCode, message, args...)
}

var (
	// ErrNameIsRequired predefine error Name is required
	ErrNameIsRequired = Error(http.StatusBadRequest, codes.InvalidArgument, ErrNameIsRequiredCustomCode, "Name is required")
	// ErrEmailOrPhoneNumberIsNotValid predefine error Email or phone number is not valid
	ErrEmailOrPhoneNumberIsNotValid = Error(http.StatusBadRequest, codes.InvalidArgument, ErrEmailOrPhoneNumberIsNotValidCustomCode, "Email or phone number is not valid")
	// ErrEmailIsNotValid predefine error Email is not valid
	ErrEmailIsNotValid = Error(http.StatusBadRequest, codes.InvalidArgument, ErrEmailIsNotValidCustomCode, "Email is not valid")
	// ErrPhoneNumberIsNotValid predefine error phone number is not valid
	ErrPhoneNumberIsNotValid = Error(http.StatusBadRequest, codes.InvalidArgument, ErrPhoneNumberIsNotValidCustomCode, "Phone number is not valid")
	// ErrIDIsRequired predefine error ID user is required
	ErrIDIsRequired = Error(http.StatusBadRequest, codes.InvalidArgument, ErrIDIsRequiredCustomCode, "ID user is required")
	// ErrTokenIsRequired predefine error Token is required
	ErrTokenIsRequired = Error(http.StatusUnauthorized, codes.Unauthenticated, ErrTokenIsRequiredCustomCode, "Token is required")
	// ErrKeyIsNotInvalidType predefine error JWTInvalidType
	ErrKeyIsNotInvalidType = Error(http.StatusInternalServerError, codes.Internal, ErrKeyIsNotInvalidTypeCustomCode, JWTInvalidType)
	// ErrEligbleAccess predefine error No right to access the API
	ErrEligbleAccess = Error(http.StatusForbidden, codes.PermissionDenied, ErrEligbleAccessCustomCode, "No right to access the API")
	// ErrTokenInvalid predefine error Token is invalid
	ErrTokenInvalid = Error(http.StatusUnauthorized, codes.Unauthenticated, ErrTokenInvalidCustomCode, "Token is invalid")
	// ErrGetTokenToRedis predefine error Error Get Token from redis
	ErrGetTokenToRedis = Error(http.StatusUnauthorized, codes.Unauthenticated, ErrGetTokenToRedisCustomCode, "Error Get Token from redis")
	// ErrTokenAlreadyExpired predefine error Token Already Expired
	ErrTokenAlreadyExpired = Error(http.StatusUnauthorized, codes.Unauthenticated, ErrTokenAlreadyExpiredCustomCode, "Token Already Expired")
	// ErrTokenReplaced predefine error Please re login for next process
	ErrTokenReplaced = Error(http.StatusUnauthorized, codes.Unauthenticated, ErrTokenReplacedCustomCode, "Please re login for next process")
	// ErrEmailExists predefine error Email is already exists
	ErrEmailExists = Error(http.StatusBadRequest, codes.InvalidArgument, ErrEmailExistsCustomCode, "Email is already exists")
	// ErrNoDialer predefine error Dialer not set yet
	ErrNoDialer = Error(http.StatusInternalServerError, codes.Internal, ErrNoDialerCustomCode, "Dialer not set yet")
	// ErrSaveTokenToRedis predefine error Error Save Token to redis
	ErrSaveTokenToRedis = Error(http.StatusInternalServerError, codes.Internal, ErrSaveTokenToRedisCustomCode, "Error Save Token to redis")
	// ErrEmailAndPasswordNotMatch predefine error Email or Password Not Match
	ErrEmailAndPasswordNotMatch = Error(http.StatusUnauthorized, codes.Unauthenticated, ErrEmailAndPasswordNotMatchCustomCode, "Email or Password Not Match")
	// ErrUnauthenticated predefine error Email or Password Not Match
	ErrUnauthenticated = Error(http.StatusUnauthorized, codes.Unauthenticated, ErrUnauthenticatedCustomCode, "Unauthenticated")
	// ErrIncorrectOTP predefine error OTP incorrect
	ErrIncorrectOTP = Error(http.StatusUnauthorized, codes.Unauthenticated, ErrIncorrectOTPCustomCode, "OTP incorrect")
	// ErrUniqueID predefine error Unique ID Not Match
	ErrUniqueID = Error(http.StatusBadRequest, codes.InvalidArgument, ErrUniqueIDCustomCode, "Unique ID Not Match")
	// ErrBuildNumber predefine error Build Number Not Match
	ErrBuildNumber = Error(http.StatusBadRequest, codes.InvalidArgument, ErrBuildNumberCustomCode, "Build Number Not Match")
	// ErrAppID predefine error App ID Not Match
	ErrAppID = Error(http.StatusBadRequest, codes.InvalidArgument, ErrAppIDCustomCode, "App ID Not Match")
	// ErrCertExpired predefine error Certificate expired
	ErrCertExpired = Error(http.StatusBadRequest, codes.InvalidArgument, ErrCertExpiredCustomCode, "Certificate expired")
	// ErrApproverUneligible predefine error Approver not eligible to sign
	ErrApproverUneligible = Error(http.StatusBadRequest, codes.InvalidArgument, ErrApproverUneligibleCustomCode, "Approver not eligible to sign")
	// ErrEmailTemplate predefine error Email should have template and/or html body
	ErrEmailTemplate = Error(http.StatusInternalServerError, codes.Internal, ErrEmailTemplateCustomCode, "Email should have template and/or html body")
	// ErrInvalidBloodType predefine error Invalid Blood Type
	ErrInvalidBloodType = Error(http.StatusBadRequest, codes.InvalidArgument, ErrInvalidBloodTypeCustomCode, "Invalid Blood Type")
	// ErrInvalidGender predefine error User has invalid gender value
	ErrInvalidGender = Error(http.StatusBadRequest, codes.InvalidArgument, ErrInvalidGenderCustomCode, "User has invalid gender value")
	// ErrUserNameAbsence predefine error User has no name
	ErrUserNameAbsence = Error(http.StatusBadRequest, codes.InvalidArgument, ErrUserNameAbsenceCustomCode, "User has no name")
	// ErrMigrationTableNameNonAlphabetic predefine error Migration server name only allowed alphabetic character
	ErrMigrationTableNameNonAlphabetic = Error(http.StatusInternalServerError, codes.Internal, ErrMigrationTableNameNonAlphabeticCustomCode, "Migration server name only allowed alphabetic character")
	// ErrUndoDefaultMigration predefine error undo default migration
	ErrUndoDefaultMigration = Error(http.StatusInternalServerError, codes.Internal, ErrUndoDefaultMigrationCustomCode, "Default migration cannot be undone")
	// ErrSaveOTPToRedis predefine error Error Save OTP to redis
	ErrSaveOTPToRedis = Error(http.StatusInternalServerError, codes.Internal, ErrSaveOTPToRedisCustomCode, "Error Save OTP to redis")
	// ErrGetOTPFromRedis predefine error OTP is invalid or has been expired
	ErrGetOTPFromRedis = Error(http.StatusBadRequest, codes.InvalidArgument, ErrGetOTPFromRedisCustomCode, "OTP is invalid or has been expired.")
	// ErrDeleteOTPFromRedis predefine error Error Delete OTP from redis
	ErrDeleteOTPFromRedis = Error(http.StatusInternalServerError, codes.Internal, ErrDeleteOTPFromRedisCustomCode, "Error Delete OTP from redis")
	// ErrOTPIsRequired predefine error OTP is required
	ErrOTPIsRequired = Error(http.StatusBadRequest, codes.InvalidArgument, ErrOTPIsRequiredCustomCode, "OTP is required")
	// ErrOTPNotMatch predefine error OTP doesn't match
	ErrOTPNotMatch = Error(http.StatusBadRequest, codes.InvalidArgument, ErrOTPNotMatchCustomCode, "OTP doesn't match")
	// ErrGenerateToken predefine err generate from jwt
	ErrGenerateToken = Error(http.StatusBadRequest, codes.InvalidArgument, ErrGenerateTokenFromJwtCode, "err generate from jwt")
	// ErrGenerateFirestoreCustomToken predefine err generate from jwt
	ErrGenerateFirestoreCustomToken = Error(http.StatusBadRequest, codes.InvalidArgument, ErrGenerateFirestoreCustomTokenFromJwtCode, "err generate custom token from jwt")
	// ErrSetTokenToRedis predefine err for set to redis
	ErrSetTokenToRedis = Error(http.StatusBadRequest, codes.InvalidArgument, ErrSetTokenToRedisCode, "err set token to redis")
	// ErrPasswordNotNull predefine err for request password nil
	ErrPasswordNotNull = Error(http.StatusBadRequest, codes.InvalidArgument, ErrPasswordNotNullCode, "err set pasword but is null value")
	// ErrEmptyMigrationList predefine migration list cannot be empty
	ErrEmptyMigrationList = Error(http.StatusInternalServerError, codes.Internal, ErrEmptyMigrationListCode, "migration list cannot be empty")
	// ErrInvalidClaims predefine invalid JWT claims
	ErrInvalidClaims = Error(http.StatusInternalServerError, codes.Internal, ErrInvalidClaimsCode, "invalid claims")
	// ErrUnknown predefine invalid JWT claims
	ErrUnknown = Error(http.StatusInternalServerError, codes.Internal, ErrUnknownCustomCode, ErrUnknownCondition)
	// ErrClockinNotInTimeRange predefine error when user clockin outside allowed time range
	ErrClockinOutsideTimeRange = Error(http.StatusBadRequest, codes.InvalidArgument, ErrClockinOutsideTimeRangeCode, "Cannot clockin outside allowed time range")
	// ErrOvertimeInSchedule predefine error when user clockin as overtime inside workhour
	ErrOvertimeInSchedule = Error(http.StatusBadRequest, codes.InvalidArgument, ErrOvertimeInScheduleCode, "Overtime is not allowed inside work hour time range")
	// ErrInvalidCompanyEmployee predefine error when user take action for another company
	ErrInvalidCompanyEmployee = Error(http.StatusForbidden, codes.PermissionDenied, ErrInvalidCompanyEmployeeCode, "Action is invalid for this company")
	// ErrNonAdminEmployee predefine error when non admin employee take action
	ErrNonAdminEmployee = Error(http.StatusBadRequest, codes.InvalidArgument, ErrNonAdminEmployeeCode, "You are not an admin for this company")
	// ErrActiveAttendanceExists predefine error when there's active attendance as clockin
	ErrActiveAttendanceExists = Error(http.StatusBadRequest, codes.InvalidArgument, ErrActiveAttendanceExistsCode, "Please clockout first.")
	// ErrMaxDailyAttendanceCount predefine error when attendance exceeding max daily clockin
	ErrMaxDailyAttendanceCount = Error(http.StatusBadRequest, codes.InvalidArgument, ErrMaxDailyAttendanceCountCode, "You cannot clockin again today.")
	// ErrActiveAttendanceNotExists predefine error when there's no active attendance as clockout
	ErrActiveAttendanceNotExists = Error(http.StatusBadRequest, codes.InvalidArgument, ErrActiveAttendanceNotExistsCode, "Please clockin first.")
	// ErrClockoutOutsideTimeRange predefine error when user clockout outside allowed time range
	ErrClockoutOutsideTimeRange = Error(http.StatusBadRequest, codes.InvalidArgument, ErrClockoutOutsideTimeRangeCode, "Cannot clockout outside allowed time range")
	// ErrEmailPhoneNumberAlreadyValidated predefine error when user clockout outside allowed time range
	ErrEmailPhoneNumberAlreadyValidated = Error(http.StatusBadRequest, codes.InvalidArgument, ErrEmailPhoneNumberAlreadyValidatedCode, "Email or phone number is already validated")
	// ErrNoEmployeeProfilePicture predefine error when employee doesn't have profile picture
	ErrNoEmployeeProfilePicture = Error(http.StatusBadRequest, codes.InvalidArgument, ErrNoEmployeeProfilePictureCode, "Employee does not have profile picture")
	// ErrFaceNotMatch predefine error when face not match
	ErrFaceNotMatch = Error(http.StatusBadRequest, codes.InvalidArgument, ErrFaceNotMatchCode, "Face not match")
	// ErrPasswordNotMatch predefine error when Password not match
	ErrPasswordNotMatch = Error(http.StatusBadRequest, codes.InvalidArgument, ErrPasswordNotMatchCode, "Password not match")
	// ErrSendNotification predefine error when Password not match
	ErrSendNotification = Error(http.StatusInternalServerError, codes.Internal, ErrSendNotificationCode, "Error when sending notifications")
	// ErrPhoneNumberExists predefine error PhoneNumber is already exists
	ErrPhoneNumberExists = Error(http.StatusBadRequest, codes.InvalidArgument, ErrPhoneNumberExistsCustomCode, "Phone number is already exists")
	// ErrPendingUserDeletionRequest predefine error pending user deletion request is exists
	ErrPendingUserDeletionRequest = Error(http.StatusBadRequest, codes.InvalidArgument, ErrPendingUserDeletionRequestCustomCode, "You have requested to delete your account. Please wait for admin's response")
	// ErrUsernameExists predefine error Username is already exists
	ErrUsernameExists = Error(http.StatusBadRequest, codes.InvalidArgument, ErrUsernameExistsCustomCode, "Username is already exists")
	// ErrInvalidPlatform predefine error Invalid url for this platform
	ErrInvalidPlatform = Error(http.StatusBadRequest, codes.InvalidArgument, ErrInvalidPlatformCustomCode, "Invalid social media platform. Use only Detik, Facebook, Instagram, Tiktok, or Youtube")
	// ErrInvalidCompanyId predefine invalid companyID on header
	ErrInvalidCompanyId = Error(http.StatusBadRequest, codes.InvalidArgument, ErrInvalidCompanyIdCustomCode, "Invalid header company ID")
	// ErrAttendanceOutsideRadius predefine invalid companyID on header
	ErrAttendanceOutsideRadius = Error(http.StatusBadRequest, codes.InvalidArgument, ErrAttendanceOutsideRadiusCustomCode, "Cannot clockin/clockout outside designated location")
)
