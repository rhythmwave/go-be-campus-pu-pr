package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"
	"time"

	firebase "firebase.google.com/go"
	"github.com/go-redis/redis/v8"
	"google.golang.org/api/option"
	"google.golang.org/grpc/codes"

	"github.com/sccicitb/pupr-backend/config"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/objects/jwt_object"
	"github.com/sccicitb/pupr-backend/utils"

	"github.com/dgrijalva/jwt-go"
)

// JWTInterface interface jwt
type JWTInterface interface {
	ExtractJWTClaims(ctx context.Context, token string, appName string) (claims *jwt_object.JWTClaims, errs *constants.ErrorResponse)
	ValidateTokenIssuer(claims *jwt_object.JWTClaims) *constants.ErrorResponse
	ValidateTokenExpire(ctx context.Context, claims *jwt_object.JWTClaims, reqToken string, purpose string, appName string, device string) *constants.ErrorResponse

	DeleteTokenFromRedis(ctx context.Context, claims *jwt_object.JWTClaims, purpose string, appName string) *constants.ErrorResponse
	GenerateJWTToken(ctx context.Context, request jwt_object.JWTRequest, expireTime time.Duration, appName string, purpose string, device string) (string, *constants.ErrorResponse)
	GenerateJWTFirestoreCustomToken(ctx context.Context, userID string) (string, *constants.ErrorResponse)
}

const (
	// AuthKeyAccessToken auth key for redis
	AuthKeyAccessToken = "ametory-auth-access-token"
	// AuthKeyRefreshToken auth key for redis
	AuthKeyRefreshToken = "ametory-auth-refresh-token"
	// AdminAuthKeyAccessToken admin auth key for redis
	AdminAuthKeyAccessToken = "ametory-admin-auth-access-token"
	// AdminAuthKeyRefreshToken admin auth key for redis
	AdminAuthKeyRefreshToken = "ametory-admin-auth-refresh-token"
)

// jwtObj struct
type jwtObj struct {
	config             *config.JWTConfig
	redis              *redis.Client
	firebase           *config.Firebase
	allowMultipleLogin bool
}

// NewJWT function to connect jwtObj to JWTInterface
// Params:
// cfg: config
// redis: redis client
// Returns JWTInterface
func NewJWT(cfg *config.JWTConfig, redis *redis.Client, firebase *config.Firebase, allowMultipleLogin bool) JWTInterface {
	return &jwtObj{
		config:             cfg,
		redis:              redis,
		firebase:           firebase,
		allowMultipleLogin: allowMultipleLogin,
	}
}

// ExtractJWTClaims function to extract jwt claims from authorization header
// Params:
// ctx: context
// token: token to extract from
func (j *jwtObj) ExtractJWTClaims(ctx context.Context, token string, appName string) (claims *jwt_object.JWTClaims, errs *constants.ErrorResponse) {
	// check authorization
	splitToken := strings.Split(token, constants.Bearer)
	if len(splitToken) != 2 {
		return nil, constants.ErrTokenIsRequired
	}
	reqToken := strings.TrimSpace(splitToken[1])

	t, err := jwt.ParseWithClaims(reqToken, &jwt_object.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.config.Secret), nil
	})
	if err != nil {
		if err.Error() == constants.JWTInvalidType {
			return nil, constants.ErrKeyIsNotInvalidType
		}
		return nil, constants.Error(http.StatusInternalServerError, codes.Internal, constants.DefaultCustomErrorCode, err.Error(), token)
	}

	claims = t.Claims.(*jwt_object.JWTClaims)
	// Validate Issuer Token
	errs = j.ValidateTokenIssuer(claims)
	if errs != nil {
		return nil, errs
	}

	// Validate token expire
	errs = j.ValidateTokenExpire(ctx, claims, reqToken, claims.TokenType, appName, claims.Device)
	if errs != nil {
		return nil, errs
	}
	return claims, nil
}

// ValidateTokenIssuer is for validate token issuer
func (j *jwtObj) ValidateTokenIssuer(claims *jwt_object.JWTClaims) *constants.ErrorResponse {
	if claims.Issuer != j.config.Issuer {
		return constants.ErrTokenInvalid
	}
	return nil
}

