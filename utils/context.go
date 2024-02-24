package utils

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/objects/jwt_object"
)

// GetJWTClaimsFromContext function to get jwt claims from context
func GetJWTClaimsFromContext(ctx context.Context) (*jwt_object.JWTClaims, *constants.ErrorResponse) {
	claims, ok := ctx.Value(constants.ClaimsContextKey).(*jwt_object.JWTClaims)
	if !ok {
		return nil, constants.ErrInvalidClaims
	}
	return claims, nil
}
