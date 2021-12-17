package authentication

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWT interface {
	GenerateToken(email string, isUser bool) string
	ValidateToken(token string) (*jwt.Token, error)
}

type authCustomClaims struct {
	Email string `json:"email"`
	ID    string `json:"id"`
	User  bool   `json:"user"`
	jwt.StandardClaims
}

func getSecretKey() string {
	//secret := os.Getenv("SECRET")
	// if secret == "" {
	// 	secret = "secret"
	// } //here we get secret from env but currently i not read env.
	return "secret"
}
func GenerateToken(email string, id string, isUser bool) string {
	claims := &authCustomClaims{
		email,
		id,
		isUser,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    id,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	t, err := token.SignedString([]byte(getSecretKey()))
	if err != nil {
		panic(err)
	}
	return t
}

func ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])

		}
		return []byte(getSecretKey()), nil
	})
}
