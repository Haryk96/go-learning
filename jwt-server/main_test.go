package main

import (
	"bytes"
	"jwt-server/router"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	router.SetupRouter(r)
	return r
}

func TestEndpoints(t *testing.T) {
	type testCase struct {
		method      string
		url         string
		body        []byte
		expectedRes int
	}
	testCases := []testCase{
		{method: "GET", url: "/ping", body: nil, expectedRes: http.StatusOK},
		{method: "POST", url: "/jwt", body: []byte(`{"name": "John", "password": "nigga-123"}`), expectedRes: http.StatusCreated},
	}

	for _, tc := range testCases {
		r := setupRouter()

		req, err := http.NewRequest(tc.method, tc.url, bytes.NewBuffer(tc.body))
		assert.NoError(t, err)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, tc.expectedRes, w.Code)
	}
}
