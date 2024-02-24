package utils

import (
	"database/sql"
	"strconv"

	"github.com/sccicitb/pupr-backend/constants"
)

// NullBooleanScan function to convert *bool to bool
// Params:
// a: *bool
// Returns bool
// if a is nil, returns false
func NullBooleanScan(a *bool) bool {
	if a == nil {
		return false
	}
	return *a
}

func NewNullBoolean(b bool) sql.NullBool {
	if !b {
		return sql.NullBool{}
	}
	return sql.NullBool{
		Bool:  b,
		Valid: true,
	}
}

func StringToBoolPointer(s string) (*bool, *constants.ErrorResponse) {
	if s == "" {
		return nil, nil
	}

	res, err := strconv.ParseBool(s)
	if err != nil {
		return nil, constants.ErrorInternalServer(err.Error())
	}

	return &res, nil
}
