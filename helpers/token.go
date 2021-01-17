package helpers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"tap-talk/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

// Create JWT Token
func CreateToken(username string) (*models.TokenDetails, error) {
	var err error
	dur, _ := strconv.Atoi(os.Getenv("TOKEN_EXP"))
	exp := time.Second * time.Duration(dur)

	token := &models.TokenDetails{}
	token.Username = username
	token.Exp = time.Now().Add(exp).Unix()
	token.UUID = uuid.New().String()

	atClaims := jwt.MapClaims{}
	atClaims["username"] = token.Username
	atClaims["uuid"] = token.UUID
	atClaims["exp"] = token.Exp
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token.Token, err = at.SignedString([]byte(os.Getenv("TOKEN_SECRET")))

	if err != nil {
		return nil, err
	}
	return token, nil
}

// Verify token header
func VerifyToken(r *http.Request) (*jwt.Token, error) {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	tokenString := ""
	if len(strArr) == 2 {
		tokenString = strArr[1]
	}
	jwtToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return jwtToken, nil
}

// Extract JWT token
func ExtractToken(r *http.Request) (*models.TokenDetails, error) {
	token := &models.TokenDetails{}
	jwtToken, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if ok && jwtToken.Valid {
		token.UUID, ok = claims["uuid"].(string)
		if !ok {
			return nil, err
		}
		token.Username, ok = claims["username"].(string)
		if !ok {
			return nil, err
		}
		return token, nil
	}
	return nil, err
}
