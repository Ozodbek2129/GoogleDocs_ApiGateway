package auth

import (
	"github.com/dgrijalva/jwt-go"
)

const (
	SIGNING_KEY = "GOoGLe_DoCs"
)

func ValidateAccessToken(tokenStr string) (bool, error) {
	_, err := ExtractAccessClaim(tokenStr)
	if err != nil {
		return false, err
	}
	return true, nil
}

func ExtractAccessClaim(tokenStr string) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(SIGNING_KEY), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		return nil, err
	}

	return &claims, nil
}

func GetUserInfoFromAccessToken(accessTokenString string) (string, string, string, string, error) {
	refreshToken, err := jwt.Parse(accessTokenString, func(token *jwt.Token) (interface{}, error) { return []byte(SIGNING_KEY), nil })
	if err != nil || !refreshToken.Valid {
		return "", "", "", "", err
	}
	claims, ok := refreshToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", "", "", "", err
	}
	userID := claims["user_id"].(string)
	email := claims["email"].(string)
	password := claims["password"].(string)
	role := claims["role"].(string)

	return userID, email, password, role , nil
}
