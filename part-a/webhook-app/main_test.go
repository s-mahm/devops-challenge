package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var a App

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code is %d. Got %d", expected, actual)
	}
}

func TestMain(m *testing.M) {
	a = App{}
	a.Initialize(os.Getenv("RC_URL"))

	code := m.Run()

	os.Exit(code)
}

func TestSendCommit(t *testing.T) {
	req, _ := http.NewRequest("POST", "/commit/8731ba2cdd98f8c7753481afa08d5148ca6dec21", nil)
	response := executeRequest(req)
	print(req.URL.String())

	checkResponseCode(t, http.StatusOK, response.Code)
	// responseLen := utf8.RuneCountInString(response.Body.String())
	// if responseLen != 40 {
	// 	t.Errorf("Expected hash of 40 characters, got %v", responseLen)
	// }
}
