package jwt

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// Constantes
const (
	ComponentName = "jwt"
	componentDir  = "server/components/" + ComponentName
)

//Claims es estructura para token
type Claims struct {
	Username string `json:"username"`
	// recommended having
	jwt.StandardClaims
}

//CreateToken para crear token
func CreateToken(secretKey string, username string) (string, error) {

	// Expires the token in 7 days
	expireToken := time.Now().Add(time.Hour * 168).Unix()

	// Create the token using your claims
	token := jwt.New(jwt.SigningMethodHS256)

	//Create a map to store our claims
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = expireToken

	// Signs the token with a secret.
	signedToken, err := token.SignedString([]byte(secretKey))

	if err != nil {
		return "", err
	}

	return signedToken, nil
}

//IsJWTValid Función para comprobar si el token es valido o no
func IsJWTValid(tokenValue string, secretKey string) bool {
	token, err := jwt.Parse(tokenValue, func(token *jwt.Token) (interface{}, error) {
		// Make sure token's signature wasn't changed
		if _, err := token.Method.(*jwt.SigningMethodHMAC); !err {
			return nil, fmt.Errorf("Unexpected siging method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return false
	}

	// Grab the tokens claims and pass it into the original request
	if token.Valid {
		return true
	}
	return false

}
