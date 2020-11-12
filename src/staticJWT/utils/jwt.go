package utils

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"

	jwt "github.com/dgrijalva/jwt-go"
)

// TokenClaims - Claims for a JWT access token.
type TokenClaims struct {
	UserID string `json:"userId"`
	Email  string `json:"email"`
	Token  string `json:"userKey"`
	jwt.StandardClaims
}

func IssueJWTToken(userID string, email, token string) (string, error) {
	//read the key from .env

	if !fileExists(".env") {
		//for test check one directory up
		if fileExists("../.env") {
			err := godotenv.Load("../.env")
			if err != nil {
				log.Println(err)
				log.Fatal("Error loading .env file")
			}
		}
	} else {
		err := godotenv.Load(".env")
		if err != nil {
			log.Println(err)
			log.Fatal("Error loading .env file")
		}
	}

	jwtKey := os.Getenv("APP_SECRET_KEY")
	tokenExpireTime := os.Getenv("TOKEN_EXPIRE_TIME")

	tokenExpireTimeInt, err := strconv.Atoi(tokenExpireTime)

	if err != nil {
		log.Fatal("Failed reading token expire time, invalid number for minute.")
	}

	expirationTime := time.Now().Add(time.Duration(tokenExpireTimeInt) * time.Minute).Unix()

	claims := &TokenClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime,
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := jwtToken.SignedString([]byte(jwtKey))
	return ("Bearer " + tokenString), err
}

func VerifyJWTToken(tknStr string) (*TokenClaims, error) {

	//read the key from .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err)
		log.Fatal("Error loading .env file")
	}
	jwtKey := os.Getenv("APP_SECRET_KEY")

	claims := &TokenClaims{}

	token, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected siging method")
		}
		return []byte(jwtKey), nil
	})
	if !token.Valid {
		return nil, err
	}
	return claims, err
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
