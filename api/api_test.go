package api

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"

	"github.com/gorilla/mux"
)

var (
	m *mux.Router
)

func setupAPITest() {
	m = mux.NewRouter()
	// m.Handle("/api/", http.StripPrefix("/api", Handler()))
	m.Handle("/api/val", http.StripPrefix("/api", Handler()))
}

func TestPostOKResponse(t *testing.T) {
	setupAPITest()
	testReq := httptest.NewRequest("POST", "/api/val", nil)
	rec := httptest.NewRecorder()
	m.ServeHTTP(rec, testReq)
	resp := rec.Result()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Should return 200 OK status, received %v", resp.StatusCode)
	}
}

func TestGetNotAllowedResponse(t *testing.T) {
	setupAPITest()
	testReq := httptest.NewRequest("GET", "/api/val", nil)
	rec := httptest.NewRecorder()
	m.ServeHTTP(rec, testReq)
	resp := rec.Result()
	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Fatalf("Should return 405 Method Not Allowed status, received %v", resp.StatusCode)
	}
}

func TestValidateValidInputs(t *testing.T) {
	setupAPITest()
	var tests = []string{
		"valid.yaml",
		"valid.json",
		"multi_valid.yaml",
		"int_or_string.yaml",
		"null_array.yaml",
		"quantity.yaml",
		"extra_property.yaml",
	}
	for _, test := range tests {
		filePath, _ := filepath.Abs("../fixtures/" + test)
		payload, _ := ioutil.ReadFile(filePath)
		testReq := httptest.NewRequest("POST", "/api/val", bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()
		m.ServeHTTP(rec, testReq)
		resp := rec.Result()
		if resp.StatusCode != http.StatusOK {
			t.Errorf("Response expected was 200; for input '"+test+"' response was %v", resp.StatusCode)
		}
	}
}
