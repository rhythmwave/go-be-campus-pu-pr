package objects

import "time"

type LoginRequest struct {
	Username string
	Password string
}

type LoginResponse struct {
	AccessToken     string
	RefreshToken    string
	AppType         string
	ExpiryTime      time.Time
	PermissionNames []string
	Name            string
	Username        string
	AdminRoleName   string
}

type UpdatePasswordRequest struct {
	OldPassword string
	NewPassword string
}

type GetAuthentication struct {
	Id                 string
	Username           string
	AuthenticationType string
	AdminId            *string
	AdminName          *string
	LecturerId         *string
	LecturerName       *string
	StudentId          *string
	StudentName        *string
	IsActive           bool
	SuspensionRemarks  *string
}

type CreateAuthentication struct {
	AuthenticationType string
	UserId             string
}

type BulkCreateAuthentication struct {
	AuthenticationType string
	UserIds            []string
}

type CreateAuthenticationResponse struct {
	Username string
	Password string
}

type BulkCreateAuthenticationResponse struct {
	UserId   string
	Name     string
	Username string
	Password string
}

type UpdateAuthentication struct {
	Id                string
	IsActive          bool
	SuspensionRemarks string
}

type GetSsoAuthResponse struct {
	Url         string
	AppId       string
	FrontendUrl string
}
