package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var AccessSecret = "secret_token"
var RefreshSecret = "refresh_token"

func GenerateAccessToken(username string) (string,error) {
	claims := jwt.MapClaims {
		"username":username,
		"exp": time.Now().Add(time.Minute * 15).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	webToken,err := token.SignedString([]byte(AccessSecret))
	if err != nil {
		return "",err
	}
	return webToken,nil
}

func GenerateRefreshToken(username string) (string,error) {
	claims := jwt.MapClaims {
		"username":username,
		"exp": time.Now().Add(time.Minute * 24 * 7).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	webToken,err := token.SignedString([]byte(RefreshSecret))
	if err != nil {
		return "",err
	}
	return webToken,nil
}

func VerifyAccessToken(tokenString string) (*jwt.Token, error){
	token,err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _,ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(AccessSecret),nil
	})
	if err != nil{
		return nil,err
	}
	return token,nil
}

func VerifyRefreshToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(RefreshSecret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}


func DecodeToken(tokenString string)(jwt.MapClaims,error){
	token,err := VerifyAccessToken(tokenString)
	if err != nil {
		return nil,err
	}

	claims,isOK := token.Claims.(jwt.MapClaims)
	if isOK && token.Valid {
		return claims,nil
	}
	return nil,fmt.Errorf("invalid decode token")
}