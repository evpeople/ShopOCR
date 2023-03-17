package logic

import "github.com/form3tech-oss/jwt-go"

// Returns a JWT token using the given secret key, the given issue time, the given
// number of seconds the token will be valid for, and the given user id.
func getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims

	return token.SignedString([]byte(secretKey))
}
