package service

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTService interface {
	GenerateToken(userID string) string
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtCustomClaim struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}


//NewJWTService create a new instance of JWTService
func NewJWTService() JWTService {
	return &jwtService{
		issuer: "yntkts",
		secretKey: getSecretKey(),
	}	
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	return secretKey
}


func (j *jwtService) GenerateToken(userID string) string {
	claims := &jwtCustomClaim{
		userID,
		jwt.StandardClaims{
			//token kadaluarsa setelah 59 minute
			ExpiresAt: time.Now().Add(time.Minute * 59).Unix(),
			Issuer: j.issuer,
			IssuedAt: time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	t, err := token.SignedString([]byte(j.secretKey))

	if err != nil {
		panic(err.Error())
	}

	return t
}


func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t.Header["alg"])
		}
	
		return []byte(j.secretKey), nil
	})
}


