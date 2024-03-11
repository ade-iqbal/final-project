package helper

import (
	"errors"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func GenerateToken(id uint, email string) string {
	claims := jwt.MapClaims{
		"id": id,
		"email": email,
	}
	_ = godotenv.Load(".env")

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	signedToken, _ := parseToken.SignedString([]byte(os.Getenv("SECRET_KEY")))
	return signedToken
}

func VerifyToken(ctx *gin.Context) (interface{}, error) {
	errResponse := errors.New("sign in to proceed")
	headerToken := GetToken(ctx)
	bearer := strings.HasPrefix(headerToken, "Bearer")
	_ = godotenv.Load(".env")

	if !bearer {
		return nil, errResponse
	}

	stringToken := strings.Split(headerToken, " ")[1]
	token, _ := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errResponse
		}

		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, errResponse
	}

	return token.Claims.(jwt.MapClaims), nil
}