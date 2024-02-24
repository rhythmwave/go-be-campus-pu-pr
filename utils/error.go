package utils

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/sccicitb/pupr-backend/constants"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GenerateGrpcError function to generate error for gRPC procedure from *constants.ErrorResponse
func GenerateGrpcError(errs *constants.ErrorResponse, customErrorFunc func(error) error) error {
	if customErrorFunc != nil {
		errs.Err = customErrorFunc(errs.Err)
	}

	errorStatus := status.New(errs.GrpcCode, errs.Err.Error())
	errorDetail := errdetails.ErrorInfo{
		Metadata: errs.ErrInfo,
	}

	errorStatus, err := errorStatus.WithDetails(&errorDetail)
	if err != nil {
		return err
	}

	return errorStatus.Err()
}

// ErrDuplicate function to handle duplicate data on database
// Params:
// m: error string
// Returns *constants.ErrorResponse
func ErrDuplicate(m string) *constants.ErrorResponse {
	if !strings.Contains(m, constants.PGDuplicateConstraint) {
		return nil
	}
	errs := strings.Split(m, constants.PGDuplicateConstraint)
	s := trimQuote(errs[1])
	return constants.Error(http.StatusInternalServerError, codes.Internal, constants.ErrDuplicateCustomCode, constants.ErrDuplicate, s)
}

// ErrForeignKey function to handle invalid foreign key value on database
// Params:
// m: error string
// Returns *constants.ErrorResponse
func ErrForeignKey(m string) *constants.ErrorResponse {
	if !strings.Contains(m, constants.PGForeignKeyConstraint) {
		return nil
	}
	errs := strings.Split(m, constants.PGForeignKeyConstraint)
	s := trimQuote(errs[1])
	return constants.Error(http.StatusInternalServerError, codes.Internal, constants.ErrForeignKeyCustomCode, constants.ErrForeignKey, s)
}

// ErrHttpClient function to handle error when accesing 3rd party
// Params:
// m: error string
// Returns *constants.ErrorResponse
func ErrHttpClient(m string) *constants.ErrorResponse {
	parsedUrl, err := url.Parse(m)
	if err != nil {
		return constants.Error(http.StatusInternalServerError, codes.Internal, constants.DefaultCustomErrorCode, err.Error())
	}

	return constants.Error(http.StatusInternalServerError, codes.Internal, constants.ErrHttpClientCustomCode, constants.ErrHttpClient, parsedUrl.Host)
}

// ErrTableDoesNotExist function to handle inexistent table on database
// Params:
// err: error
// Returns *constants.ErrorResponse
func ErrTableDoesNotExist(err error) *constants.ErrorResponse {
	strArray := strings.Split(err.Error(), "\"")
	if len(strArray) != 3 {
		return nil
	}

	if strings.Contains(err.Error(), fmt.Sprintf(constants.PGTableDoesNotExist, strArray[1])) {
		return constants.Error(http.StatusInternalServerError, codes.Internal, constants.ErrTableDoesNotExistCustomCode, constants.ErrTableDoesNotExist, strArray[1])
	}

	return nil
}

// ErrGetTokenFromRedis function to handle error when get token from redis
// Params:
// m: error string
// Returns *constants.ErrorResponse
func ErrGetTokenFromRedis(m string) *constants.ErrorResponse {
	httpStatusCode := http.StatusInternalServerError
	grpcStatusCode := codes.Internal

	if m == constants.RedisNilValue {
		httpStatusCode = http.StatusUnauthorized
		grpcStatusCode = codes.Unauthenticated
	}
	return constants.Error(httpStatusCode, grpcStatusCode, constants.ErrGetTokenFromRedisCustomCode, constants.ErrGetTokenFromRedis, m)
}

// ErrInvalidDateFormat function to handle invalid date format
// Params:
// m: error string
// Returns *constants.ErrorResponse
func ErrInvalidDateFormat(m string) *constants.ErrorResponse {
	return constants.Error(http.StatusBadRequest, codes.InvalidArgument, constants.ErrInvalidDateFormatCustomCode, constants.ErrInvalidDateFormat, m)
}

// ErrDataNotFound function to handle data not found
// Params:
// m: error string
// Returns *constants.ErrorResponse
func ErrDataNotFound(m string) *constants.ErrorResponse {
	return constants.Error(http.StatusNotFound, codes.NotFound, constants.ErrDataNotFoundCustomCode, constants.ErrDataNotFound, m)
}

// ErrInvalidOrderKey function to handle invalid order key
// Params:
// m: list error string
// Returns *constants.ErrorResponse
func ErrInvalidOrderKey(m []string) *constants.ErrorResponse {
	return constants.Error(http.StatusBadRequest, codes.InvalidArgument, constants.ErrInvalidOrderKeyCustomCode, constants.ErrInvalidOrderKey, m...)
}

// ErrInvalidOrder function to handle order other than asc and desc
// Params:
// m: list error string
// Returns *constants.ErrorResponse
func ErrInvalidOrder(m []string) *constants.ErrorResponse {
	return constants.Error(http.StatusBadRequest, codes.InvalidArgument, constants.ErrInvalidOrderCustomCode, constants.ErrInvalidOrder, m...)
}

// ErrEmptyValue function to handle empty request value
// Params:
// m: request key
// Returns *constants.ErrorResponse
func ErrEmptyValue(m string) *constants.ErrorResponse {
	return constants.Error(http.StatusBadRequest, codes.InvalidArgument, constants.ErrEmptyValueCustomCode, constants.ErrEmptyValue, m)
}

// ErrMailTooFrequent function to handle send mail too frequently
// Params:
// t: time countdown from function getremainingtime
// Returns *constants.ErrorResponse
func ErrMailTooFrequent(t TimeCountdown) *constants.ErrorResponse {
	var remainingTime []string
	if t.Days > 0 {
		remainingTime = append(remainingTime, fmt.Sprintf("%d %s", t.Days, constants.Days))
	}
	if t.Hours > 0 {
		remainingTime = append(remainingTime, fmt.Sprintf("%d %s", t.Hours, constants.Hours))
	}
	if t.Minutes > 0 {
		remainingTime = append(remainingTime, fmt.Sprintf("%d %s", t.Minutes, constants.Minutes))
	}
	if t.Seconds > 0 {
		remainingTime = append(remainingTime, fmt.Sprintf("%d %s", t.Seconds, constants.Seconds))
	}

	return constants.Error(http.StatusBadRequest, codes.InvalidArgument, constants.ErrTooFrequentMailCustomCode, constants.ErrTooFrequentMail, strings.Join(remainingTime, " "))
}
