package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Auth struct {
	Issuer        string
	Audience      string
	Secret        string
	TokenExpiry   time.Duration
	RefreshExpiry time.Duration
	CookieDomain  string
	CookiePath    string
	CookieName    string
}

type jwtUser struct {
	Id        int64  `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type TokenPairs struct {
	Token        string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type Claims struct {
	jwt.RegisteredClaims
}

func (j *Auth) GenerateTokenPair(user *jwtUser) (TokenPairs, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["name"] = fmt.Sprintf("%s %s", user.FirstName, user.LastName)
	claims["sub"] = user.Id
	claims["aud"] = j.Audience
	claims["iss"] = j.Issuer
	claims["iat"] = time.Now().UTC().Unix()
	claims["typ"] = "JWT"

	claims["exp"] = time.Now().UTC().Add(j.TokenExpiry).Unix()

	signedAccessToken, err := token.SignedString([]byte(j.Secret))

	if err != nil {
		return TokenPairs{}, err
	}

	refreshToken := jwt.New(jwt.SigningMethodHS256)

	refreshTokenClaims := refreshToken.Claims.(jwt.MapClaims)
	refreshTokenClaims["sub"] = fmt.Sprint(user.Id)
	refreshTokenClaims["iat"] = time.Now().UTC().Unix()

	refreshTokenClaims["exp"] = time.Now().UTC().Add(j.RefreshExpiry).Unix()

	signedRefreshToken, err := refreshToken.SignedString([]byte(j.Secret))

	if err != nil {
		return TokenPairs{}, err
	}
	return TokenPairs{
		Token:        signedAccessToken,
		RefreshToken: signedRefreshToken,
	}, nil
}
