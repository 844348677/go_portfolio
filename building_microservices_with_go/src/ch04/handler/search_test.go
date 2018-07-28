package handler

import (
	"testing"
	"net/http/httptest"
	"net/http"
)

func TestSearchHandlerReturnsBadRequestWhenNoSearchCriterialSent(t *testing.T){
	handler := SearchHandler{}
	request := httptest.NewRequest("GET","/search",nil)
	response := httptest.NewRecorder()

	handler.ServeHTTP(response,request)

	if response.Code != http.StatusBadRequest {
		t.Errorf("expected Bad request go %v ",response.Code)
	}
}

// the net/http/httptest package
// Mocks or Spies

// httptest.NewRequest
// return an incoming server request which we can then pass to our http.Handler
// func NewRequest(method , target string , body io.Reader) *http.Request
// io.Reader file which will correspond to the body of the request

// httptest.NewRecorder
// ResponseWriter type is an implementation of http.ResponseWriter

// Go does not have an assertion library as you would find with RSpec or JUnit

// got test -v -race ./....
// the -v flag will print the output in a verbose style

// the -race flag enables
// Go's race detector which holds discover bugs with concurrency problems

//refactored our test to add a setup method

// Setup
// Execute
// Assert

// Bad test with duplicated code can be worse than bad code

