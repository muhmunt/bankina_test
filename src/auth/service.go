package auth

import (
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

var SECRET_KEY = os.Getenv("SECRET_KEY")

func GenerateToken(userID int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signToken, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		return signToken, err
	}

	return signToken, nil
}

func ValidateToken(signedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Invalid token")
		}

		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}

func ExtractBearerToken(r *http.Request) (string, bool) {
	header := r.Header.Get("Authorization")

	if header == "" || !strings.HasPrefix(header, "Bearer ") {

		return "", false
	}

	token := strings.TrimPrefix(header, "Bearer ")

	return token, true
}

func ValidateLevel(userRole, role string) bool {
	if userRole != role {
		return false
	}

	return true
}
