package service

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"strconv"
	"time"
)

type JWTService interface {
	GenerateToken(email string) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type AuthCustomClaims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

type jwtServices struct {
	secretKey string
	issure    string
}

func getSecretKey() string {
	secret := viper.GetString("jwt.secret")
	if secret == "" {
		secret = "default_secret"
	}
	return secret
}

func getIssuer() string {
	secret := viper.GetString("jwt.issuer")
	if secret == "" {
		secret = "backend-api"
	}
	return secret
}

func NewJWTService() JWTService {
	return &jwtServices{
		secretKey: getSecretKey(),
		issure:    getIssuer(),
	}
}

func (j *jwtServices) GenerateToken(userID string) (string, error) {
	configExpired, err := strconv.ParseInt(viper.GetString("jwt.expired"), 10, 64)
	if err != nil {
		configExpired = 24
	}
	expiredTime := time.Now().Add(time.Duration(configExpired) * time.Hour)
	claims := &AuthCustomClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiredTime),
			Issuer:    j.issure,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	t, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		return "", err
	}
	return t, nil
}

func (j *jwtServices) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("invalid token - %s", token.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})
}
