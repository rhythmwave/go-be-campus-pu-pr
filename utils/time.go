package utils

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/sccicitb/pupr-backend/constants"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// TimeCountdown struct for time count down
type TimeCountdown struct {
	Total   int
	Days    int
	Hours   int
	Minutes int
	Seconds int
}

// GetRemainingTime function to get remaining time
// Params:
// currentTime: current time
// timeout: time out
// Returns TimeCountdown
func GetRemainingTime(currentTime, timeout time.Time) TimeCountdown {
	difference := timeout.Sub(currentTime)

	total := int(difference.Seconds())
	days := int(total / (60 * 60 * 24))
	hours := int(total / (60 * 60) % 24)
	minutes := int(total/60) % 60
	seconds := int(total % 60)

	return TimeCountdown{
		Total:   total,
		Days:    days,
		Hours:   hours,
		Minutes: minutes,
		Seconds: seconds,
	}
}

// NewNullTime function to handle nullable value in SQL
// Params:
// t: time to be checked
// Returns sql.NullString
func NewNullTime(t time.Time) sql.NullTime {
	if t.IsZero() {
		return sql.NullTime{}
	}
	return sql.NullTime{
		Time:  t,
		Valid: true,
	}
}

// NullTimeDurationScan function to handle pointer duration
// Params:
// t: duration
// defaultDuration default Duration
// Returns duration
func NullTimeDurationScan(t *time.Duration, defaultDuration time.Duration) time.Duration {
	if t != nil {
		if *t != 0 {
			return *t
		}
	}

	return defaultDuration
}

// NextResendTime function to calculate duration to be able to send next mail
// Params:
// lastSend: time email last sent
// interval: interval to resend mail
func NextResendTime(lastSend time.Time, interval *time.Duration) TimeCountdown {
	now := time.Now()

	resendLimitDuration := NullTimeDurationScan(interval, constants.DefaultOTPResendTimePeriod)

	timeLimit := lastSend.Add(resendLimitDuration)

	diff := GetRemainingTime(now, timeLimit)

	return diff
}

// GetHourMinuteIntFromTime function to get hour and minute as integer
func GetHourMinuteIntFromTime(t time.Time) (int, *constants.ErrorResponse) {
	result, err := strconv.Atoi(t.Format("1504"))
	if err != nil {
		return 0, constants.Error(http.StatusInternalServerError, codes.Internal, constants.DefaultCustomErrorCode, err.Error())
	}

	return result, nil
}

func GetHourAndMinuteFromIntTime(t int) (int, int, *constants.ErrorResponse) {
	var hour int
	var minute int

	stringTime := strconv.Itoa(t)
	hour, err := strconv.Atoi(stringTime[:len(stringTime)-2])
	if err != nil {
		return hour, minute, constants.Error(http.StatusInternalServerError, codes.Internal, constants.DefaultCustomErrorCode, err.Error())
	}
	minute, err = strconv.Atoi(stringTime[len(stringTime)-2:])
	if err != nil {
		return hour, minute, constants.Error(http.StatusInternalServerError, codes.Internal, constants.DefaultCustomErrorCode, err.Error())
	}

	return hour, minute, nil
}

// StringToTime function to handle pointer string nulled
// Params:
// s: time.Time to be checked
// Returns string
func SafetyDate(s *time.Time) string {
	if s == nil {
		return ""
	}
	var d time.Time = *s
	return d.Format(constants.DateRFC)
}

// SafetyTimestamppb function to return proto timestamp
// Params:
// t time.Time to check
// loc *time.Location current location
// Returns *timestamppb.Timestamp
func SafetyTimestamppb(t time.Time, loc *time.Location) *timestamppb.Timestamp {
	_parse, err := time.ParseInLocation("2006-01-02T15:04:05", t.Format("2006-01-02T15:04:05"), loc)
	if err != nil {
		fmt.Println("ERROR TIME STAMP", err)
		return nil
	}

	return timestamppb.New(_parse)
}

// NullTimeScan function to convert time pointer to time
// Params:
// value: time pointer
// Returns time
// if value is null, returns empty time
func NullTimeScan(value *time.Time) time.Time {
	if value == nil {
		return time.Time{}
	}

	return *value
}

// func CalculateBusinessDayDuration()
