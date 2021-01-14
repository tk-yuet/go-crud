package encrypt

import (
	"fmt"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Should be in .env
var secret = []byte("secret")

type JwtToken struct {
	Token string `json:"token"`
}

func NewJwtToken(tokenStr string) *JwtToken {
	newInstance := JwtToken{
		Token: tokenStr,
	}
	return &newInstance
}

func GetUserIdByJwtToken(tokenString string) int64 {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return secret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userIdStr := fmt.Sprintf("%v", claims["userId"])
		uid, err := strconv.ParseInt(userIdStr, 10, 64)
		if err != nil {
			return 0
		}
		return uid

	} else {
		fmt.Println(err)
	}
	return 0
}

func GenerateNewToken(userId int) string {
	expirationTime := time.Now().Add(60 * time.Minute)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": "1",
		"StandardClaims": jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	})

	tokenString, _ := token.SignedString(secret)
	return tokenString
}
