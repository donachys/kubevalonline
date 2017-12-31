package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/donachys/kubevalonline/router"
	"github.com/garethr/kubeval/kubeval"
	"github.com/gorilla/mux"
)

func Handler() *mux.Router {
	r := router.API()
	r.Get(router.Validate).HandlerFunc(val)
	return r
}

func val(rw http.ResponseWriter, req *http.Request) {
	success := true
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(fmt.Sprint("500 - internal server error")))
	}
	results, err := kubeval.Validate(body, "")
	if err != nil {
		log.Println(err)
		er := []string{err.Error()}
		re := []ResponseElement{ResponseElement{Errors: er}}
		rr := ResultsResponse{Results: re}
		resp, err := json.Marshal(rr)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte(fmt.Sprint("500 - internal server error")))
		} else {
			rw.Header().Set("Content-Type", "application/json; charset=utf-8")
			rw.Write(resp)
		}
	} else {
		success = logResults(results, success)
		rr := ResultsResponse{Results: buildResultsResponse(results, success)}
		logResultResponse(rr)
		resp, err := json.Marshal(rr)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte(fmt.Sprint("500 - internal server error")))
		} else {
			rw.Header().Set("Content-Type", "application/json; charset=utf-8")
			rw.Write(resp)
		}
	}
}

type ResponseElement struct {
	Kind    string
	Success bool
	Errors  []string
}
type ResultsResponse struct {
	Results []ResponseElement
}

func buildResultsResponse(results []kubeval.ValidationResult, success bool) []ResponseElement {
	r := len(results)
	resp := make([]ResponseElement, r, r)
	for i, result := range results {
		re := ResponseElement{Kind: result.Kind}
		n := len(result.Errors)
		if n > 0 {
			re.Success = false
			errs := make([]string, n, n)
			for j, resErr := range result.Errors {
				errs[j] = resErr.Description()
			}
			re.Errors = errs
		} else {
			re.Success = true
		}

		resp[i] = re
	}
	return resp
}
func logResults(results []kubeval.ValidationResult, success bool) bool {
	for _, result := range results {
		if len(result.Errors) > 0 {
			success = false
			log.Println("The document", result.FileName, "contains an invalid", result.Kind)
			for _, desc := range result.Errors {
				log.Println("--->", desc)
			}
		} else {
			log.Println("The document", result.FileName, "contains a valid", result.Kind)
		}
	}
	return success
}

func logResultResponse(r ResultsResponse) {
	for _, result := range r.Results {
		log.Println(result)
	}
}
