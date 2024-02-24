package objects

type SsoFindSession struct {
	ID               string                  `json:"_id"`
	Username         string                  `json:"username"`
	AccessToken      string                  `json:"access_token"`
	ExpiresIn        int64                   `json:"expires_in"`
	RefreshToken     string                  `json:"refresh_token"`
	RefreshExpiresIn int64                   `json:"refresh_expires_in"`
	LocalSession     string                  `json:"local_session"`
	AppAcces         []SsoFindSessionAppAcce `json:"appAcces"`
	ExpireAt         string                  `json:"expireAt"`
	V                int64                   `json:"__v"`
}

type SsoFindSessionAppAcce struct {
	ID        string               `json:"_id"`
	Roles     []SsoFindSessionRole `json:"roles"`
	Class     []interface{}        `json:"class"`
	UserID    string               `json:"userId"`
	AppID     SsoFindSessionAppID  `json:"appId"`
	CreatedAt string               `json:"createdAt"`
	UpdatedAt string               `json:"updatedAt"`
	V         int64                `json:"__v"`
}

type SsoFindSessionAppID struct {
	ID          string        `json:"_id"`
	Name        string        `json:"name"`
	Domain      string        `json:"domain"`
	RedirectURI string        `json:"redirect_uri"`
	UserCount   int64         `json:"userCount"`
	TagID       string        `json:"tagId"`
	ClusterID   []string      `json:"clusterId"`
	IsPublic    bool          `json:"isPublic"`
	OtpSetting  bool          `json:"otpSetting"`
	Thumbnail   string        `json:"thumbnail"`
	Description string        `json:"description"`
	Roles       []string      `json:"roles"`
	Class       []interface{} `json:"class"`
	CreatedAt   string        `json:"createdAt"`
	UpdatedAt   string        `json:"updatedAt"`
	V           int64         `json:"__v"`
}

type SsoFindSessionRole struct {
	ID        string `json:"_id"`
	Name      string `json:"name"`
	Detail    string `json:"detail"`
	AppID     string `json:"appId"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	V         int64  `json:"__v"`
}

type SsoUserInfo struct {
	Status  int64           `json:"status"`
	Message string          `json:"message"`
	Data    SsoUserInfoData `json:"data"`
}

type SsoUserInfoData struct {
	ID         string   `json:"_id"`
	SubjectID  string   `json:"subjectId"`
	Name       string   `json:"name"`
	Username   string   `json:"username"`
	Password   string   `json:"password"`
	Nip        string   `json:"nip"`
	Email      string   `json:"email"`
	IsVerified bool     `json:"isVerified"`
	Phone      int64    `json:"phone"`
	Apps       []string `json:"apps"`
	CreatedAt  string   `json:"createdAt"`
	UpdatedAt  string   `json:"updatedAt"`
	V          int64    `json:"__v"`
}

type SsoLogout struct {
	Status int64  `json:"status"`
	Data   string `json:"data"`
}
