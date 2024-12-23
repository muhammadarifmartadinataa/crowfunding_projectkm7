package auth

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt"
)

type Service interface {
	GenerateToken(userID int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
}

//var SECRET_KEY = []byte("CROWFUND1NG_MIN1PROJECT_S3cR3T_k3y")

func NewService() *jwtService {
	return &jwtService{}
}

// Generate Token
func (s *jwtService) GenerateToken(userID int) (string, error) {
	//payload
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return signedToken, err
	}
	return signedToken, nil

}

// Validasi Token
func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Invalid Token")
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil {
		return token, err
	}
	return token, nil
}
