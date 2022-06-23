package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"rgb/internal/conf"
	"rgb/internal/store"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func testSetup() *gin.Engine {
	store.ResetTestDatabase()
	cfg := conf.NewTestConfig()
	JwtSetup(cfg)
	return SetRouter(cfg)
}

func userJSON(user store.User) string {
	body, err := json.Marshal(map[string]interface{}{
		"Username": user.Username,
		"Password": user.Password,
	})

	if err != nil {
		log.Panic().Err(err).Msg("Error Marshaling JSON body")
	}
	return string(body)
}

func jsonRes(body *bytes.Buffer) map[string]interface{} {
	jsonValue := &map[string]interface{}{}
	err := json.Unmarshal(body.Bytes(), jsonValue)

	if err != nil {
		log.Panic().Err(err).Msg("Error Unmarshalling json value")
	}
	return *jsonValue
}

func performRequest(router *gin.Engine, method, path, body string) *httptest.ResponseRecorder {
	req, err := http.NewRequest(method, path, strings.NewReader(body))
	if err != nil {
		log.Panic().Err(err).Msg("Error creating new request")
	}

	rec := httptest.NewRecorder()
	req.Header.Add("Content-Type", "application/json")
	router.ServeHTTP(rec, req)
	return rec
}
