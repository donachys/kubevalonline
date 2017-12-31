package app

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

var (
	m *mux.Router
)

func setupAppTest() {
	m = mux.NewRouter()
	m.Handle("/", Handler())
}

func TestIndexValidResponse(t *testing.T) {
	setupAppTest()
	testReq := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()
	m.ServeHTTP(rec, testReq)
	resp := rec.Result()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Should return 200 OK status, received %v", resp.StatusCode)
	}
}
