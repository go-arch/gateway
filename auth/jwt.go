package auth

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"Gateway/constants"
	"time"

)

var (
	verifyKey *rsa.PublicKey
	signKey *rsa.PrivateKey
)

func InitiateJwt() {

	signKeyPath, err := filepath.Abs(constants.SIGN_KEY_PATH)
	signBytes, err := ioutil.ReadFile(signKeyPath)
	Fatal(err)

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	Fatal(err)
	verifyKeyPath, err := filepath.Abs(constants.PUB_KEY_PATH)

	verifyByte, err := ioutil.ReadFile(verifyKeyPath)
	Fatal(err)

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyByte)
	Fatal(err)

}

func ValidateMiddleware(w http.ResponseWriter, r *http.Request) map[string]interface{} {

	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor, func(token *jwt.Token) (interface{}, error) {
		return verifyKey, nil
	})
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Token is Not valid")
		return map[string]interface{}{}
	} else if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(token.Claims)

		return claims
		//next(w, r.WithContext(ctx))
		//token.Claims
	} else {
		w.WriteHeader(http.StatusUnauthorized) //TODO code repeated
		fmt.Fprint(w, "Token is Not valid")
		return map[string]interface{}{}

	}
	return map[string]interface{}{}
}

func GenerateJwtToken(Id User) (string, error) {
	token := jwt.New(jwt.SigningMethodRS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["id"] = Id.ID
	claims["role"] = Id.Role
	token.Claims = claims
	return token.SignedString(signKey)
}

func Fatal(err error) {
	if err != nil {
		fmt.Println(err)
	}
}


