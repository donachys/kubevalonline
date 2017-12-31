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

func TestIndexOKResponse(t *testing.T) {
	setupAppTest()
	testReq := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()
	m.ServeHTTP(rec, testReq)
	resp := rec.Result()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Should return 200 OK status, received %v", resp.StatusCode)
	}
}

func TestBadRouteNotFoundResponse(t *testing.T) {
	setupAppTest()
	testReq := httptest.NewRequest("GET", "/path/to/route", nil)
	rec := httptest.NewRecorder()
	m.ServeHTTP(rec, testReq)
	resp := rec.Result()
	if resp.StatusCode != http.StatusNotFound {
		t.Fatalf("Should return 404 Not Found status, received %v", resp.StatusCode)
	}
}
