package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"rgb/internal/conf"
	"rgb/internal/store"
	"strconv"
	"time"

	"github.com/cristalhq/jwt/v4"
	"github.com/rs/zerolog/log"
)

var (
	jwtSigner   jwt.Signer
	jwtVerifier jwt.Verifier
)

func JwtSetup(conf conf.Config) {
	var err error
	key := []byte(conf.JwtSecret)

	jwtSigner, err = jwt.NewSignerHS(jwt.HS256, key)
	if err != nil {
		log.Panic().Err(err).Msg("Error creating JWT signer")
	}

	jwtVerifier, err = jwt.NewVerifierHS(jwt.HS256, key)
	if err != nil {
		log.Panic().Err(err).Msg("Error creating JWT verifier")
	}
}

func GenerateJWT(user *store.User) string {
	claims := &jwt.RegisteredClaims{
		ID:        fmt.Sprint(user.ID),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 5)),
	}

	builder := jwt.NewBuilder(jwtSigner)
	token, err := builder.Build(claims)
	if err != nil {
		log.Panic().Err(err).Msg("Error Buidling JWT")
	}
	return token.String()
}

func VerifyJWT(tokenStr string) (int, error) {
	tokenBytes := []byte(tokenStr)
	token, err := jwt.Parse(tokenBytes, jwtVerifier)
	if err != nil {
		log.Error().Err(err).Str("tokenStr", tokenStr).Msg("Error Parsing JWT")
		return 0, err
	}

	if err = jwtVerifier.Verify(token); err != nil {
		log.Error().Err(err).Msg("Error verifying token")
		return 0, err
	}

	var newClaims jwt.RegisteredClaims
	if err = json.Unmarshal(token.Claims(), &newClaims); err != nil {
		log.Error().Err(err).Msg("Error Unmarshalling the claims")
		return 0, err
	}

	if err = jwt.ParseClaims(tokenBytes, jwtVerifier, &newClaims); err != nil {
		log.Error().Err(err).Msg("Error parsing the the claims")
		return 0, err
	}

	if notExpired := newClaims.IsValidAt(time.Now()); !notExpired {
		return 0, errors.New("Token expired.")
	}

	id, err := strconv.Atoi(newClaims.ID)
	if err != nil {
		log.Error().Err(err).Str("claims.ID", newClaims.ID).Msg("Error converting claims ID to number")
		return 0, errors.New("ID in token is not valid")
	}
	return id, err
}
