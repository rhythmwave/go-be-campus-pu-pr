package otp

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/sccicitb/pupr-backend/constants"
	"google.golang.org/grpc/codes"

	"github.com/go-redis/redis/v8"
)

type otpObj struct {
	redis   *redis.Client
	appName string
}

// OTPData base otp struct
type OTPData struct {
	OTP       string    `json:"otp"`
	CreatedAt time.Time `json:"created_at"`
}

func (j *otpObj) GenerateOTP(ctx context.Context, length int, hour int, id string, purpose string) (string, *constants.ErrorResponse) {
	var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	otpByte := make([]byte, length)
	n, err := io.ReadAtLeast(rand.Reader, otpByte, length)
	if n != length || err != nil {
		return "", constants.Error(http.StatusInternalServerError, codes.Internal, constants.DefaultCustomErrorCode, err.Error())
	}
	for i := 0; i < len(otpByte); i++ {
		otpByte[i] = table[int(otpByte[i])%len(table)]
	}

	otp := string(otpByte)

	// Save token to redis
	errs := j.SaveOTPToRedis(ctx, id, purpose, hour, otp)
	if errs != nil {
		return "", constants.ErrSaveOTPToRedis
	}

	return otp, nil
}

func (j *otpObj) SaveOTPToRedis(ctx context.Context, id string, purpose string, hour int, otp string) *constants.ErrorResponse {
	now := time.Now()

	otpData := OTPData{
		OTP:       otp,
		CreatedAt: now,
	}
	otpJSON, err := json.Marshal(otpData)
	if err != nil {
		return constants.Error(http.StatusInternalServerError, codes.Internal, constants.DefaultCustomErrorCode, err.Error())
	}

	key := fmt.Sprintf("%s-%s:%s", j.appName, purpose, id)
	ttl := time.Duration(hour) * time.Hour
	err = j.redis.Set(ctx, key, string(otpJSON), ttl).Err()
	if err != nil {
		return constants.Error(http.StatusInternalServerError, codes.Internal, constants.DefaultCustomErrorCode, err.Error())
	}

	return nil
}

func (j *otpObj) GetOTPFromRedis(ctx context.Context, id string, purpose string) (OTPData, *constants.ErrorResponse) {
	var result OTPData

	key := fmt.Sprintf("%s-%s:%s", j.appName, purpose, id)
	val, err := j.redis.Get(ctx, key).Result()
	if err != nil {
		return result, constants.Error(http.StatusInternalServerError, codes.Internal, constants.DefaultCustomErrorCode, err.Error())
	}
	err = json.Unmarshal([]byte(val), &result)
	if err != nil {
		return result, constants.Error(http.StatusInternalServerError, codes.Internal, constants.DefaultCustomErrorCode, err.Error())
	}

	return result, nil
}

func (j *otpObj) DeleteOTPFromRedis(ctx context.Context, id string, purpose string) *constants.ErrorResponse {
	key := fmt.Sprintf("%s-%s:%s", j.appName, purpose, id)
	_, err := j.redis.Del(ctx, key).Result()
	if err != nil {
		return constants.Error(http.StatusInternalServerError, codes.Internal, constants.DefaultCustomErrorCode, err.Error())
	}

	return nil
}
