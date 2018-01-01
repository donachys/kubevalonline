package api

import (
	"bytes"
	"encoding/json"
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

func TestOKResponseForValidInputs(t *testing.T) {
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

func TestSuccessResponseForValidInputs(t *testing.T) {
	setupAPITest()
	var tests = []string{
		"valid.yaml",
		"valid.json",
		"multi_valid.yaml",
		"int_or_string.yaml",
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
		} else {
			respBody, _ := ioutil.ReadAll(resp.Body)
			rr := ResultsResponse{}
			json.Unmarshal(respBody, &rr)
			for _, result := range rr.Results {
				if result.Success != true {
					t.Errorf("Response for test \"%v\" included Failure:\n %+v", test, rr)
				}
			}
		}
	}
}

func TestOKResponseForInvalidInputs(t *testing.T) {
	setupAPITest()
	var tests = []string{
		"blank.yaml",
		"missing-kind.json",
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

func TestFailureResponseForInvalidInputs(t *testing.T) {
	setupAPITest()
	var tests = []string{
		"null_array.yaml",
		"blank.yaml",
		"missing-kind.json",
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
		} else {
			respBody, _ := ioutil.ReadAll(resp.Body)
			rr := ResultsResponse{}
			json.Unmarshal(respBody, &rr)
			for _, result := range rr.Results {
				if result.Success != false {
					t.Errorf("Response for test \"%v\" included Success:\n %+v", test, rr)
				}
			}
		}
	}
}
