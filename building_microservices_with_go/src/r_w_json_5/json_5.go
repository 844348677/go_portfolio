package r_w_json_5

import (
	"net/http"
	"encoding/json"
)

// 性能问题　
// Encoder Decoder 减少中间变量，减少代码　直接使用流处理

func helloWorldHandler(w http.ResponseWriter,r *http.Request){
	var request helloWorldRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil{
		http.Error(w,"Bad request " ,http.StatusBadRequest)
		return
	}

	response := helloWorldResponse{Message:"Hello " + request.Name}
	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}

type helloWorldRequest struct{
	Name string `json:"name"`
}
type helloWorldResponse struct{
	Message string `json:"message"`
}

// Routing in net/http
// DefaultServeMux is instance of ServeMux

// ListenAndServe function second parameter nil
// -> DefaultServeMux

//Path
// http.Handle("/images/",newFooHandler())
// http.Handle("/images/persian",newBarHandler())
// http.handle("/images",newBuzzHandler())
// /images		=>Buzz
// /images/		=> Foo
// /images/cat	=> Foo
// /images/cat.jpg	=>Foo
// /images/persion/cat.jpg	=> Bar

// Convenience handlers

// FileServer
// FileServer  server static file , images or other content
// func FileServer(root FileSystem) Handler
// http.Handle("/images", http.FileServer(http.Dir("./images")))
// map the contents of the file system path ./images to the server route /images

// func NotFoundHandler() Handler

//func RedirectHandler(url string,code int) Handler

// func StripPrefix(prefix string,h Handler) Handler

// func TimeoutHandler(h Handler, dt time.Duration, msg string) Handler

