package middleware

import (
	"github.com/dgrijalva/jwt-go"
)

type JWT struct {
	JWTKey []byte
}

func NewJWT(jwtKey string) *JWT {
	return &JWT{
		JWTKey: []byte(jwtKey),
	}
}

func (j *JWT) CreateToken(claims Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.JWTKey)
}

func (j *JWT) ValidateJWT(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return j.JWTKey, nil
	})
	if token != nil {
		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			return claims, nil
		}
	}
	return nil, err
}

type Claims struct {
	UnionID string `json:"union_id"`
	jwt.StandardClaims
}

func NewClaims(unionID string) Claims {
	return Claims{
		UnionID:        unionID,
		StandardClaims: jwt.StandardClaims{},
	}
}