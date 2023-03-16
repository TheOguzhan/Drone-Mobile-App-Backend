package utils

import (
	"encoding/json"
	"os"
	"time"

	"github.com/cristalhq/jwt/v5"
	"github.com/google/uuid"
)

type userClaims struct {
	jwt.RegisteredClaims
	User_ID string `json:"user_id,omitempty"`
}

func VerifyToken(token string) (bool, *userClaims, error) {
	key := []byte(os.Getenv("JWT_SECRET"))
	verifier, verifier_err := jwt.NewVerifierHS(jwt.HS256, key)
	if verifier_err != nil {
		return false, nil, verifier_err
	}
	newToken, err := jwt.Parse([]byte(token), verifier)
	if err != nil {
		return false, nil, err
	}

	var newClaims userClaims

	errClaims := json.Unmarshal(newToken.Claims(), &newClaims)

	if errClaims != nil {
		return false, nil, errClaims
	}

	return true, &newClaims, nil

}

func CreateJWTToken(user_id string) (*jwt.Token, error) {
	key := []byte(os.Getenv("JWT_SECRET"))
	signer, err := jwt.NewSignerHS(jwt.HS256, key)

	if err != nil {
		return nil, err
	}

	id, _ := uuid.NewUUID()

	claims := &userClaims{

		RegisteredClaims: jwt.RegisteredClaims{
			ID:        id.String(),
			Audience:  []string{"user"},
			Issuer:    "app-idea",
			Subject:   "",
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(time.Hour * 24)},
			IssuedAt:  &jwt.NumericDate{Time: time.Now()},
		},
		User_ID: user_id,
	}
	builder := jwt.NewBuilder(signer)

	got_token, err := builder.Build(claims)

	if err != nil {
		return nil, err
	}
	return got_token, nil
}
