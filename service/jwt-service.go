package service

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

//jwt service
type JWTService interface {
	GenerateToken(email string) (*string,error)
	ValidateToken(context.Context, string) (*jwt.Token, error)
}
type authCustomClaims struct {
	ID string `json:"id"`
	//	Uuu   bool   `json:"uuu"`
	jwt.StandardClaims
}

type jwtServices struct {
	secretKey string
	//issure    string
}

//auth-jwt
func JWTAuthService() JWTService {
	return &jwtServices{
		secretKey: "secretkey",
		//	issure:    "Abenezer",
	}
}

// func getSecretKey() string {
// 	secret := os.Getenv("KEY")
// 	if secret == "" {
// 		secret = "mysecretkey"
// 	}
// 	return secret
// }

func (service *jwtServices) GenerateToken(ID string) (*string,error) {
	claims := &authCustomClaims{
		ID, 
		//	uuu,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    ID,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	t, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		
		return nil,err
	}

	return &t,nil
}

func (service *jwtServices) ValidateToken(ctx context.Context, encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("invalid token %v", token.Header["alg"])

		}
		return []byte(service.secretKey), nil
	})

}