// ValidateTokenExpire is for validate Token Expire
func (j *jwtObj) ValidateTokenExpire(ctx context.Context, claims *jwt_object.JWTClaims, reqToken string, purpose string, appName string, device string) *constants.ErrorResponse {
	redisKey := utils.GenerateRedisKey(appName, purpose, claims.ID, device, claims.UniqueKey, j.allowMultipleLogin)
	// check token to redis
	token, err := j.getTokenFromRedis(ctx, redisKey)
	if err != nil {
		return utils.ErrGetTokenFromRedis(err.Error())
	}
	if token != reqToken {
		return constants.ErrTokenReplaced
	}

	return nil
}

func (j *jwtObj) getTokenFromRedis(ctx context.Context, key string) (string, error) {
	val, err := j.redis.Get(ctx, key).Result()
	if err != nil {
		return val, err
	}
	return val, nil
}

// DeleteTokenFromRedis function to delete token from redis
// Params:
// ctx: context
// id: user ID / admin ID
// authKey: redis authorization key
// Returns *constants.ErrorResponse
func (j *jwtObj) DeleteTokenFromRedis(ctx context.Context, claims *jwt_object.JWTClaims, purpose string, appName string) *constants.ErrorResponse {
	redisKey := utils.GenerateRedisKey(appName, purpose, claims.ID, claims.Device, claims.UniqueKey, j.allowMultipleLogin)
	_, err := j.redis.Del(ctx, redisKey).Result()
	if err != nil {
		return constants.Error(http.StatusInternalServerError, codes.Internal, constants.DefaultCustomErrorCode, err.Error(), redisKey)
	}

	return nil
}

// Generate Token
// Params:
// ctx: context
// id: user ID / admin ID
// authKey: redis authorization key
// Returns *constants.ErrorResponse
func (j *jwtObj) GenerateJWTToken(ctx context.Context, request jwt_object.JWTRequest, expireTime time.Duration, appName string, purpose string, device string) (string, *constants.ErrorResponse) {
	JWTSignatureKey := []byte(j.config.Secret)
	claims := jwt_object.JWTClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    j.config.Issuer,
			ExpiresAt: time.Now().Add(time.Duration(expireTime)).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		ID:          request.ID,
		Name:        request.Name,
		Email:       request.Email,
		FlutterUdid: request.FlutterUdid,
		IsAdmin:     request.IsAdmin,
		Role:        request.Role,
		TokenType:   purpose,
		Device:      device,
		UniqueKey:   utils.RandomString(6),
		Permissions: request.Permissions,
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	// create token client
	signedToken, err := token.SignedString(JWTSignatureKey)
	if err != nil {
		return "", constants.ErrGenerateToken
	}

	redisKey := utils.GenerateRedisKey(appName, purpose, claims.ID, device, claims.UniqueKey, j.allowMultipleLogin)
	errs := j.setTokenToRedis(ctx, signedToken, expireTime, redisKey)
	if errs != nil {
		return "", constants.ErrGenerateToken
	}

	return signedToken, nil
}

func (j *jwtObj) GenerateJWTFirestoreCustomToken(ctx context.Context, userID string) (string, *constants.ErrorResponse) {
	opt := option.WithCredentialsFile(j.firebase.KeyFileDir)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return "", constants.ErrGenerateFirestoreCustomToken
	}

	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
		return "", constants.ErrGenerateFirestoreCustomToken
	}

	signedToken, err := client.CustomToken(ctx, userID)
	if err != nil {
		log.Fatalf("error minting custom token: %v\n", err)
		return "", constants.ErrGenerateFirestoreCustomToken
	}

	return signedToken, nil
}

func (j *jwtObj) setTokenToRedis(ctx context.Context, token string, expireTime time.Duration, key string) *constants.ErrorResponse {
	_, err := j.redis.Set(ctx, key, token, expireTime).Result()
	if err != nil {
		return constants.ErrSetTokenToRedis
	}

	return nil
}
