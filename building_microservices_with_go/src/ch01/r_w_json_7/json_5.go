package r_w_json_5

import (
	"net/http"
	"encoding/json"
)

type helloWorldRequest struct{
	Name string `json:"name"`
}
type helloWorldResponse struct{
	Message string `json:"message"`
}

// create a Handler rather than just using HandleFunc

type validationHandler struct{
	next http.Handler
}
// implement the methods in the Handlers interface
// we are going to be chaining handlers together
func (h validationHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request){
	var request helloWorldRequest
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&request)
	if err!= nil{
		http.Error(rw,"Bad request ",http.StatusBadRequest)
		return
	}
	// it has the resposibility for calling ServeHTTP or returning a respose
	h.next.ServeHTTP(rw,r)
}

func newValidationHandler(next http.Handler) http.Handler{
	return validationHandler{next:next}
}

type helloWorldHandler struct{}

func (h helloWorldHandler) ServeHTTP(rw http.ResponseWriter,r *http.Request){
	// name ??
	// response := helloWorldResponse{Message:"Hello " + name}
}


// Context !!!!