package jwt_object

import (
	"github.com/dgrijalva/jwt-go"
)

// JWTClaims struct of standard JWT middleware
type JWTClaims struct {
	jwt.StandardClaims
	ID          string
	Email       string
	Name        string
	Role        string
	FlutterUdid string
	Timezone    string
	TokenType   string
	Device      string
	IsAdmin     bool
	CompanyID   string
	EmployeeID  string
	UniqueKey   string
	Permissions []string
}

// JWTRequest struct for request jwt
type JWTRequest struct {
	ID          string   `json:"id"`
	Name        string   `json:"nama"`
	Email       string   `json:"email"`
	Role        string   `json:"role"`
	FlutterUdid string   `json:"flutter_udid"`
	IsAdmin     bool     `json:"is_admin"`
	Timezone    string   `json:"timezone"`
	Permissions []string `json:"permissions"`
}

// JWTSimpleRequest struct for request jwt
type JWTSimpleRequest struct {
	UID    string               `json:"uid"`
	Claims JWTSimpleChildClaims `json:"claims"`
}

// JWTSimpleChildClaims struct for request jwt
type JWTSimpleChildClaims struct {
	UID string `json:"uid"`
	Alg string `json:"alg"`
}

// JWTSimpleClaims struct for claims response jwt
type JWTSimpleClaims struct {
	jwt.StandardClaims
	JWTSimpleRequest
}
