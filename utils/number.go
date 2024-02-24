package utils

import (
	"database/sql"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var (
	numberRunes = []rune("1234567890")
)

// NullIntScan function to convert integer pointer to integer
// Params:
// a: integer pointer
// Returns integer
// if a is nil, returns 0
func NullIntScan(a *int) int {
	if a != nil {
		return *a
	}

	return 0
}

// NullUint32Scan function to convert unsign integer pointer to unsign integer
// Params:
// a: unsign integer pointer
// Returns unsign integer
// if a is nil, returns 0
func NullUint32Scan(a *uint32) uint32 {
	if a != nil {
		return *a
	}

	return 0
}

// NullUint32Scan function to convert unsign integer pointer to unsign integer
// Params:
// a: unsign integer pointer
// Returns unsign integer
// if a is nil, returns 0
func NullInt64Scan(a *int64) int64 {
	if a != nil {
		return *a
	}

	return 0
}

// NullFloatScan function to convert float64 pointer to float64
// Params:
// a: float64 pointer
// Returns float64
// if a is nil, returns 0
func NullFloatScan(a *float64) float64 {
	if a != nil {
		return *a
	}

	return 0.0
}

// NullFloat32Scan function to convert float32 pointer to float32
// Params:
// a: float32 pointer
// Returns float32
// if a is nil, returns 0
func NullFloat32Scan(a *float32) float32 {
	if a != nil {
		return *a
	}
	return 0
}

// ScanIntToNullValue function to convert integer to integer pointer
// Params:
// a: integer
// Returns integer pointer
// if a == 0, returns nil
func ScanIntToNullValue(a int) *int {
	if a == 0 {
		return nil
	}

	return &a
}

// NullFloat64ScanFromNullableString function to convert string pointer to float64
// Params:
// a: string pointer
// Returns float64
// if a == nil or not numeric, returns 0
func NullFloat64ScanFromNullableString(a *string) float64 {
	if a != nil {
		value, err := strconv.ParseFloat(*a, 64)
		if err != nil {
			return 0.0
		}
		return value
	}
	return 0.0
}

// CountTotalPage function to count total page
// Params:
// total: total data
// perPage: data per page
// Returns total page
func CountTotalPage(total, perPage int) int {
	if (total % perPage) > 0 {
		return (total / perPage) + 1
	}
	return total / perPage
}

// CommaSeparated function to convert float64 to comma separated string
// Params:
// v: float64
// Returns comma separated string
func CommaSeparated(v float64) string {
	sign := ""

	// Min float64 can't be negated to a usable value, so it has to be special cased.
	if v == math.MinInt64 {
		return "-9,223,372,036,854,775,808"
	}

	if v < 0 {
		sign = "-"
		v = 0 - v
	}

	parts := []string{"", "", "", "", "", "", ""}
	j := len(parts) - 1

	for v > 999 {
		parts[j] = fmt.Sprintf("%.0f", math.Floor(math.Mod(v, 1000)))
		switch len(parts[j]) {
		case 2:
			parts[j] = "0" + parts[j]
		case 1:
			parts[j] = "00" + parts[j]
		}
		v = v / 1000
		j--
	}
	parts[j] = strconv.Itoa(int(v))
	return sign + strings.Join(parts[j:], ",")
}

func AbsoluteInteger(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

// NewNullFloat64 function to handle nullable value in SQL
// Params:
// s: float64 to be checked
// Returns sql.NullFloat64
func NewNullFloat64(s *float64) sql.NullFloat64 {
	if s == nil {
		return sql.NullFloat64{}
	}
	if *s == 0 {
		return sql.NullFloat64{}
	}
	return sql.NullFloat64{
		Float64: *s,
		Valid:   true,
	}
}

// NewNullInt32 function to handle nullable value in SQL
// Params:
// s: int32 to be checked
// Returns sql.NullInt32
func NewNullInt32(s int32) sql.NullInt32 {
	if s == 0 {
		return sql.NullInt32{}
	}
	return sql.NullInt32{
		Int32: s,
		Valid: true,
	}
}

// NewNullInt64 function to handle nullable value in SQL
// Params:
// s: int64 to be checked
// Returns sql.NullInt64
func NewNullInt64(s int64) sql.NullInt64 {
	if s == 0 {
		return sql.NullInt64{}
	}
	return sql.NullInt64{
		Int64: s,
		Valid: true,
	}
}

// RandomInteger Generate Random Integer
func RandomInteger(n int) int {
	rand.Seed(time.Now().Unix())
	b := make([]rune, n)
	for i := range b {
		b[i] = numberRunes[rand.Intn(len(numberRunes))]
	}
	res, _ := strconv.Atoi(string(b))
	return res
}
