package otp

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/sccicitb/pupr-backend/constants"
)

// OTPInterface interface for otp service
type OTPInterface interface {
	GenerateOTP(ctx context.Context, length int, hour int, id string, purpose string) (string, *constants.ErrorResponse)
	SaveOTPToRedis(ctx context.Context, id string, purpose string, hour int, otp string) *constants.ErrorResponse
	GetOTPFromRedis(ctx context.Context, id string, purpose string) (OTPData, *constants.ErrorResponse)
	DeleteOTPFromRedis(ctx context.Context, id string, purpose string) *constants.ErrorResponse
}

// NewOTP function to connect mail to OTPInterface
// Params:
// redis: redis client
// Returns OTPInterface
func NewOTP(redis *redis.Client, appName string) OTPInterface {
	return &otpObj{
		redis:   redis,
		appName: appName,
	}
}
