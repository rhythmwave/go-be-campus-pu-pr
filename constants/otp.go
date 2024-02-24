package constants

import "time"

const (
	// DefaultOTPLength default otp length
	DefaultOTPLength = 6
	// DefaultOTPLifeHour default otp life time (hour)
	DefaultOTPLifeHour = 24
	// DefaultOTPResendTimePeriod period to default value to resend OTP
	DefaultOTPResendTimePeriod = 1 * time.Minute
)

const (
	// OTPUserEmailAccountActivationPurpose user email otp account activation purpose for redis
	OTPUserEmailAccountActivationPurpose = "user-email-otp-account-activation"
	// OTPUserPhoneNumberAccountActivationPurpose user email otp account activation purpose for redis
	OTPUserPhoneNumberAccountActivationPurpose = "user-phone-number-otp-account-activation"
	// OTPUserLoginPurpose login OTP for redis
	OTPUserLoginPurpose = "otp-user-login"
	// OTPUserForgotPassword forgot password OTP for redis
	OTPUserForgotPassword = "otp-user-forgot-password"
)

const (
	// OTPAdminEmailAccountActivationPurpose admin email otp account activation purpose for redis
	OTPAdminEmailAccountActivationPurpose = "admin-email-otp-account-activation"
	// OTPAdminPhoneNumberAccountActivationPurpose admin email otp account activation purpose for redis
	OTPAdminPhoneNumberAccountActivationPurpose = "admin-phone-number-otp-account-activation"
	// OTPAdminLoginPurpose login OTP for redis
	OTPAdminLoginPurpose = "otp-admin-login"
	// OTPAdminForgotPassword forgot password OTP for redis
	OTPAdminForgotPassword = "otp-admin-forgot-password"
)
