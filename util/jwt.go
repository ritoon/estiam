package util

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

var mySigningKey = []byte("AllYourBase")

type MyCustomClaims struct {
	UserUUID  string `json:"user_uuid"`
	FirstName string
	jwt.StandardClaims
}

func CreateJWT(signKey []byte, userUUID, firstName string) (string, error) {

	// Create the Claims
	claims := MyCustomClaims{
		userUUID,
		firstName,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
			Issuer:    "Estiam",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", ss), nil
}
